package utils

// SourceDocs:
// MarkdownFile
// Asset
// Stylesheet

// RenderedDocs:
// HTMLFile
// Assets

type Pages map[string][]byte
type Assets []Asset

type Asset struct {
	Name string // the name used to fetch this asset
	Type string // the type (basically the file extension without the .)
	Data []byte
}

type SourceDocs struct {
	Pages    Pages // the standalone pages to serve
	Includes Pages // the markdown that can be included in other pages
	Styles   Assets
	Assets   Assets
}

type RenderedDocs struct {
	Pages  Pages // the pages to be served
	Static Assets
}
