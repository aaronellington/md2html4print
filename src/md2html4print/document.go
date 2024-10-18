package md2html4print

import "bytes"

type document struct {
	Title           string `yaml:"title"`
	PageSize        string
	TableOfContents bool
	Pages           []page
}

func (d document) Generate() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := indexTemplate().Execute(buf, d); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
