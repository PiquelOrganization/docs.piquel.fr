package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piquel-fr/piquel-docs/config"
	"github.com/piquel-fr/piquel-docs/handlers"
	"github.com/piquel-fr/piquel-docs/middleware"
	"github.com/piquel-fr/piquel-docs/render"
	"github.com/piquel-fr/piquel-docs/source"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Initializing documentation service...\n")

	config := config.LoadConfig()

	router := mux.NewRouter()
	middleware.Setup(router)

	source := source.NewGitSource(config)
	source.Update()

	renderer, err := render.NewRealRenderer(source)
	if err != nil {
		panic(err)
	}

	var staticHandler http.Handler
	if assetsPath := source.GetAssetsPath(); assetsPath != "" {
		staticHandler = http.StripPrefix("/", http.FileServer(http.Dir(assetsPath)))
	}
	handler := handlers.NewHandler(config, source, renderer, staticHandler)
	router.PathPrefix("/").Handler(handler).Methods(http.MethodPost, http.MethodGet)

	done := make(chan error)
	go func() {
		address := fmt.Sprintf("0.0.0.0:%s", config.Envs.Port)
		log.Printf("[Server] Starting server on %s\n", address)

		err := http.ListenAndServe(address, router)
		done <- err
	}()

	err = <-done
	if err != http.ErrServerClosed {
		panic(err)
	}

	log.Printf("[Server] Shut down without issue\n")
	log.Printf("Stopped documentation service\n")
}
