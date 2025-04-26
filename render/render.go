package render

import (
	"maps"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
	"github.com/gorilla/mux"
)

type Renderer interface {
	RenderOutput() utils.RenderedDocs
}

func NewRealRenderer(config *config.Config, router *mux.Router, source source.Source) Renderer {
	return &RealRenderer{config: config, router: router, source: source}
}

type RealRenderer struct {
	config *config.Config
	router *mux.Router
	source source.Source

	input  utils.SourceDocs
	output utils.RenderedDocs
}

func (r *RealRenderer) renderDocs() {
	includes := utils.Files{}
	for name, data := range r.input.Includes {
		newName := strings.TrimRight(name, ".md")
		includes[newName] = data
	}

	// TODO: render properly
	// namely: includes & styles
}

func (r *RealRenderer) renderHTML() {
	outputPages := make(utils.Files)

	for route, data := range r.input.Pages {
		route = strings.ReplaceAll(route, ".md", ".html")
		outputPages[route] = markdownToHTML(data)
	}

	r.output.Pages = outputPages
	r.output.Static = r.input.Styles
	maps.Copy(r.output.Static, r.input.Assets)
}

func (r *RealRenderer) RenderOutput() utils.RenderedDocs {
	r.input = r.source.LoadFiles()
	r.renderDocs()
	r.renderHTML()
	return r.output
}
