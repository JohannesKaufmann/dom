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

	aNode := dom.FindFirstNode(doc, func(node *html.Node) bool {
		return dom.NodeName(node) == "a"
	})

	next := dom.GetNextNeighborNodeExcludingOwnChild(aNode)
	fmt.Printf("next %s node is %q \n", dom.NodeName(next), next.Data)
}
