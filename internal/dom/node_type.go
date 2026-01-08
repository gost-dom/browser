package dom

import "strconv"

type NodeType int

const (
	NodeTypeElement               NodeType = 1
	NodeTypeAttribute             NodeType = 2
	NodeTypeText                  NodeType = 3
	NodeTypeCDataSection          NodeType = 4
	NodeTypeProcessingInstruction NodeType = 7
	NodeTypeComment               NodeType = 8
	NodeTypeDocument              NodeType = 9
	NodeTypeDocumentType          NodeType = 10
	NodeTypeDocumentFragment      NodeType = 11
)

// CanHaveChildren returns true for note types that allow child nodes
func (t NodeType) CanHaveChildren() bool {
	switch t {
	case NodeTypeElement:
		return true
	case NodeTypeDocument:
		return true
	case NodeTypeDocumentFragment:
		return true
	default:
		return false
	}
}

// IsCharacterDataNode returns true the 4 node types that are [Characterdata]
// nodes.
//
// [Characterdata]: https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
func (t NodeType) IsCharacterDataNode() bool {
	switch t {
	case NodeTypeText:
		return true
	case NodeTypeCDataSection:
		return true
	case NodeTypeComment:
		return true
	case NodeTypeProcessingInstruction:
		return true
	default:
		return false
	}
}

// CanBeAChild returns true for node types that are allowed as children. Node,
// special rules may apply, such as a document node can only contain one child
// element.
func (t NodeType) CanBeAChild() bool {
	if t.IsCharacterDataNode() {
		return true
	}
	switch t {
	case NodeTypeDocumentFragment:
		return true
	case NodeTypeDocumentType:
		return true
	case NodeTypeElement:
		return true
	default:
		return false
	}
}

// String returns name of the node type. For invalid values, a string
// representation of the integer value is returned.
func (t NodeType) String() string {
	switch t {
	case NodeTypeElement:
		return "Element"
	case NodeTypeAttribute:
		return "Attribute"
	case NodeTypeText:
		return "Text"
	case NodeTypeCDataSection:
		return "CDataSection"
	case NodeTypeProcessingInstruction:
		return "ProcessingInstruction"
	case NodeTypeComment:
		return "Comment"
	case NodeTypeDocument:
		return "Document"
	case NodeTypeDocumentType:
		return "DocumentType"
	case NodeTypeDocumentFragment:
		return "DocumentFragment"
	default:
		return strconv.Itoa(int(t))
	}
}
