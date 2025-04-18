package server

import (
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
)

type Server struct {
    Router http.Handler
	Config *config.Config
}
