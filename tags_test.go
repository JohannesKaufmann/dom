package dom

import (
	"testing"

	"golang.org/x/net/html"
)

func TestNodeName(t *testing.T) {
	runs := []struct {
		name string
		node *html.Node
	}{
		{
			name: "",
			node: nil,
		},
		{
			name: "",
			node: &html.Node{
				Type: 10,
			},
		},
		{
			name: "#error",
			node: &html.Node{
				Type: html.ErrorNode,
			},
		},
		{
			name: "#text",
			node: &html.Node{
				Type: html.TextNode,
				Data: "some boring text",
			},
		},

		{
			name: "#document",
			node: &html.Node{
				Type: html.DocumentNode,
			},
		},
		{
			name: "#comment",
			node: &html.Node{
				Type: html.CommentNode,
			},
		},
		{
			name: "html",
			node: &html.Node{
				Type: html.DoctypeNode,
				// E.g. for `<!DOCTYPE html>` it would be "html"
				Data: "html",
			},
		},
		// - - - - - - - - - - //
		{
			name: "div",
			node: &html.Node{
				Type: html.ElementNode,
				Data: "div",
			},
		},
		{
			name: "a",
			node: &html.Node{
				Type: html.ElementNode,
				Data: "a",
			},
		},
	}
	for _, run := range runs {
		t.Run(run.name, func(t *testing.T) {
			output := NodeName(run.node)
			if output != run.name {
				t.Errorf("expected '%s' but got '%s'", run.name, output)
			}
		})
	}
}

func TestIsInlineNode(t *testing.T) {
	if out := IsInlineNode("strong"); out != true {
		t.Error("expected different output")
	}

	if out := IsInlineNode("div"); out != false {
		t.Error("expected different output")
	}
	if out := IsInlineNode("magic"); out != false {
		t.Error("expected different output")
	}
}

func TestIsBlockNode(t *testing.T) {
	if out := IsBlockNode("div"); out != true {
		t.Error("expected different output")
	}

	if out := IsBlockNode("strong"); out != false {
		t.Error("expected different output")
	}
	if out := IsBlockNode("magic"); out != false {
		t.Error("expected different output")
	}
}
