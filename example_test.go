package dom_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/JohannesKaufmann/dom"
	"golang.org/x/net/html"
)

func ExampleFindFirstNode() {
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

	fmt.Println(dom.GetAttributeOr(firstLink, "href", ""))
	// Output: github.com/JohannesKaufmann/dom
}
