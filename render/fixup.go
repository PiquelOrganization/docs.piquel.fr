package render

import (
	"bytes"

	"github.com/gomarkdown/markdown/ast"
)

func (r *RealRenderer) fixupAST(doc ast.Node, config *RenderConfig) ast.Node {
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		switch node := node.(type) {
		case *ast.Link:
			r.fixupLink(node, entering, config)
		}

		return ast.GoToNext
	})
	return doc
}

func (r *RealRenderer) fixupLink(link *ast.Link, entering bool, config *RenderConfig) {
	if !entering {
		return
	}

	if bytes.HasPrefix(link.Destination, []byte("http")) {
		link.AdditionalAttributes = append(link.AdditionalAttributes, "target=\"_blank\"")
	} else {
		if config.RootPath != "" {
			link.Destination = append([]byte(config.RootPath), link.Destination...)
		}
	}
}
