package render

import (
	"fmt"
	"regexp"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
)

type Renderer interface {
	RenderAllFiles(config *RenderConfig) (map[string][]byte, error)
	RenderFile(path string, config *RenderConfig) ([]byte, error)
}

type RenderConfig struct {
	config.DocsConfig

	highlightStyle *chroma.Style // the style used to format code blocks
}

type RealRenderer struct {
	source source.Source

	singleline    *regexp.Regexp
	multiline     *regexp.Regexp
	htmlFormatter *html.Formatter
}

func NewRealRenderer(source source.Source) (Renderer, error) {
	singleline, err := regexp.Compile(`(?m)^{ *([a-z]+)(?: *\"(.*)\")? */}$`)
	if err != nil {
		return nil, err
	}

	multiline, err := regexp.Compile(`(?m)^{ *([a-z]+) *}\n?((?:.|\n)*?)\n?{/}$`)
	if err != nil {
		return nil, err
	}

	htmlFormatter := html.New()
	if htmlFormatter == nil {
		return nil, fmt.Errorf("Error creating html formatter")
	}

	return &RealRenderer{source, singleline, multiline, htmlFormatter}, nil
}

func (r *RealRenderer) RenderAllFiles(config *RenderConfig) (map[string][]byte, error) {
	if err := r.getHighlightStyle(config); err != nil {
		return nil, err
	}

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
	// just in case already created in RenderAllFiles
	if config.highlightStyle == nil {
		if err := r.getHighlightStyle(config); err != nil {
			return nil, err
		}
	}

	file, err := r.source.LoadRouteFile(path)
	if err != nil {
		return nil, err
	}

	custom, err := r.renderCustom(file, config)
	if err != nil {
		return nil, err
	}
	doc := r.parseMarkdown(custom, config)
	doc = r.fixupAST(doc, config)
	html := r.renderHTML(doc, config)
	return r.addStyles(html, config), nil
}
