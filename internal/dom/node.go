package dom

import "github.com/gost-dom/browser/internal/entity"

type Node struct {
	entity.Entity
	// revision of the node is incremented on any change. Used by
	// LiveHtmlCollection to check if a node has been changed.
	rev      int
	Children []*Node
	Parent   *Node
}

func NewNode() *Node {
	return &Node{}
}
