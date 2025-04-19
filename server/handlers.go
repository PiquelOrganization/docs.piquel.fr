package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GithubPushHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var payload any
    err := decoder.Decode(&payload)
    if err != nil {
        panic(err)
    }
}
