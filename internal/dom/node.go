package dom

import (
	"fmt"

	"github.com/gost-dom/browser/internal/entity"
)

type Node struct {
	entity.Entity
	// revision of the node is incremented on any change. Used by
	// LiveHtmlCollection to check if a node has been changed.
	rev      int
	Children []*Node
	Parent   *Node
	Type     NodeType

	OwnerDocument *Node
}

func NewNode(ownerDocument *Node, nodeType NodeType) *Node {
	var node = Node{Type: nodeType}
	node.SetOwnerDocument(ownerDocument)
	return &node
}

func (n *Node) SetOwnerDocument(ownerDocument *Node) {
	if n.OwnerDocument == ownerDocument {
		return
	}
	if ownerDocument != nil && ownerDocument.Type != NodeTypeDocument {
		panic(fmt.Sprintf("Invalid owner document: %v", ownerDocument.Type))
	}
	n.OwnerDocument = ownerDocument
	for _, c := range n.Children {
		c.SetOwnerDocument(ownerDocument)
	}
}

func (n *Node) SetParent(parent *Node) {
	if parent == nil {
		n.Parent = nil
	} else {
		n.SetOwnerDocument(parent.OwnerDocument)
		n.Parent = parent
	}
}
