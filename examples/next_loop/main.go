package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/JohannesKaufmann/dom"
	"golang.org/x/net/html"
)

func main() {
	input := `<p>The <i><a href="/about">library</a></i> is amazing<p>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dom.RenderRepresentation(doc))

	// - - - //

	node := doc
	for node != nil {
		fmt.Println(dom.NodeName(node))

		node = dom.GetNextNeighborNode(node)
	}

	// #document
	// html
	// head
	// body
	// p
	// #text
	// i
	// a
	// #text
	// #text
	// p
}
