package main

import (
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/server"
	"github.com/gorilla/mux"
)

func main() {
    log.Printf("Initializing documentation service...\n")

    // config
    config := config.LoadConfig()

    // router
    log.Printf("[Router] Initializing...\n")
    router := mux.NewRouter()
    router.HandleFunc("/", testHandler).Methods(http.MethodGet)
    log.Printf("[Router] Initialized\n")

    server := server.InitServer(router, config)
    server.Serve()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}
