package render

import (
	"fmt"
	"io"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/gomarkdown/markdown/ast"
)

func (r *RealRenderer) renderCodeBlock(w io.Writer, codeBlock *ast.CodeBlock, entering bool, config *RenderConfig) error {
	lang := string(codeBlock.Info)
	source := string(codeBlock.Literal)
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	iterator, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}

	return r.htmlFormatter.Format(w, config.highlightStyle, iterator)
}

func (r *RealRenderer) getHighlightStyle(config *RenderConfig) error {
	styleName := config.StyleName
	if styleName == "" {
		styleName = "tokyonight"
	}

	config.highlightStyle = styles.Get(styleName)
	if config.highlightStyle == nil {
		return fmt.Errorf("Couldn't find style %s", styleName)
	}

	return nil
}
