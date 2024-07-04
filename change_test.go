package dom

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestRemoveNode(t *testing.T) {
	child := &html.Node{
		Data: "child",
	}
	doc := &html.Node{
		Data: "parent",
	}

	if doc.FirstChild != nil {
		t.Error("expected FirstChild to be nil")
	}
	if child.Parent != nil {
		t.Error("expected Parent to be nil")
	}
	doc.AppendChild(child)
	if doc.FirstChild == nil {
		t.Error("expected FirstChild not to be nil anymore")
	}
	if child.Parent == nil {
		t.Error("expected Parent not to be nil anymore")
	}

	RemoveNode(child)
	if doc.FirstChild != nil {
		t.Error("expected FirstChild to be nil again")
	}
	if child.Parent != nil {
		t.Error("expected Parent to be nil again")
	}

	// Should not crash if run again...
	RemoveNode(child)
	if doc.Parent != nil {
		t.Error("expected Parent to still be nil")
	}
}

func TestReplaceNode(t *testing.T) {
	node1 := &html.Node{
		Data: "original",
	}
	node2 := &html.Node{
		Data: "replacement",
	}

	doc := &html.Node{}
	doc.AppendChild(node1)

	ReplaceNode(node1, node1)
	if doc.FirstChild != node1 {
		t.Error("expected the node1 to still be in place")
	}

	ReplaceNode(node1, node2)
	if doc.FirstChild != node2 {
		t.Error("expected the node2 to take the place")
	}
	if node1.Parent != nil {
		t.Error("expected node1 to not have a parent anymore")
	}
}

func TestUnwrapNode(t *testing.T) {
	child1 := &html.Node{
		Type: html.ElementNode,
		Data: "child1",
	}
	child2 := &html.Node{
		Type: html.ElementNode,
		Data: "child2",
	}
	child2.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "child2.1",
	})

	parent := &html.Node{
		Type: html.ElementNode,
		Data: "parent",
	}
	parent.AppendChild(child1)
	parent.AppendChild(child2)

	root := &html.Node{
		Type: html.ElementNode,
		Data: "root",
	}
	root.AppendChild(parent)

	// - - - - - //

	expectedBefore := strings.TrimSpace(`
root
├─parent
│ ├─child1
│ ├─child2
│ │ ├─child2.1
	`)
	expectedAfter := strings.TrimSpace(`
root
├─child1
├─child2
│ ├─child2.1
	`)

	if RenderRepresentation(root) != expectedBefore {
		t.Error("expected a different initial render")
	}
	if root.FirstChild != parent {
		t.Error("expected the parent to be under root")
	}
	UnwrapNode(parent)
	if root.FirstChild == parent {
		t.Error("expected the parent to not be under root anymore")
	}
	if RenderRepresentation(root) != expectedAfter {
		t.Error("expected a different final render")
	}

	UnwrapNode(root)
	if root.Data != "root" {
		t.Error("expected the root to still be the root")
	}
}

/*
func TestWrapNode(t *testing.T) {
	child1 := &html.Node{
		Type: html.ElementNode,
		Data: "child1",
	}
	child2 := &html.Node{
		Type: html.ElementNode,
		Data: "child2",
	}
	child2.AppendChild(&html.Node{
		Type: html.ElementNode,
		Data: "child2.1",
	})

	parent := &html.Node{
		Type: html.ElementNode,
		Data: "parent",
	}
	parent.AppendChild(child1)
	parent.AppendChild(child2)

	root := &html.Node{
		Type: html.ElementNode,
		Data: "root",
	}
	root.AppendChild(parent)

	// - - - - - //
	wrapper := &html.Node{
		Type: html.ElementNode,
		Data: "wrapper",
	}

	if render(t, root) != "<root><parent><child1></child1><child2><child2.1></child2.1></child2></parent></root>" {
		t.Error("expected a different initial render")
	}
	if root.FirstChild != parent {
		t.Error("expected the parent to be under root")
	}
	WrapNode(parent, wrapper)
	if render(t, root) != "<root><wrapper><parent><child1></child1><child2><child2.1></child2.1></child2></parent></wrapper></root>" {
		t.Error("expected a different final render")
	}
	if root.FirstChild != wrapper {
		t.Error("expected the wrapper to be under root")
	}
}
*/
