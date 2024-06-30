package dom

import "golang.org/x/net/html"

func initGetNeighbor(
	firstChildFunc func(node *html.Node) *html.Node,
	prevNextFunc func(node *html.Node) *html.Node,
	goUpUntilFunc func(node *html.Node) bool,
) func(node *html.Node) *html.Node {

	return func(node *html.Node) *html.Node {
		// First look at the children
		if child := firstChildFunc(node); child != nil {
			return child
		}

		// Otherwise my prev/next sibling
		if sibling := prevNextFunc(node); sibling != nil {
			return sibling
		}

		for {
			// Finally, continously go upwards until we find an element with a sibling
			node = node.Parent
			if node == nil {
				// We reached the top
				return nil
			}

			if goUpUntilFunc(node) {
				// Don't go too far up...
				return nil
			}

			sibling := prevNextFunc(node)
			if sibling != nil {
				return sibling
			}
		}
	}
}

func FindFirstNode(startNode *html.Node, matchFn func(node *html.Node) bool) *html.Node {
	nextFunc := initGetNeighbor(
		FirstChildNode,
		NextSiblingNode,
		func(node *html.Node) bool {
			// We should not get higher up than the startNode...
			return node == startNode
		},
	)

	child := startNode.FirstChild
	for child != nil {
		if matchFn(child) {
			return child
		}

		child = nextFunc(child)
	}
	return nil
}
