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
	if ownerDocument != nil && ownerDocument.Type != NodeTypeDocument {
		panic(fmt.Sprintf("Invalid owner document: %v", ownerDocument.Type))
	}
	return &Node{OwnerDocument: ownerDocument, Type: nodeType}
}
