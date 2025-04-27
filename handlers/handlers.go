package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/render"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
)

type Handler struct {
	source   source.Source
	renderer render.Renderer

	staticHandler http.Handler // the handler that will serve static files
	homePage      string
}

func NewHandler(config *config.Config, source source.Source, renderer render.Renderer, staticHandler http.Handler) *Handler {
	return &Handler{source, renderer, staticHandler, config.HomePage}
}

func (h *Handler) GithubPushHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// header := r.Header.Get("X-Hub-Signature-256")
	// header = strings.Split(header, "=")[1]

	// isGithub, err := utils.VerifySignature(string(body), header, s.config.WebhookSecret)
	// if err != nil {
	// 	panic(err)
	// }

	// if isGithub {
	// 	s.Restart()
	// }

	if err := h.source.Update(); err != nil {
		panic(err)
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")

	if path == "gh-push" && r.Method == http.MethodPost {
		h.GithubPushHandler(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are available for documentation", http.StatusBadRequest)
		return
	}

	if strings.Contains(path, ".") {
		if h.staticHandler != nil {
			h.staticHandler.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if path == "" {
		h.handleDocsPath(w, r, h.homePage)
		return
	}

	h.handleDocsPath(w, r, path)
}

func (h *Handler) handleDocsPath(w http.ResponseWriter, r *http.Request, path string) {
	renderConfig := &render.RenderConfig{}

	root := r.URL.Query().Get("root")
	renderConfig.RootPath = fmt.Sprintf("/%s/", strings.Trim(root, "/"))

	html, err := h.renderer.RenderFile(path, renderConfig)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		panic(err)
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}
