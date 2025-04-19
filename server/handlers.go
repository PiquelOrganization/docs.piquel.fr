package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) GithubPushHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var payload any
    err := decoder.Decode(&payload)
    if err != nil {
        panic(err)
    }

    json, err := json.Marshal(payload)
    if err != nil {
        panic(err)
    }

    log.Printf("Received:\n %s", json)
}
