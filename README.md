# dom

[![Go Reference](https://pkg.go.dev/badge/github.com/JohannesKaufmann/dom.svg)](https://pkg.go.dev/github.com/JohannesKaufmann/dom)

Helper functions for "net/html" that make it easier to interact with `*html.Node`.

## Installation

```bash
go get -u github.com/JohannesKaufmann/dom
```

> [!NOTE]
> This "dom" libary was developed for the needs of the [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) library.
> That beeing said, please submit any functions that you need.

## Usage

```go
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

	fmt.Println("href:", dom.GetAttributeOr(firstLink, "href", ""))
}
```

## Node vs Element

The naming scheme in this library is:

- "Node" means `*html.Node{}`
  - This means _any_ node in the tree of nodes.
- "Element" means `*html.Node{Type: html.ElementNode}`
  - This means _only_ nodes with the type of `ElementNode`. For example `<p>`, `<span>`, `<a>`, ... but not `#text`, `<!--comment-->`, ...

For most functions, there are two versions. For example:

- `FirstChildNode()` and `FirstChildElement()`
- `GetChildNodes()` and `GetChildElements()`
- ...

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/JohannesKaufmann/dom.svg)](https://pkg.go.dev/github.com/JohannesKaufmann/dom)

### Attributes

- `GetAttribute` and `GetAttributeOr`

---

### Children & Siblings

- `AllChildNodes` and `AllChildElements`

- `FirstChildNode` and `FirstChildElement`

- `PrevSiblingNode` and `PrevSiblingElement`

- `NextSiblingNode` and `NextSiblingElement`

### Find Nodes

```go
firstParagraph := dom.FindFirstNode(doc, func(node *html.Node) bool {
    return dom.NodeName(node) == "p"
})
// *html.Node


allParagraphs := dom.FindAllNodes(doc, func(node *html.Node) bool {
    return dom.NodeName(node) == "p"
})
// []*html.Node
```

---

### Get next/previous neighbors

What is special about this? The order!

If you are somewhere in the DOM, you can call `GetNextNeighborNode` to get the next node, even if it is _further up_ the tree. The order is the same as you would see the elements in the DOM.

```go
node := startNode
for node != nil {
    fmt.Println(dom.NodeName(node))

    node = dom.GetNextNeighborNode(node)
}
```

If we start the `for` loop at the `<button>` and repeatedly call `GetNextNeighborNode` this would be the _order_ that the nodes are _visited_.

```text
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─nav
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "up"
│ │ │ ├─main
│ │ │ │ ├─button   *️⃣
│ │ │ │ │ ├─span  0️⃣
│ │ │ │ │ │ ├─#text "start"  1️⃣
│ │ │ │ ├─div  2️⃣
│ │ │ │ │ ├─h3  3️⃣
│ │ │ │ │ │ ├─#text "heading"  4️⃣
│ │ │ │ │ ├─p  5️⃣
│ │ │ │ │ │ ├─#text "description"  6️⃣
│ │ │ ├─footer  7️⃣
│ │ │ │ ├─p  8️⃣
│ │ │ │ │ ├─#text "down"  9️⃣
```

If you only want to visit the ElementNode's (and skip the `#text` Nodes) you can use `GetNextNeighborElement` instead.

If you want to skip the children you can use `GetNextNeighborNodeExcludingOwnChild`. In the example above, when starting at the `<button>` the next node would be the `<div>`.

The same functions also exist for the previous nodes, e.g. `GetPrevNeighborNode`.

---

### RemoveNode

### ReplaceNode

### UnwrapNode

```text
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─article   *️⃣
│ │ │ │ ├─h3
│ │ │ │ │ ├─#text "Heading"
│ │ │ │ ├─p
│ │ │ │ │ ├─#text "short description"
```

If we take the input above and run `UnwrapNode(articleNode)` we can "unwrap" the `<article>`. That means removing the `<article>` while _keeping_ the children (`<h3>` and `<p>`).

```text
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─h3
│ │ │ │ ├─#text "Heading"
│ │ │ ├─p
│ │ │ │ ├─#text "short description"
```

---

### RenderRepresentation

```go
import (
	"fmt"
	"log"
	"strings"

	"github.com/JohannesKaufmann/dom"
	"golang.org/x/net/html"
)

func main() {
	input := `<a href="/about">Read More</a>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dom.RenderRepresentation(doc))
}
```

The tree representation helps to visualize the tree-structure of the DOM.
And the `#text` nodes stand out.

> [!TIP]
> This function could be useful for debugging & testcases.

```text
├─#document
│ ├─html
│ │ ├─head
│ │ ├─body
│ │ │ ├─a (href=/about)
│ │ │ │ ├─#text "Read More"
```

While the normal "net/html" [`Render()`](https://pkg.go.dev/golang.org/x/net/html#Render) function would have produced this:

```
<html><head></head><body><a href="/about">Read More</a></body></html>
```
