package render

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PiquelOrganization/docs.piquel.fr/utils"
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
	styleName := config.HighlightStyle
	if styleName == "" {
		styleName = "tokyonight"
	}

	config.highlightStyle = styles.Get(styleName)
	if config.highlightStyle == nil {
		return utils.NewError(fmt.Sprintf("Couldn't find the style %s", styleName), http.StatusBadRequest)
	}

	return nil
}
