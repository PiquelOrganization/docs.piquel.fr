package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
)

type Server struct {
	http.Server
	config *config.Config
    requestingRestart bool
}

func NewServer(config *config.Config, handler http.Handler) Server {
	address := fmt.Sprintf("0.0.0.0:%s", config.Port)

	return Server{
		Server: http.Server{
			Addr:    address,
			Handler: handler,
		},
		config: config,
	}
}

func (s *Server) Serve(done chan error) {
	log.Printf("[Server] Starting server on %s\n", s.Server.Addr)

	err := s.ListenAndServe()
	done <- err
}

func (s *Server) IsRequestingRestart() bool {
    return s.requestingRestart
}

func (s *Server) Restart() {
	log.Printf("[Server] Restart request received...")
    s.requestingRestart = true
	s.Shutdown(context.Background())
}

func (s *Server) Shutdown(context context.Context) {
	log.Printf("[Server] Shutting down...")
	s.Server.Shutdown(context)
}
