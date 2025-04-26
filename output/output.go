package output

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strings"
//
// 	"github.com/PiquelOrganization/docs.piquel.fr/config"
// 	"github.com/PiquelOrganization/docs.piquel.fr/render"
// 	"github.com/PiquelOrganization/docs.piquel.fr/utils"
// 	"github.com/gorilla/mux"
// )
//
// type Output interface {
// 	Output() error
// }
//
// type RouterOutput struct {
// 	router   *mux.Router
// 	config   *config.Config
// 	renderer render.Renderer
// }
//
// func NewRouterOutput(router *mux.Router, config *config.Config, renderer render.Renderer) Output {
// 	return &RouterOutput{router, config, renderer}
// }
//
// func (o *RouterOutput) Output() error {
// 	output := o.renderer.RenderAllFiles()
//
// 	if o.config.HomePage != "" {
// 		homePage := strings.ReplaceAll(o.config.HomePage, ".md", ".html")
// 		pageData := output.Pages[homePage]
// 		handler := utils.GenerateHandler(pageData)
// 		o.router.HandleFunc("/", handler).Methods(http.MethodGet)
// 	}
//
// 	for route, data := range output.Pages {
// 		route = strings.TrimRight(route, ".html")
// 		handler := utils.GenerateHandler(data)
// 		o.router.HandleFunc(fmt.Sprintf("/%s", route), handler).Methods(http.MethodGet)
// 	}
//
// 	staticDir := "/home/piquel/tmp"
// 	for file, data := range output.Static {
// 		path := fmt.Sprintf("%s/%s", staticDir, file)
// 		err := os.WriteFile(path, data, os.FileMode(int(0644)))
// 		if err != nil {
// 			return err
// 		}
// 	}
//
// 	o.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))
// 	return nil
// }
