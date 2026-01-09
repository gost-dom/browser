package dom

import "github.com/gost-dom/browser/internal/entity"

type NodeList interface {
	entity.Components
	Length() int

	// Item returns the node with the specified zero-based index. If the index
	// is out of range, the function returns nil.
	Item(index int) Node

	All() []Node
	setNodes([]Node)
}

type nodeList struct {
	entity.Entity
	nodes *[]Node
}

func (l *nodeList) empty() bool { return l == nil || l.nodes == nil }

func (l *nodeList) Length() int {
	if l.empty() {
		return 0
	}
	return len(*l.nodes)
}

func (l *nodeList) Item(index int) Node {
	if index >= l.Length() || index < 0 {
		return nil
	}
	return (*l.nodes)[index]
}

func (l *nodeList) All() []Node {
	if l.empty() {
		return nil
	}
	return *l.nodes
}
func (l *nodeList) setNodes(nodes []Node) {
	if l == nil {
		panic("nodeList.setNodes: nodelist is nil")
	}
	if l.nodes == nil {
		panic("nodeList.setNodes: nodelist.nodes is nil")
	}
	*l.nodes = nodes
}
