package render

import (
	"bytes"
	"slices"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

const tailwindBase = `
    h1 { font-size: 2em; }
    h2 { font-size: 1.5em; }
    h3 { font-size: 1.17em; }
    h4 { font-size: 1em; }
    h5 { font-size: 0.83em; }
    h6 { font-size: 0.67em; }
`

func (r *RealRenderer) parseMarkdown(md []byte, config *RenderConfig) ast.Node {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	return p.Parse(md)
}

func (r *RealRenderer) renderHTML(doc ast.Node, config *RenderConfig) []byte {
	htmlFlags := html.CommonFlags

	if config.FullPage {
		htmlFlags = htmlFlags | html.CompletePage
	}

	options := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(options)

	return markdown.Render(doc, renderer)
}

func (r *RealRenderer) addStyles(html []byte, config *RenderConfig) []byte {
	var styles []byte

	if config.UseTailwind {
		styles = append(styles, []byte(tailwindBase)...)
	}

	if styles == nil {
		return html
	}

	styles = slices.Concat([]byte("<style>\n"), styles, []byte("</style>\n"))
	return slices.Concat(html, styles)
}

func (r *RealRenderer) fixupAST(doc ast.Node, config *RenderConfig) ast.Node {
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		switch node := node.(type) {
		case *ast.Link:
			r.fixupLink(node, entering, config)
		}

		return ast.GoToNext
	})
	return doc
}

func (r *RealRenderer) fixupLink(link *ast.Link, entering bool, config *RenderConfig) {
	if !entering {
		return
	}

	if bytes.HasPrefix(link.Destination, []byte("http")) {
		link.AdditionalAttributes = append(link.AdditionalAttributes, "target=\"_blank\"")
	} else {
		link.Destination = slices.Concat([]byte(config.RootPath), utils.FormatLocalPath(link.Destination, ".md"))
	}
}
