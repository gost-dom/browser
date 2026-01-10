package dom

import (
	"slices"

	intdom "github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/css"
)

// ElementParent is a type alias for ParentNode that allows it to be embedded in
// other interfaces and types, without conflicting with the ParentNode method on
// [Node]
type ElementParent = ParentNode

type NonElementParentNode interface {
	GetElementById(string) Element
}

// parentNode implements the functions defined in the [ParentNode] IDL Mixin
// interface.
type parentNode struct {
	node *node
}

func (n parentNode) Append(nodes ...Node) (err error) {
	if node := n.collapseNodes(nodes); node != nil {
		_, err = n.node.self().InsertBefore(node, nil)
	}
	return
}

func (n parentNode) Prepend(nodes ...Node) (err error) {
	if node := n.collapseNodes(nodes); node != nil {
		_, err = n.node.self().InsertBefore(node, n.node.FirstChild())
	}
	return
}

func (n parentNode) ReplaceChildren(nodes ...Node) (err error) {
	node := n.collapseNodes(nodes)
	if err = n.node.assertCanAddNode(node); err == nil {
		return n.node.replaceNodes(0, len(n.node.ptr().Children), node)
	}
	return
}

func (n parentNode) collapseNodes(nodes []Node) Node {
	return collapseNodes(n.node.OwnerDocument(), nodes)
}

// nodeOfNodes creates a single node from a list of nodes. The return value
// depends on the length of the list
//
// - If the list is empty it returns nil
// - If the list contains a single element, it returns the element
// - If the list contains multiple elements, it returns a [DocumentFragment]
//
// This helps implement the WHAT cookbook rules, which explains that multiple
// arguments to [ParentNode.Append] or [ParentNode.ReplaceChildren] should treat
// multiple children as a document fragment. This makes the valid child
// validation easier, as the same rule applies for single vs. multiple elements.
func collapseNodes(owner Document, nodes []Node) Node {
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		return nodes[0]
	default:
		fragment := owner.CreateDocumentFragment()
		for _, n := range nodes {
			fragment.AppendChild(n)
		}
		return fragment
	}
}

// expandFragment is the opposite of [collapseNodes], it returns a single [Node]
// into a list of nodes.
//
// - A nil value returns a nil slice (empty slice)
// - A [DocumentFragment] returns its children
// - Any other node returns a single-element slice of itself.
func expandNode(node Node) []*intdom.Node {
	if node == nil {
		return nil
	}

	if _, ok := node.(DocumentFragment); ok {
		return slices.Clone(node.ptr().Children)
	} else {
		return []*intdom.Node{node.ptr().Node}
	}
}

func (n parentNode) Children() HTMLCollection {
	return newHtmlCollection(n.node.self())
}

func (n parentNode) FirstElementChild() Element {
	return n.Children().Item(0)
}

func (n parentNode) LastElementChild() Element {
	c := n.Children()
	return n.Children().Item(c.Length() - 1)
}

func (f parentNode) ChildElementCount() int {
	return len(f.node.childElements())
}

func (n parentNode) QuerySelector(pattern string) (Element, error) {
	n.node.Logger().Debug("parentNode.QuerySelector", "pattern", pattern)
	nodes, err := n.QuerySelectorAll(pattern)
	if err != nil {
		return nil, err
	}
	// TODO, it should be a list of elements, not nodes, then the cast, and
	// error isn't necessary
	result := nodes.Item(0)
	element, _ := result.(Element)
	return element, nil
}

func (n parentNode) QuerySelectorAll(pattern string) (NodeList, error) {
	sel, err := css.Parse(pattern)
	if err != nil {
		return nil, err
	}
	htmlNode, m := toHtmlNodeAndMap(n.node.self())

	nodes := sel.Select(htmlNode)
	result := make([]*intdom.Node, len(nodes))
	for i, node := range nodes {
		result[i] = m[node].ptr().Node
	}
	return &nodeList{nodes: &result}, nil
}
