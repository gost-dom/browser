package dom

type ChildNode interface {
	Remove()
}

type childNode struct {
	node *node
}

func (n childNode) Remove() {
	s := n.node.self()
	if s == nil {
		panic("NIL SELF")
	}
	if parent := s.ParentNode(); parent != nil {
		parent.RemoveChild(s)
	}
}
