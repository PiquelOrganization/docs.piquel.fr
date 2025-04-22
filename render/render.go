package render

import (
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
	"github.com/gorilla/mux"
)

type Renderer interface {
	Init()
	RenderDocs()
	RenderHTML()
	GetOutput() utils.RenderedDocs
	SetupRouter()
}

func NewRenderer(config *config.Config, router *mux.Router, source source.Source) Renderer {
	return &RealRenderer{config: config, router: router, source: source}
}

type RealRenderer struct {
	config *config.Config
	router *mux.Router
	source source.Source

	input  utils.SourceDocs
	output utils.RenderedDocs
}

func (r *RealRenderer) Init() {
	r.input = r.source.LoadFiles()
}

func (r *RealRenderer) RenderDocs() {
	// TODO: render properly
	// namely: includes & styles
}

func (r *RealRenderer) RenderHTML() {
	outputPages := make(utils.Files)

	for route, data := range r.input.Pages {
		outputPages[route] = utils.MarkdownToHTML(data)
	}

	r.output.Pages = outputPages
}

func (r *RealRenderer) GetOutput() utils.RenderedDocs {
	return r.output
}

func (r *RealRenderer) SetupRouter() {
	for route, data := range r.output.Pages {
		handler := utils.GenerateHandler(data)
		r.router.HandleFunc(route, handler).Methods(http.MethodGet)
	}
	// TODO: serve assets & styles
}
