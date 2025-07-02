package dom

type ChildNode interface {
	Remove()
}

type childNode struct {
	node
}

func newChildNode(ownerDocument Document) childNode { return childNode{newNode(ownerDocument)} }

func (n *childNode) Remove() {
	s := n.self
	s.Parent().RemoveChild(s)
}
