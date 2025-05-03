package render

import (
	"io"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/gomarkdown/markdown/ast"
)

func (r *RealRenderer) renderCodeBlock(w io.Writer, codeBlock *ast.CodeBlock, entering bool) error {
	_, err := io.WriteString(w, "\n<pre><code>\n")
	if err != nil {
		return err
	}

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

	err = r.htmlFormatter.Format(w, r.highlightStyle, iterator)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, "\n</pre></code>\n")
	return err
}
