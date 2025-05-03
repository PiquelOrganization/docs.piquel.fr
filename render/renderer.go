package render

import (
	"regexp"

	"github.com/PiquelOrganization/docs.piquel.fr/source"
)

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

	singleline *regexp.Regexp
	multiline  *regexp.Regexp
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

	return &RealRenderer{
		source:     source,
		singleline: singleline,
		multiline:  multiline,
	}, nil
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
	file, err := r.source.LoadRouteFile(path)
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
