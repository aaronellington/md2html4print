package main

import (
	"fmt"
	"os"

	"github.com/aaronellington/md2html4print/src/md2html4print"
)

func main() {
	userData, err := md2html4print.ParseDocument(os.DirFS("."))
	if err != nil {
		panic(err)
	}

	html, err := userData.Generate()
	if err != nil {
		panic(err)

	}

	fmt.Println(string(html))
}
