package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/utils"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Initializing documentation service...\n")

	config := config.LoadConfig()
	router := mux.NewRouter()

	// TODO: actually populate the var
	markdownFiles := utils.Files{}

	htmlFiles := utils.Files{}
	for _, file := range markdownFiles {
		html := utils.MarkdownToHTML(file.Data)
		htmlFile := utils.File{Path: file.Path, Data: html}
		htmlFiles = append(htmlFiles, htmlFile)
	}

	for _, file := range htmlFiles {
		handler := utils.GenerateHandler(file.Data)
		router.HandleFunc(file.Path, handler).Methods(http.MethodGet)
	}

	address := fmt.Sprintf("0.0.0.0:%s", config.Port)

	log.Printf("[Server] Starting server on %s!\n", address)
	log.Fatalf("%s\n", http.ListenAndServe(address, router))
}
