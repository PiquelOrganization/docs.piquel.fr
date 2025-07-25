package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/piquel-fr/piquel-docs/config"
	"github.com/piquel-fr/piquel-docs/render"
	"github.com/piquel-fr/piquel-docs/source"
	"github.com/piquel-fr/piquel-docs/utils"
)

type Handler struct {
	source        source.Source
	renderer      render.Renderer
	staticHandler http.Handler // the handler that will serve static files
	docsConfig    *config.DocsConfig
}

func NewHandler(config *config.Config, source source.Source, renderer render.Renderer, staticHandler http.Handler) *Handler {
	return &Handler{source, renderer, staticHandler, &config.Config}
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
	path := utils.FormatLocalPathString(r.URL.Path, ".md")

	if path == "/gh-push" && r.Method == http.MethodPost {
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
		http.NotFound(w, r)
		return
	}

	if path == "/" {
		h.handleDocsPath(w, r, h.docsConfig.HomePage)
		return
	}

	h.handleDocsPath(w, r, path)
}

func (h *Handler) handleDocsPath(w http.ResponseWriter, r *http.Request, path string) {
	queryConfig := &config.DocsConfig{}
	root := r.URL.Query().Get("root")
	if root != "" {
		queryConfig.Root = utils.FormatLocalPathString(root, "")
	}
	_, queryConfig.UseTailwind = r.URL.Query()["tailwind"]
	queryConfig.HighlightStyle = r.URL.Query().Get("highlight_style")
	_, queryConfig.FullPage = r.URL.Query()["full_page"]

	jsonConfig := &config.DocsConfig{}
	if r.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&jsonConfig)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			panic(err)
		}
		jsonConfig.Root = utils.FormatLocalPathString(jsonConfig.Root, ".md")
	}

	repoConfig := h.docsConfig

	renderConfig := &render.RenderConfig{}
	renderConfig.FullPage = queryConfig.FullPage || jsonConfig.FullPage || repoConfig.FullPage
	renderConfig.UseTailwind = queryConfig.UseTailwind || jsonConfig.UseTailwind || repoConfig.UseTailwind

	if queryConfig.Root == "" {
		if jsonConfig.Root == "" {
			renderConfig.Root = repoConfig.Root
		} else {
			renderConfig.Root = jsonConfig.Root
		}
	} else {
		renderConfig.Root = queryConfig.Root
	}

	if queryConfig.HighlightStyle == "" {
		if jsonConfig.HighlightStyle == "" {
			renderConfig.HighlightStyle = repoConfig.HighlightStyle
		} else {
			renderConfig.HighlightStyle = jsonConfig.HighlightStyle
		}
	} else {
		renderConfig.HighlightStyle = queryConfig.HighlightStyle
	}

	html, err := h.renderer.RenderFile(path, renderConfig)
	if err != nil {
		if internalError, ok := err.(utils.Error); ok {
			internalError.Handle(w)
			return
		}

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
