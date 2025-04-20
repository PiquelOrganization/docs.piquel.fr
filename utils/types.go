package utils

type Files map[string][]byte

type SourceDocs struct {
	Pages    Files // the standalone pages to serve
	Includes Files // the markdown that can be included in other pages
	Styles   Files
	Assets   Files
}

type RenderedDocs struct {
	Pages  Files // the pages to be served
	Static Files
}
