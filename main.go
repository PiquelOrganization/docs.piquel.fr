package main

import (
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/render"
	"github.com/PiquelOrganization/docs.piquel.fr/server"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/gorilla/mux"
)

var DocsServer server.Server

func main() {
	log.Printf("Initializing documentation service...\n")

	config := config.LoadConfig()

	runDocsService(config)
}

func runDocsService(config *config.Config) {
	log.Printf("Starting documentation service...\n")

	if config.UseGit {
		// TODO: clone/pull repo
	}

	router := mux.NewRouter()
	source := source.NewSource(config)
	renderer := render.NewRenderer(config, router, source)

	renderer.RenderDocs()
	renderer.RenderHTML()
	renderer.SetupRouter()

	DocsServer := server.NewServer(config, router)

	done := make(chan error)
	go DocsServer.Serve(done)

	err := <-done
	if err != http.ErrServerClosed {
		panic(err)
	}

	log.Printf("[Server] Shut down without issue\n")
	log.Printf("Stopped documentation service\n")

	if DocsServer.IsRequestingRestart() {
		runDocsService(config)
	}
}
