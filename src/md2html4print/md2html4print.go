package md2html4print

import (
	"bytes"
	"io"
	"io/fs"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

func ParseDocument(fileSystem fs.FS) (document, error) {
	sourceDirectoryList, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return document{}, err
	}

	d := document{
		Title:    "Default Title",
		Pages:    []page{},
		PageSize: "letter",
	}

	for _, sourceDirectoryItem := range sourceDirectoryList {
		if strings.HasSuffix(sourceDirectoryItem.Name(), ".yaml") {
			y, err := fileSystem.Open(sourceDirectoryItem.Name())
			if err != nil {
				return document{}, err
			}

			yData, err := io.ReadAll(y)
			if err != nil {
				return document{}, err
			}

			if err := yaml.Unmarshal(yData, &d); err != nil {
				return document{}, err
			}

			continue
		}

		if !strings.HasSuffix(sourceDirectoryItem.Name(), ".md") {
			continue
		}

		source, err := fs.ReadFile(fileSystem, sourceDirectoryItem.Name())
		if err != nil {
			return document{}, err
		}

		page, err := parseFileContents(source)
		if err != nil {
			return document{}, err
		}

		d.Pages = append(d.Pages, page)
	}

	return d, nil
}

func parseFileContents(contents []byte) (page, error) {
	contentParts := bytes.SplitN(contents, []byte("---"), 3)

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	// Parse the Markdown
	var buf bytes.Buffer
	if err := md.Convert(contentParts[2], &buf); err != nil {
		return page{}, err
	}

	// Parse the FrontMatter
	frontMatter := frontMatter{}
	if err := yaml.Unmarshal(contentParts[1], &frontMatter); err != nil {
		return page{}, err
	}

	return page{
		HTML:        string(buf.Bytes()),
		FrontMatter: frontMatter,
	}, nil
}
