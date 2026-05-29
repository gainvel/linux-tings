package render

import (
	"strings"

	"refr/internal/config"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

type Highlighter struct {
	cfg   config.SyntaxConfig
	style *chroma.Style
}

func NewHighlighter(syntax config.SyntaxConfig) *Highlighter {
	h := &Highlighter{cfg: syntax}
	h.style = buildStyle(syntax)
	return h
}

func (h *Highlighter) Highlight(code, language string) string {
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	iter, err := lexer.Tokenise(nil, code)
	if err != nil {
		return code
	}

	formatter := formatters.Get("terminal256")
	if formatter == nil {
		return code
	}

	var buf strings.Builder
	if err := formatter.Format(&buf, h.style, iter); err != nil {
		return code
	}
	return buf.String()
}

func buildStyle(cfg config.SyntaxConfig) *chroma.Style {
	base := styles.Get(cfg.Base)
	if base == nil {
		base = styles.Fallback
	}

	builder := base.Builder()

	overrides := map[chroma.TokenType]string{
		chroma.Keyword:       cfg.Keyword,
		chroma.LiteralString: cfg.String,
		chroma.LiteralNumber: cfg.Number,
		chroma.Comment:       cfg.Comment,
		chroma.NameClass:     cfg.Type,
		chroma.NameFunction:  cfg.Function,
		chroma.Operator:      cfg.Operator,
		chroma.NameVariable:  cfg.Variable,
		chroma.Error:         cfg.Error,
	}

	for tokenType, spec := range overrides {
		if spec == "" {
			continue
		}
		builder.Add(tokenType, spec)
	}

	style, err := builder.Build()
	if err != nil {
		return base
	}
	return style
}
