package dom

import (
	intdom "github.com/gost-dom/browser/internal/dom"
	"golang.org/x/net/html"
)

type DocumentFragment interface {
	RootNode
}

type documentFragment struct {
	node
	parentNode
	rootNode
}

func NewDocumentFragment(ownerDocument Document) DocumentFragment {
	result := &documentFragment{node: newNode(ownerDocument, intdom.NodeTypeDocumentFragment)}
	result.parentNode = parentNode{&result.node}
	result.rootNode = rootNode{&result.node}
	result.SetSelf(result)
	return result
}

func (f *documentFragment) cloneNode(doc Document, deep bool) Node {
	result := NewDocumentFragment(doc)
	if deep {
		result.Append(f.cloneChildren()...)
	}
	return result
}

func (d *documentFragment) GetElementById(id string) Element {
	return rootNodeHelper{d}.GetElementById(id)
}

func (d *documentFragment) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DocumentNode,
	}
}

func (d *documentFragment) NodeName() string { return "#document-fragment" }
