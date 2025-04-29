package render

import (
	"slices"

	"github.com/PiquelOrganization/docs.piquel.fr/source"
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

type Renderer interface {
	RenderAllFiles(config *RenderConfig) (map[string][]byte, error)
	RenderFile(path string, config *RenderConfig) ([]byte, error)
}

type RenderConfig struct {
	RootPath    string // this will be prepended to any local URLs in the markdown
	UseTailwind bool   // wether to use tailwind classes and settings (notably restore the proper size of titles)
	FullPage    bool   // wether to render a full page (add <!DOCTYPE html> to the top of the page
}

type RealRenderer struct {
	source source.Source
}

func NewRealRenderer(source source.Source) Renderer {
	return &RealRenderer{source: source}
}

func (r *RealRenderer) RenderAllFiles(config *RenderConfig) (map[string][]byte, error) {
	files := map[string][]byte{}
	fileNames, err := r.source.GetAllMarkdown()
	if err != nil {
		return map[string][]byte{}, err
	}

	for _, file := range fileNames {
		renderedFile, err := r.RenderFile(file, config)
		if err != nil {
			return map[string][]byte{}, err
		}
		files[file] = renderedFile
	}
	return files, nil
}

func (r *RealRenderer) RenderFile(path string, config *RenderConfig) ([]byte, error) {
	file, err := r.source.LoadFile(path)
	if err != nil {
		return []byte{}, err
	}

	custom, err := r.renderCustom(file, config)
	if err != nil {
		return []byte{}, err
	}
	doc := r.parseMarkdown(custom, config)
	doc = r.fixupAST(doc, config)
	html := r.renderHTML(doc, config)
	return r.addStyles(html, config), nil
}

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

	styles = slices.Concat([]byte("<style>\n"), styles, []byte("</styles>\n"))
	return slices.Concat(html, styles)
}
