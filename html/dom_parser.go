package html

import (
	"fmt"
	"io"

	"github.com/gost-dom/browser/dom"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type domParser struct{}

// Parses a HTML or XML from an [io.Reader] instance. The parsed nodes will
// have reference the window, e.g. letting events bubble to the window itself.
// The document pointer will be replaced by the created document.
//
// The document is updated using a pointer rather than returned as a value, as
// parseing process can e.g. execute script tags that require the document to
// be set on the window _before_ the script is executed.
func (p domParser) ParseReader(window Window, document *dom.Document, reader io.Reader) error {
	*document = newHTMLDocument(window)
	return parseIntoDocument(*document, reader)
}

func (p domParser) ParseFragment(
	document dom.Document,
	reader io.Reader,
) (dom.DocumentFragment, error) {
	return ParseFragment(document, reader)
}

func ParseFragment(
	ownerDocument dom.Document,
	reader io.Reader,
) (dom.DocumentFragment, error) {
	nodes, err := html.ParseFragment(reader, &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	result := ownerDocument.CreateDocumentFragment()
	if err == nil {
		for _, child := range nodes {
			iterate(ownerDocument, result, child)
		}
	}
	return result, err
}

func NewDOMParser() domParser { return domParser{} }

type ElementSteps interface {
	AppendChild(parent dom.Node, child dom.Node) dom.Node
}

func parseIntoDocument(doc dom.Document, r io.Reader) error {
	node, err := html.Parse(r)
	if err != nil {
		return err
	}
	iterateChildren(doc, doc, node)
	return nil
}

// convertNS converts the namespace URI from x/net/html to the _right_
// namespace.
// SVG elements have namespace "svg"
func convertNS(ns string) string {
	switch ns {
	case "svg":
		return "http://www.w3.org/2000/svg"
	default:
		return ns
	}
}

func createElementFromNode(
	d dom.Document,
	parent dom.Node,
	source *html.Node,
) dom.Element {
	if parent == nil {
		panic("Elements must have a parent")
	}

	var newElm dom.Element
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

func iterateChildren(d dom.Document, dest dom.Node, source *html.Node) {
	if dest, ok := dest.(interface{ Content() dom.DocumentFragment }); ok {
		iterateChildren(d, dest.Content(), source)
		return
	}
	for child := range source.ChildNodes() {
		iterate(d, dest, child)
	}
}

func iterate(d dom.Document, dest dom.Node, child *html.Node) {
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
