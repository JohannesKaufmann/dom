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
<p>
	<i>Javascript</i>
	<i class="lang__old">PHP</i>
	<i>Golang</i>
<p>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	// - - - - - //

	italicNodes := dom.FindAllNodes(doc, func(node *html.Node) bool {
		return dom.NodeName(node) == "i"
	})

	for _, node := range italicNodes {
		if dom.HasClass(node, "lang__old") {
			newNode := &html.Node{
				Type: html.TextNode,
				Data: "🪦",
			}
			dom.ReplaceNode(node, newNode)
		}
	}

	// - - - - - //

	emptyTextNodes := dom.FindAllNodes(doc, func(node *html.Node) bool {
		name := dom.NodeName(node)
		text := dom.CollectText(node)

		return name == "#text" && strings.TrimSpace(text) == ""
	})

	for _, node := range emptyTextNodes {
		dom.RemoveNode(node)
	}

	// - - - - - //

	fmt.Println(dom.RenderRepresentation(doc))
}
