package markdown

import (
	"bytes"
	"fmt"
	"os"

	wikilink "github.com/abhinav/goldmark-wikilink"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type Markdown struct {
	gm goldmark.Markdown
}

func NewMarkdown() *Markdown {
	return &Markdown{
		gm: goldmark.New(
			goldmark.WithExtensions(
				extension.GFM,
				&wikilink.Extender{},
				meta.Meta,
			),
			goldmark.WithRendererOptions(
				html.WithHardWraps(),
				html.WithXHTML(),
			),
		),
	}
}

func (md *Markdown) Convert(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	context := parser.NewContext()
	if err := md.gm.Convert(data, &buf, parser.WithContext(context)); err != nil {
		return "", fmt.Errorf("fail to convert md to html: %v", err)
	}

	return buf.String(), nil
}
