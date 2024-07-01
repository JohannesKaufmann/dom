package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/JohannesKaufmann/dom"
	"golang.org/x/net/html"
)

func main() {
	input := `<a href="/about">Read More</a>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	// - - - //

	fmt.Println(dom.RenderRepresentation(doc))

	// - - - //

	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
