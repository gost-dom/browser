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
	Target   string
	Children []*Node
	Parent   *Node
	Type     NodeType
	// For attributes, the fully qualified node name. For elements, the tag name
	Name string

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

// Returns node's nodename
//
// spec: https://dom.spec.whatwg.org/#dom-node-nodename
func (n *Node) NodeName() string {
	switch n.Type {
	case NodeTypeDocumentType, NodeTypeElement, NodeTypeAttribute:
		return n.Name
	case NodeTypeText:
		return "#text"
	case NodeTypeCDataSection:
		return "#cdata-section"
	case NodeTypeProcessingInstruction:
		return n.Target
	case NodeTypeComment:
		return "#comment"
	case NodeTypeDocument:
		return "#document"
	case NodeTypeDocumentFragment:
		return "#document-fragment"
	default:
		panic("gost: Node.NodeName(): invalid node type")
	}
}
