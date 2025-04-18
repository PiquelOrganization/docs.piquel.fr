package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
)

func InitServer(router http.Handler, config *config.Config) *Server {
    log.Printf("[Server] Initializing...\n")

    server := &Server{
        Router: router,
        Config: config,
    }

    log.Printf("[Server] Intialized\n")

    return server
}

func (s *Server) Serve() {
    address := fmt.Sprintf("0.0.0.0:%s", s.Config.Port)

    log.Printf("[Server] Starting server on %s!\n", address)
    log.Fatalf("%s\n", http.ListenAndServe(address, s.Router))
}
