package dom

import "golang.org/x/net/html"

type DocumentType interface {
	Node
	ChildNode
	Name() string
	PublicId() string
	SystemId() string
}

type documentType struct {
	node
	childNode
	name     string
	publicID string
	systemID string
}

func NewDocumentType(
	qualifiedName, publicID, systemID string,
	ownerDocument Document,
) DocumentType {
	res := &documentType{
		node:     newNode(ownerDocument),
		name:     qualifiedName,
		publicID: publicID,
		systemID: systemID,
	}
	res.childNode = childNode{&res.node}
	res.SetSelf(res)
	return res
}

func (t *documentType) Name() string       { return t.name }
func (t *documentType) PublicId() string   { return t.publicID }
func (t *documentType) SystemId() string   { return t.systemID }
func (t *documentType) NodeType() NodeType { return NodeTypeDocumentType }

func (t *documentType) IsEqualNode(n Node) bool {
	other, ok := n.(*documentType)
	return ok && other.name == t.name && other.publicID == t.publicID &&
		other.systemID == t.systemID
}

func (t *documentType) cloneNode(doc Document, deep bool) Node {
	return NewDocumentType(t.name, t.publicID, t.systemID, doc)
}

func (t *documentType) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DoctypeNode,
		Data: t.name,
	}
}
