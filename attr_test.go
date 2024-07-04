package dom

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestGetAttribute(t *testing.T) {
	node := &html.Node{
		Attr: []html.Attribute{
			{
				Key: "previouskey",
				Val: "previousval",
			},
			{
				Key: "mykey",
				Val: "myval",
			},
		},
	}

	actual, found := GetAttribute(node, "mykey")
	if !found {
		t.Error("expected found to be true")
	}
	if actual != "myval" {
		t.Error("expected different value")
	}

	actual, found = GetAttribute(node, "unknownkey")
	if found {
		t.Error("expected found to be false")
	}
	if actual != "" {
		t.Error("expected empty value")
	}
}

func TestGetAttributeOr(t *testing.T) {
	node := &html.Node{
		Attr: []html.Attribute{
			{
				Key: "previouskey",
				Val: "previousval",
			},
			{
				Key: "mykey",
				Val: "myval",
			},
			{
				Key: "nextkey",
				Val: "nextval",
			},
		},
	}

	actual := GetAttributeOr(node, "mykey", "myfallback")
	if actual != "myval" {
		t.Error("expected different value")
	}

	actual = GetAttributeOr(node, "unknownkey", "myfallback")
	if actual != "myfallback" {
		t.Error("expected different fallback value")
	}
}

func TestGetClasses(t *testing.T) {
	node1 := &html.Node{
		Type: html.ElementNode,
		Data: "h1",
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: " form form--theme-xmas  form--simple",
			},
		},
	}

	classes := GetClasses(node1)
	if !reflect.DeepEqual(classes, []string{"form", "form--theme-xmas", "form--simple"}) {
		t.Error("the slice of classes dont match")
	}

	node2 := &html.Node{
		Type: html.ElementNode,
		Data: "h1",
		Attr: []html.Attribute{},
	}
	classes = GetClasses(node2)
	if len(classes) != 0 {
		t.Error("expected no classes")
	}
}

func TestHasID(t *testing.T) {
	node1 := &html.Node{
		Type: html.ElementNode,
		Data: "h1",
		Attr: []html.Attribute{
			{
				Key: "id",
				Val: " city__name ",
			},
		},
	}

	if HasID(node1, "city__name") != true {
		t.Error("expected different output")
	}
	if HasID(node1, "city__image") != false {
		t.Error("expected different output")
	}

	node2 := &html.Node{
		Type: html.ElementNode,
		Data: "h1",
		Attr: []html.Attribute{},
	}
	if HasID(node2, "city__name") != false {
		t.Error("expected different output")
	}
}

func TestHasClass(t *testing.T) {
	node := &html.Node{
		Type: html.ElementNode,
		Data: "h1",
		Attr: []html.Attribute{
			{
				Key: "class",
				Val: " form form--theme-xmas  form--simple",
			},
		},
	}

	if HasClass(node, "form--theme-xmas") != true {
		t.Error("expected different output")
	}

	if HasClass(node, "xmas") != false {
		t.Error("expected different output")
	}
}

func TestCollectText(t *testing.T) {
	input := `
	<h2 class="article__title">Hello <span>world</span></h2>
	<p>Some description</p>
	`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	heading := FindFirstNode(doc, func(node *html.Node) bool {
		return HasClass(node, "article__title")
	})

	expected := "Hello world"
	output := CollectText(heading)
	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}
