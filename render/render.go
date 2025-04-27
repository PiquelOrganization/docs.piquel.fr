package render

import (
	"bytes"
	"io"

	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Renderer interface {
	RenderAllFiles(config *RenderConfig) (map[string][]byte, error)
	RenderFile(path string, config *RenderConfig) ([]byte, error)
}

type RenderConfig struct {
	RootPath string // this will be prepended to any local URLs in the markdown
}

func NewRealRenderer(source source.Source) Renderer {
	return &RealRenderer{source: source}
}

type RealRenderer struct {
	source source.Source
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

	// TODO: do the custom rendering

	doc := r.parseMarkdown(file, config)
	return r.renderHTML(doc, config), nil
}

func (r *RealRenderer) parseMarkdown(md []byte, config *RenderConfig) ast.Node {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	return p.Parse(md)
}

func (r *RealRenderer) renderHTML(doc ast.Node, config *RenderConfig) []byte {
	htmlFlags := html.CommonFlags
	options := html.RendererOptions{
		Flags:          htmlFlags,
		RenderNodeHook: r.renderHook(config),
	}
	renderer := html.NewRenderer(options)

	return markdown.Render(doc, renderer)
}

func (r *RealRenderer) renderHook(config *RenderConfig) html.RenderNodeFunc {
	if config == nil {
		return func(io.Writer, ast.Node, bool) (ast.WalkStatus, bool) { return ast.GoToNext, false }
	}

	return func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
		if link, ok := node.(*ast.Link); ok {
			r.renderLink(w, link, entering, config)
			return ast.GoToNext, false
		}

		return ast.GoToNext, false
	}
}

func (r *RealRenderer) renderLink(w io.Writer, link *ast.Link, entering bool, config *RenderConfig) {
	if !entering {
		return
	}

	if bytes.HasPrefix(link.Destination, []byte("http")) {
		link.AdditionalAttributes = append(link.AdditionalAttributes, "target=\"_blank\"")
	} else {
		link.Destination = append([]byte(config.RootPath), link.Destination...)
	}
}
