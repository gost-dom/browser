package dom

import "golang.org/x/net/html"

type DocumentType interface {
	Node
	ChildNode
	Name() string
}

type documentType struct {
	childNode
	name string
}

func NewDocumentType(name string, ownerDocument Document) DocumentType {
	result := &documentType{newChildNode(ownerDocument), name}
	result.SetSelf(result)
	return result
}

func (t *documentType) Name() string       { return t.name }
func (t *documentType) NodeType() NodeType { return NodeTypeDocumentType }

func (t *documentType) CloneNode(deep bool) Node {
	return NewDocumentType(t.name, t.OwnerDocument())
}

func (t *documentType) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DoctypeNode,
		Data: t.name,
	}
}
