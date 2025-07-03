package dom

import "golang.org/x/net/html"

type DocumentFragment interface {
	RootNode
}

type documentFragment struct {
	rootNode
}

func NewDocumentFragment(ownerDocument Document) DocumentFragment {
	result := &documentFragment{newRootNode(ownerDocument)}
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

func (d *documentFragment) NodeType() NodeType { return NodeTypeDocumentFragment }

func (d *documentFragment) NodeName() string { return "#document-fragment" }
