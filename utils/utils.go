package utils

import (
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/gorilla/mux"
)

type Files []File

type File struct {
	Path string
	Data []byte
}

func SetupRouterFromLoadedFiles(router *mux.Router, files Files) {
	for _, file := range files {
		handler := GenerateHandler(file.Data)
		router.HandleFunc(file.Path, handler).Methods(http.MethodGet)
	}
}

func TranslateFiles(markdownFiles Files) Files {
	htmlFiles := []File{}
	for _, file := range markdownFiles {
		html := MarkdownToHTML(file.Data)
		htmlFile := File{Path: file.Path, Data: html}
		htmlFiles = append(htmlFiles, htmlFile)
	}
	return htmlFiles
}

func MarkdownToHTML(md []byte) []byte {
	// markdown parser
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	// html renderer
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	options := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(options)

	return markdown.ToHTML(md, p, renderer)
}

func GenerateHandler(html []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(html)
	}
}
