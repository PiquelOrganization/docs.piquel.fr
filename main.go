package main

import (
	"log"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/config"
	"github.com/PiquelOrganization/docs.piquel.fr/git"
	"github.com/PiquelOrganization/docs.piquel.fr/middleware"
	"github.com/PiquelOrganization/docs.piquel.fr/output"
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

	router := mux.NewRouter()
    middleware.Setup(router)

	source := source.NewSource(config)
	renderer := render.NewRenderer(config, router, source)
	output := output.NewOutput(router, config, renderer)
	DocsServer := server.NewServer(config, router)

	if config.UseGit {
		router.HandleFunc("/gh-push", DocsServer.GithubPushHandler).Methods(http.MethodPost)

		if err := git.Status(config.DataPath); err == nil {
			err = git.Pull(config.DataPath)
			if err != nil {
				panic(err)
			}
		} else {
			err := git.Clone(config.Repository, config.DataPath)
			if err != nil {
				panic(err)
			}
		}
	}

	output.Output()

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
