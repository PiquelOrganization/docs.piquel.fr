package render

import (
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Renderer interface {
	RenderAllFiles(config RenderConfig) (map[string][]byte, error)
	RenderFile(path string, config RenderConfig) ([]byte, error)
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

func (r *RealRenderer) RenderAllFiles(config RenderConfig) (map[string][]byte, error) {
	files := map[string][]byte{}
	fileNames, err := r.source.GetAllFiles()
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

func (r *RealRenderer) RenderFile(path string, config RenderConfig) ([]byte, error) {
	file, err := r.source.LoadFile(path)
	if err != nil {
		return []byte{}, err
	}

	// TODO: do the custom rendering

	return r.renderHTML(file, config), nil
}

func (r *RealRenderer) renderHTML(md []byte, config RenderConfig) []byte {
	// TODO: modify the renderer to use RenderConfig data

	// markdown parser
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	// html renderer
	htmlFlags := html.CommonFlags
	options := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(options)

	return markdown.ToHTML(md, p, renderer)
}
