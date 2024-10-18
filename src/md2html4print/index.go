package md2html4print

import (
	_ "embed"
	"text/template"
)

//go:embed resources/index.go.html
var indexTemplateString string

func indexTemplate() *template.Template {
	var indexTemplate, err = template.New("index").Parse(indexTemplateString)
	if err != nil {
		panic(err)
	}

	return indexTemplate
}
