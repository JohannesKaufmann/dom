package dom

import (
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
