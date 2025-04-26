package render

import (
	"github.com/PiquelOrganization/docs.piquel.fr/source"
)

type Renderer interface {
	RenderAllFiles() map[string][]byte
	RenderFile(path string, config RenderConfig) []byte
}

func NewRealRenderer(source source.Source) Renderer {
	return &RealRenderer{source: source}
}

type RealRenderer struct {
	source source.Source
}

func (r *RealRenderer) RenderAllFiles() map[string][]byte {
	return map[string][]byte{}
}

func (r *RealRenderer) RenderFile(path string, config RenderConfig) []byte {
	return []byte{}
}

// func (r *RealRenderer) renderDocs() {
// 	includes := utils.Files{}
// 	for name, data := range r.input.Includes {
// 		newName := strings.TrimRight(name, ".md")
// 		includes[newName] = data
// 	}
//
// 	// TODO: render properly
// 	// namely: includes & styles
// }
//
// func (r *RealRenderer) renderHTML() {
// 	outputPages := make(utils.Files)
//
// 	for route, data := range r.input.Pages {
// 		route = strings.ReplaceAll(route, ".md", ".html")
// 		outputPages[route] = markdownToHTML(data)
// 	}
//
// 	r.output.Pages = outputPages
// 	r.output.Static = r.input.Styles
// 	maps.Copy(r.output.Static, r.input.Assets)
// }
