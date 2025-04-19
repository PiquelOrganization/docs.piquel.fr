package server

import (
	"io"
	"net/http"
	"strings"

	"github.com/PiquelOrganization/docs.piquel.fr/utils"
)

func (s *Server) GithubPushHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	header := r.Header.Get("X-Hub-Signature-256")
	header = strings.Split(header, "=")[1]

	isGithub, err := utils.VerifySignature(string(body), header, s.config.WebhookSecret)
	if err != nil {
		panic(err)
	}

	if isGithub {
		s.Restart()
	}
}
