package builder

import (
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/spf13/viper"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/anchor"
)

func NewParser() goldmark.Markdown {
	// 初始化Markdown解析器

	parserOption := goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	)

	defaultExtensions := goldmark.WithExtensions(
		extension.GFM,
		meta.Meta,
		emoji.Emoji,
		extension.CJK,
		extension.DefinitionList,
		extension.Footnote,
		extension.Typographer,
		&anchor.Extender{
			Texter:   anchor.Text("#"),
			Position: anchor.Before,
		},
	)

	var parser goldmark.Markdown
	if viper.GetBool("style.codeBlock.enable") {
		otherExtentions := goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle(viper.GetString("style.codeBlock.style")),
				highlighting.WithFormatOptions(
					html.WithLineNumbers(true),
				),
			),
		)
		parser = goldmark.New(
			parserOption, defaultExtensions, otherExtentions,
		)
	} else {
		parser = goldmark.New(parserOption, defaultExtensions)
	}
	return parser
}
