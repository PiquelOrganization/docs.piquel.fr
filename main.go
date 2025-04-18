package main

import (
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/server"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
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
	router := mux.NewRouter()

	if config.UseGit {
		// TODO: clone/pull repo
	}

	markdownFiles := source.LoadFiles(config)
	htmlFiles := utils.TranslateFiles(markdownFiles)
	utils.SetupRouterFromLoadedFiles(router, htmlFiles)

	// TODO: admin routes that restart the entire webserver

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
