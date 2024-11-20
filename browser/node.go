package browser

import (
	"fmt"

	"golang.org/x/net/html"
)

type Node interface {
	EventTarget
	AppendChild(node Node) Node
	ChildNodes() []Node
	Connected() bool
	NodeName() string
	Parent() Node
	// unexported
	createHtmlNode() *html.Node
	// toHtmlNode(Node, map[*html.Node]Node) *html.Node
	populateNodeMap(map[*html.Node]Node)
	setParent(node Node)
}

type node struct {
	eventTarget
	childNodes []Node
	name       string
	htmlNode   *html.Node
	parent     Node
}

func newNode(htmlNode *html.Node) node {
	return node{newEventTarget(), []Node{}, htmlNode.Data, htmlNode, nil}
}

func (parent *node) AppendChild(child Node) Node {
	parent.childNodes = append(parent.childNodes, child)
	return child
}

func (n *node) ChildNodes() []Node { return n.childNodes }

func (n *node) Parent() Node { return n.parent }

func (n *node) setParent(parent Node) { n.parent = parent }

func (n *node) Connected() (result bool) {
	if n.parent != nil {
		result = n.parent.Connected()
	}
	return
}

func (n *node) wrappedNode() *html.Node {
	return n.htmlNode
}

func (n *node) NodeName() string {
	return "#node"
}

// Temporary hack while the code depends on the html.Node data for e.g., CSS
// selectors.
//
// NOTE: Because Go doesn't have "virtual functions", if you need to be able to
// interact with the element of the correct subtype, that subtype needs to
// implement this function as well. E.g., it's implemented on the Element type
// too, as we need to have Element properties. If not, only the embedded value
// of the Element is stored in the map
func (n *node) populateNodeMap(m map[*html.Node]Node) {
	m[n.htmlNode] = n
	for _, c := range n.childNodes {
		c.populateNodeMap(m)
	}
}

func (n *node) createHtmlNode() *html.Node {
	panic(fmt.Sprintf("You must implement this on the specialised node: %v %v", n.htmlNode.Type, n))
}

type NodeIterator struct{ Node }

func (n NodeIterator) toHtmlNode(m map[*html.Node]Node) *html.Node {
	htmlNode := n.Node.createHtmlNode()
	if m != nil {
		m[htmlNode] = n.Node
	}
	for _, child := range n.ChildNodes() {
		htmlNode.AppendChild(NodeIterator{child}.toHtmlNode(m))
	}
	return htmlNode
}
