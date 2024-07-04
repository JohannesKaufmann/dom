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
<h1>Github</h1>

<div>
	<h2 class="repo__name">JohannesKaufmann/dom</h2>

	<nav>
		<h3>Code</h3>
		<h3>Issues</h3>
	</nav>
</div>
	`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	// - - - //

	headingNodes := dom.FindAllNodes(doc, func(node *html.Node) bool {
		name := dom.NodeName(node)
		return dom.NameIsHeading(name)
	})

	nameNode := dom.FindFirstNode(doc, func(node *html.Node) bool {
		return dom.HasClass(node, "repo__name")
	})
	repoName := dom.CollectText(nameNode)

	fmt.Printf("count:%d name:%q\n", len(headingNodes), repoName)
	// count:4 name:"JohannesKaufmann/dom"
}
