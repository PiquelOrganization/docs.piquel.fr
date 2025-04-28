package render

import (
	"bytes"
	"io"
	"log"
	"regexp"

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

	custom, err := r.renderCustom(file, config)
	if err != nil {
		return []byte{}, err
	}
	doc := r.parseMarkdown(custom, config)
	return r.renderHTML(doc, config), nil
}

func (r *RealRenderer) renderCustom(md []byte, config *RenderConfig) ([]byte, error) {
	multiline, err := regexp.Compile(`(?m)^{ *([a-z]+)(?: *\"(.*)\")? *}\n?((?:.|\n)*?)\n?{/}$`)
	if err != nil {
		return []byte{}, err
	}
	log.Printf("%s\n", multiline.FindAll(md, -1))

	singleline, err := regexp.Compile(`(?m)^{ *([a-z]+)(?: *\"(.*)\")? */}$`)
	if err != nil {
		return []byte{}, err
	}
	log.Printf("%s\n", singleline.FindAll(md, -1))

	return md, nil
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
	return func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
		switch node := node.(type) {
		case *ast.Link:
			r.renderLink(w, node, entering, config)
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
		if config.RootPath != "" {
			link.Destination = append([]byte(config.RootPath), link.Destination...)
		}
	}
}
