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
	SetupRouter()
}

func NewRenderer(config *config.Config, router *mux.Router, source source.Source) Renderer {
	return &RealRenderer{config: config, router: router, source: source}
}

type RealRenderer struct {
	config *config.Config
	router *mux.Router
	source source.Source

	files utils.Files
}

func (r *RealRenderer) Init() {
	r.files = r.source.LoadFiles()
}

func (r *RealRenderer) RenderDocs() {
	// TODO: render the documentation (namely include syntax, ...)
}

func (r *RealRenderer) RenderHTML() {
	htmlFiles := utils.Files{}
	for _, file := range r.files {
		html := utils.MarkdownToHTML(file.Data)
		htmlFile := utils.File{Route: file.Route, Data: html}
		htmlFiles = append(htmlFiles, htmlFile)
	}
	r.files = htmlFiles
}

func (r *RealRenderer) SetupRouter() {
	for _, file := range r.files {
		handler := utils.GenerateHandler(file.Data)
		r.router.HandleFunc(file.Route, handler).Methods(http.MethodGet)
	}
}
