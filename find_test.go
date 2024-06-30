package dom

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestContainsNode(t *testing.T) {
	node := parse(t, `<root><start><span><target></target></span><span><target></target></span></start></root><next></next>`, "root")

	var called1 int
	res1 := ContainsNode(node.FirstChild, func(n *html.Node) bool {
		if n.Data == "next" {
			t.Error("the next node should not have been visited")
		}

		called1++
		return n.Data == "target"
	})
	if res1 != true {
		t.Error("expected true")
	}
	if called1 != 2 {
		t.Error("expected fn to be called 2 times")
	}

	var called2 int
	res2 := ContainsNode(node.FirstChild, func(n *html.Node) bool {
		if n.Data == "next" {
			t.Error("the next node should not have been visited")
		}

		called2++
		return n.Data == "else"
	})
	if res2 != false {
		t.Error("expected false")
	}
	if called2 != 4 {
		t.Error("expected fn to be called 4 times")
	}
}

func TestFindFirstNode(t *testing.T) {
	input := `
<html>
	<head></head>
	<body>
		<article>
			<div>
				<h3>Heading</h3>
				<p>short description</p>
			</div>
		</article>

		<section>
			<h4>Heading</h4>
			<p>another description</p>
		</section>
	</body>
</html>
	`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	article := FindFirstNode(doc, func(node *html.Node) bool {
		return NodeName(node) == "article"
	})
	if article == nil || article.Data != "article" {
		t.Error("got different node")
	}

	h3 := FindFirstNode(article, func(node *html.Node) bool {
		return NodeName(node) == "h3"
	})
	if h3 == nil || h3.Data != "h3" {
		t.Error("got different node")
	}

	h4 := FindFirstNode(article, func(node *html.Node) bool {
		return NodeName(node) == "h4"
	})
	if h4 != nil {
		t.Error("expected nil node")
	}
}

func TestFindAllNodes(t *testing.T) {
	input := `
<html>
	<head></head>
	<body>
		<article>
			<div>
				<h3>Heading</h3>
				<p>short description</p>
			</div>
		</article>

		<section>
			<h4>Heading</h4>
			<p>another description</p>
		</section>
	</body>
</html>
	`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	paragraphs := FindAllNodes(doc, func(node *html.Node) bool {
		return NodeName(node) == "p"
	})
	if len(paragraphs) != 2 {
		t.Error("expected 2 nodes")
	}
	if paragraphs[0].Data != "p" || paragraphs[1].Data != "p" {
		t.Error("expected paragraph nodes")
	}
}
