package utils

import (
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

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

		w.Header().Set("Content-Type", "text/html")
		w.Write(html)
	}
}
