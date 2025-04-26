package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/handlers"
	"github.com/PiquelOrganization/docs.piquel.fr/middleware"
	"github.com/PiquelOrganization/docs.piquel.fr/render"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Initializing documentation service...\n")

	config := config.LoadConfig()

	router := mux.NewRouter()
	middleware.Setup(router)

	source := source.NewGitSource(config)
	source.Update()

	renderer := render.NewRealRenderer(source)

	handler := handlers.NewHandler(config, source, renderer)
	router.HandleFunc("/gh-push", handler.GithubPushHandler).Methods(http.MethodPost)

	done := make(chan error)
	go func() {
		address := fmt.Sprintf("0.0.0.0:%s", config.Port)
		log.Printf("[Server] Starting server on %s\n", address)

		err := http.ListenAndServe(address, router)
		done <- err
	}()

	err := <-done
	if err != http.ErrServerClosed {
		panic(err)
	}

	log.Printf("[Server] Shut down without issue\n")
	log.Printf("Stopped documentation service\n")
}
