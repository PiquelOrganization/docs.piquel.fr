package server

import (
	"log"
	"net/http"
)

func (s *Server) GithubPushHandler(w http.ResponseWriter, r *http.Request) {
    body := []byte{}
    _, err := r.Body.Read(body)
    if err != nil {
        panic(err)
    }

    log.Printf("Received: %s", string(body[:]))
}
