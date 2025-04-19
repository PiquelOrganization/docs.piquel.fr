package server

import (
	"log"
	"net/http"
)

func (s *Server) GithubPushHandler(w http.ResponseWriter, r *http.Request) {
    body := []byte{}
    size, err := r.Body.Read(body)
    if err != nil {
        panic(err)
    }

    log.Printf("Received: %d bytes\n%s\n", size, string(body[:]))
}
