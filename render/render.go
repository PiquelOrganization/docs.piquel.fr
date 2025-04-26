package render

import (
	"github.com/PiquelOrganization/docs.piquel.fr/source"
)

type Renderer interface {
	RenderAllFiles(config RenderConfig) (map[string][]byte, error)
	RenderFile(path string, config RenderConfig) ([]byte, error)
}

type RenderConfig struct{}

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
	// TODO
	return []byte{}, nil
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
