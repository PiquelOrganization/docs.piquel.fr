package output

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/render"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
	"github.com/gorilla/mux"
)

type Output interface {
	Output() error
}

type RouterOutput struct {
	router   *mux.Router
	renderer render.Renderer
}

func NewOutput(router *mux.Router, renderer render.Renderer) Output {
	return &RouterOutput{router, renderer}
}

func (o *RouterOutput) Output() error {
	output := o.renderer.RenderOutput()

	for route, data := range output.Pages {
		handler := utils.GenerateHandler(data)
		o.router.HandleFunc(route, handler).Methods(http.MethodGet)
	}

	staticDir := "/home/piquel/tmp"
	for file, data := range output.Static {
		path := fmt.Sprintf("%s/%s", staticDir, strings.Trim(file, "/"))
		err := os.WriteFile(path, data, os.FileMode(int(0644)))
		if err != nil {
			return err
		}
	}

	o.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))
	return nil
}
