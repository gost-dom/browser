package dom

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// ParseDocument parses an HTML or XML document.
func ParseDocument(document Document, reader io.Reader) error {
	return parseIntoDocument(document, reader)
}

// ParseFragment parses an HTML or XML DocumentFragment.
func ParseFragment(doc Document, r io.Reader) (DocumentFragment, error) {
	nodes, err := html.ParseFragment(r, &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	result := doc.CreateDocumentFragment()
	if err == nil {
		for _, child := range nodes {
			iterate(doc, result, child)
		}
	}
	return result, err
}

func parseIntoDocument(doc Document, r io.Reader) error {
	node, err := html.Parse(r)
	if err != nil {
		return err
	}
	iterateChildren(doc, doc, node)
	return nil
}

func convertNS(ns string) string {
	switch ns {
	case "svg":
		return "http://www.w3.org/2000/svg"
	case "math":
		return "http://www.w3.org/1998/Math/MathML"
	default:
		return ns
	}
}

func createElementFromNode(d Document, parent Node, source *html.Node) Element {
	if parent == nil {
		panic("Elements must have a parent")
	}

	var newElm Element
	if source.Namespace == "" {
		newElm = d.CreateElement(source.Data)
	} else {
		newElm = d.CreateElementNS(convertNS(source.Namespace), source.Data)
	}

	for _, a := range source.Attr {
		newElm.SetAttribute(a.Key, a.Val)
	}

	newNode, err := parent.AppendChild(newElm)
	if err != nil {
		panic(err)
	}
	iterateChildren(d, newNode, source)
	return newElm
}

func iterateChildren(d Document, dest Node, source *html.Node) {
	if dest, ok := dest.(interface{ Content() DocumentFragment }); ok {
		iterateChildren(d, dest.Content(), source)
		return
	}
	for child := range source.ChildNodes() {
		iterate(d, dest, child)
	}
}

func iterate(d Document, dest Node, child *html.Node) {
	switch child.Type {
	case html.ElementNode:
		createElementFromNode(d, dest, child)
	case html.TextNode:
		dest.AppendChild(d.CreateTextNode(child.Data))
	case html.DoctypeNode:
		dest.AppendChild(d.CreateDocumentType(child.Data))
	case html.CommentNode:
		dest.AppendChild(d.CreateComment(child.Data))
	default:
		panic(fmt.Sprintf("Node not yet supported: %v", child.Type))
	}
}
