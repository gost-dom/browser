package dom

// parentNode implements the functions defined in the [ParentNode] IDL Mixin
// interface.
type parentNode struct {
	node *node
}

func newParentNode(n *node) parentNode {
	return parentNode{n}
}

func (n parentNode) Append(nodes ...Node) (err error) {
	if node := n.nodeOfNodes(nodes); node != nil {
		_, err = n.node.self.InsertBefore(node, nil)
	}
	return
}

func (n parentNode) Prepend(nodes ...Node) (err error) {
	if node := n.nodeOfNodes(nodes); node != nil {
		_, err = n.node.self.InsertBefore(node, n.node.FirstChild())
	}
	return
}

func (n parentNode) ReplaceChildren(nodes ...Node) (err error) {
	if node := n.nodeOfNodes(nodes); node != nil {
		if err = n.node.assertCanAddNode(node); err == nil {
			for c := n.node.FirstChild(); c != nil; c = n.node.FirstChild() {
				n.node.RemoveChild(c)
			}
			n.node.self.InsertBefore(node, nil)
		}
	}
	return
}

func (n parentNode) nodeOfNodes(nodes []Node) Node {
	switch len(nodes) {
	case 0:
		return nil
	case 1:
		return nodes[0]
	default:
		fragment := n.node.getSelf().OwnerDocument().CreateDocumentFragment()
		for _, n := range nodes {
			fragment.AppendChild(n)
		}
		return fragment
	}
}

func (n parentNode) Children() HTMLCollection {
	return newHtmlCollection(n.node.getSelf())
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
