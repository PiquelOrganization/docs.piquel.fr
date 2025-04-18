package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    log.Printf("Initializing documentation service...\n")

    // TODO: load config (env)
    // TODO: load documentation

    router := mux.NewRouter()

    // TODO: middleware (logging, ...)

    log.Printf("[Router] Initialized\n")

    router.HandleFunc("/", testHandler).Methods(http.MethodGet)

    address := fmt.Sprintf("0.0.0.0:8080")

    log.Printf("[Router] Starting router...\n")
    log.Printf("[Router] Listening on %s!\n", address)

    err := http.ListenAndServe(address, router)
    panic(err)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}
