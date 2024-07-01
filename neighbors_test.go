package dom

import (
	"strconv"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var inputOneOriginal = `
<html>
	<head></head>
	<body>
		<nav>
			<p>up</p>
		</nav>
		<main>
			<button>
				<span>start</span>
			</button>
			<div>
				<h3>heading</h3>
				<p>description</p>
			</div>
		</main>
		<footer>
			<p>down</p>
		</section>
	</body>
</html>
`

func TestInitGetNeighbor(t *testing.T) {
	testCases := []struct {
		desc     string
		fn       func(*html.Node) *html.Node
		expected string
	}{
		{
			desc: "GetPrevNeighborNode",
			fn:   GetPrevNeighborNode,
			expected: `
├─#document
│ ├─html
│ │ ├─head (match=6)
│ │ ├─body
│ │ │ ├─nav (match=3)
│ │ │ │ ├─p (match=4)
│ │ │ │ │ ├─#text (match=5) "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span (match=1)
│ │ │ │ │ │ ├─#text (match=2) "start"
│ │ │ │ ├─div
│ │ │ │ │ ├─h3
│ │ │ │ │ │ ├─#text "heading"
│ │ │ │ │ ├─p
│ │ │ │ │ │ ├─#text "description"
│ │ │ ├─footer
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "down"
			`,
		},
		{
			desc: "GetPrevNeighborElement",
			fn:   GetPrevNeighborElement,
			expected: `
├─#document
│ ├─html
│ │ ├─head (match=4)
│ │ ├─body
│ │ │ ├─nav (match=2)
│ │ │ │ ├─p (match=3)
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span (match=1)
│ │ │ │ │ │ ├─#text "start"
│ │ │ │ ├─div
│ │ │ │ │ ├─h3
│ │ │ │ │ │ ├─#text "heading"
│ │ │ │ │ ├─p
│ │ │ │ │ │ ├─#text "description"
│ │ │ ├─footer
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "down"
			`,
		},
		{
			desc: "GetPrevNeighborNodeExcludingOwnChild",
			fn:   GetPrevNeighborNodeExcludingOwnChild,
			expected: `
├─#document
│ ├─html
│ │ ├─head (match=2)
│ │ ├─body
│ │ │ ├─nav (match=1)
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span
│ │ │ │ │ │ ├─#text "start"
│ │ │ │ ├─div
│ │ │ │ │ ├─h3
│ │ │ │ │ │ ├─#text "heading"
│ │ │ │ │ ├─p
│ │ │ │ │ │ ├─#text "description"
│ │ │ ├─footer
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "down"
			`,
		},
		{
			desc: "GetPrevNeighborElementExcludingOwnChild",
			fn:   GetPrevNeighborElementExcludingOwnChild,
			expected: `
├─#document
│ ├─html
│ │ ├─head (match=2)
│ │ ├─body
│ │ │ ├─nav (match=1)
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span
│ │ │ │ │ │ ├─#text "start"
│ │ │ │ ├─div
│ │ │ │ │ ├─h3
│ │ │ │ │ │ ├─#text "heading"
│ │ │ │ │ ├─p
│ │ │ │ │ │ ├─#text "description"
│ │ │ ├─footer
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "down"
			`,
		},

		// - - - - - - - - - - - - - - - - //

		{
			desc: "GetNextNeighborNode",
			fn:   GetNextNeighborNode,
			expected: `
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─nav
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span (match=1)
│ │ │ │ │ │ ├─#text (match=2) "start"
│ │ │ │ ├─div (match=3)
│ │ │ │ │ ├─h3 (match=4)
│ │ │ │ │ │ ├─#text (match=5) "heading"
│ │ │ │ │ ├─p (match=6)
│ │ │ │ │ │ ├─#text (match=7) "description"
│ │ │ ├─footer (match=8)
│ │ │ │ ├─p (match=9)
│ │ │ │ │ ├─#text (match=10) "down"
			`,
		},
		{
			desc: "GetNextNeighborElement",
			fn:   GetNextNeighborElement,
			expected: `
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─nav
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span (match=1)
│ │ │ │ │ │ ├─#text "start"
│ │ │ │ ├─div (match=2)
│ │ │ │ │ ├─h3 (match=3)
│ │ │ │ │ │ ├─#text "heading"
│ │ │ │ │ ├─p (match=4)
│ │ │ │ │ │ ├─#text "description"
│ │ │ ├─footer (match=5)
│ │ │ │ ├─p (match=6)
│ │ │ │ │ ├─#text "down"
			`,
		},
		{
			desc: "GetNextNeighborNodeExcludingOwnChild",
			fn:   GetNextNeighborNodeExcludingOwnChild,
			expected: `
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─nav
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span
│ │ │ │ │ │ ├─#text "start"
│ │ │ │ ├─div (match=1)
│ │ │ │ │ ├─h3
│ │ │ │ │ │ ├─#text "heading"
│ │ │ │ │ ├─p
│ │ │ │ │ │ ├─#text "description"
│ │ │ ├─footer (match=2)
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "down"
			`,
		},
		{
			desc: "GetNextNeighborElementExcludingOwnChild",
			fn:   GetNextNeighborElementExcludingOwnChild,
			expected: `
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─nav
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button (match=0)
│ │ │ │ │ ├─span
│ │ │ │ │ │ ├─#text "start"
│ │ │ │ ├─div (match=1)
│ │ │ │ │ ├─h3
│ │ │ │ │ │ ├─#text "heading"
│ │ │ │ │ ├─p
│ │ │ │ │ │ ├─#text "description"
│ │ │ ├─footer (match=2)
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "down"
			`,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			var replacer = strings.NewReplacer(
				"\n", "",
				"\t", "",
			)
			var inputOne = replacer.Replace(inputOneOriginal)

			doc, err := html.Parse(strings.NewReader(inputOne))
			if err != nil {
				t.Fatal(err)
			}

			button := FindFirstNode(doc, func(node *html.Node) bool {
				return NodeName(node) == "button"
			})

			// - - - - //
			var i int
			node := button
			for node != nil {
				// We record at what point each node was visited...
				node.Attr = append(node.Attr, html.Attribute{
					Key: "match",
					Val: strconv.Itoa(i),
				})
				i++

				node = testCase.fn(node)
			}
			// - - - - //

			r := RenderRepresentation(doc)
			t.Logf("rendered:\n%s\n", r)

			if r != strings.TrimSpace(testCase.expected) {
				t.Error("the representations dont match")
			}
		})
	}
}
