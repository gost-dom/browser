package dom

import (
	intdom "github.com/gost-dom/browser/internal/dom"
	"golang.org/x/net/html"
)

type DocumentType interface {
	Node
	ChildNode
	Name() string
}

type documentType struct {
	node
	childNode
}

func NewDocumentType(name string, ownerDocument Document) DocumentType {
	res := &documentType{node: newNode(ownerDocument, intdom.NodeTypeDocumentType)}
	res.Node.Name = name
	res.childNode = childNode{&res.node}
	res.SetSelf(res)
	return res
}

func (t *documentType) Name() string { return t.Node.Name }

func (t *documentType) cloneNode(doc Document, deep bool) Node {
	return NewDocumentType(t.Name(), doc)
}

func (t *documentType) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DoctypeNode,
		Data: t.Name(),
	}
}
