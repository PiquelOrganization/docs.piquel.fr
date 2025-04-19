package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) GithubPushHandler(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    var test any
    err := decoder.Decode(&test)
    if err != nil {
        panic(err)
    }

    log.Printf("Received:\n %v\n\nwith payload:\n%v", r, test)
}
