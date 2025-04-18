package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/source"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Initializing documentation service...\n")

	config := config.LoadConfig()
	router := mux.NewRouter()
	source := source.GetSource(config)

	markdownFiles := source.LoadFiles()

	htmlFiles := utils.TranslateFiles(markdownFiles)
	utils.SetupRouterFromLoadedFiles(router, htmlFiles)

	// TODO: admin routes that restart the entire webserver

	address := fmt.Sprintf("0.0.0.0:%s", config.Port)

	log.Printf("[Server] Starting server on %s!\n", address)

	server := http.Server{
		Addr:    address,
		Handler: router,
	}

    err := server.ListenAndServe()
}
