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
	router := mux.NewRouter()
	source := source.GetSource(config)

	markdownFiles := source.LoadFiles()

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

	log.Printf("[Server] Shut down without issue")
}
