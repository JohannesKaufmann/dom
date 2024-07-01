package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/JohannesKaufmann/dom"
	"golang.org/x/net/html"
)

func main() {
	input := `
	<ul>
		<li><a href="github.com/JohannesKaufmann/dom">dom</a></li>
		<li><a href="github.com/JohannesKaufmann/html-to-markdown">html-to-markdown</a></li>
	</ul>
	`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	// - - - //

	firstLink := dom.FindFirstNode(doc, func(node *html.Node) bool {
		return dom.NodeName(node) == "a"
	})

	fmt.Println("href:", dom.GetAttributeOr(firstLink, "href", ""))
}
