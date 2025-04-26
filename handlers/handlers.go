package handlers

import (
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/render"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
)

type Handler struct {
	config   *config.Config
	source   source.Source
	renderer render.Renderer
}

func NewHandler(config *config.Config, source source.Source, renderer render.Renderer) *Handler {
	return &Handler{config, source, renderer}
}

func (h *Handler) RootHandler(w http.ResponseWriter, r *http.Request) {
	h.handleDocsPath(w, r, h.config.HomePage)
}

func (h *Handler) handleDocsPath(w http.ResponseWriter, r *http.Request, path string) {
	// TODO
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
