package dom

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestRenderRepresentation_NoAttributes(t *testing.T) {
	input := "<a>\nText\n</a>"
	expected := "├─body\n│ ├─a\n│ │ ├─#text \"\\nText\\n\""

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	body := FindFirstNode(doc, func(node *html.Node) bool {
		return node.DataAtom == atom.Body
	})

	output := RenderRepresentation(body)

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}

func TestRenderRepresentation_MultipleAttributes(t *testing.T) {
	input := `<a href="/page.html" target="_blank" class="button primary">Text</a>`
	expected := `├─body
│ ├─a (href="/page.html" target="_blank" class="button primary")
│ │ ├─#text "Text"`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	body := FindFirstNode(doc, func(node *html.Node) bool {
		return node.DataAtom == atom.Body
	})

	output := RenderRepresentation(body)

	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}

func TestRenderRepresentation_Root(t *testing.T) {
	input := `<img src="/img.png" />`
	expected := strings.TrimSpace(`

#document
├─html
│ ├─head
│ ├─body
│ │ ├─img (src="/img.png")

`)

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	output := RenderRepresentation(doc)

	fmt.Println(output)
	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}
