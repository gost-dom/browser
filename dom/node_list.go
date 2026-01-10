package dom

import (
	"iter"
	"slices"

	intdom "github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/entity"
)

// NodeList corresponds to the NodeList IDL interface.
//
// see also: https://developer.mozilla.org/en-US/docs/Web/API/NodeList
type NodeList interface {
	entity.Components
	Length() int

	// Item returns the node with the specified zero-based index. If the index
	// is out of range, the function returns nil.
	Item(index int) Node

	All() iter.Seq[Node]
}

// nodeList wraps a slice of Node values, and implements the NodeList interface
// for it. It uses a pointer value to support live collections.
type nodeList struct {
	entity.Entity
	nodes *[]*intdom.Node
}

func newStaticNodeList(n []*intdom.Node) NodeList {
	data := slices.Clone(n)
	return &nodeList{nodes: &data}
}

func newDynamicNodeList(n *[]*intdom.Node) NodeList {
	return &nodeList{nodes: n}
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
	return getNode((*l.nodes)[index])
}

func (l *nodeList) All() iter.Seq[Node] {
	return func(yield func(Node) bool) {
		for i := range l.Length() {
			if !yield(l.Item(i)) {
				return
			}
		}
	}
}
