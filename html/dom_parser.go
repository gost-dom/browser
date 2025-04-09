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
			element := createElementFromNode(ownerDocument, result, child)
			result.AppendChild(element)
		}
	}
	return result, err
}

func NewDOMParser() domParser { return domParser{} }

type ElementSteps interface {
	AppendChild(parent dom.Node, child dom.Node) dom.Node
}

type BaseRules struct{}

func (r BaseRules) AppendChild(parent dom.Node, child dom.Node) dom.Node {
	res, err := parent.AppendChild(child)
	if err != nil {
		panic(err)
	}
	return res
}

type TemplateElementRules struct{ BaseRules }

func (TemplateElementRules) AppendChild(parent dom.Node, child dom.Node) dom.Node {
	template, ok := child.(HTMLTemplateElement)
	if !ok {
		panic("Parser error, applying tepmlate rules to non-template element")
	}
	parent.AppendChild(child)
	return template.Content()
}

var ElementMap = map[atom.Atom]ElementSteps{
	// atom.Script:   ScriptElementRules{},
	atom.Template: TemplateElementRules{},
}

func parseIntoDocument(doc dom.Document, r io.Reader) error {
	node, err := html.Parse(r)
	if err != nil {
		return err
	}
	iterate(doc, doc, node)
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

	rules := ElementMap[source.DataAtom]
	if rules == nil {
		rules = BaseRules{}
	}
	var newNode dom.Node
	var newElm dom.Element
	if source.Namespace == "" {
		newElm = d.CreateElement(source.Data)
	} else {
		newElm = d.CreateElementNS(convertNS(source.Namespace), source.Data)
	}
	for _, a := range source.Attr {
		newElm.SetAttribute(a.Key, a.Val)
	}
	newNode = newElm
	newNode = rules.AppendChild(parent, newElm)
	iterate(d, newNode, source)
	return newElm
}

func iterate(d dom.Document, dest dom.Node, source *html.Node) {
	for child := range source.ChildNodes() {
		switch child.Type {
		case html.ElementNode:
			createElementFromNode(d, dest, child)
		case html.TextNode:
			dest.AppendChild(d.CreateText(child.Data))
		case html.DoctypeNode:
			dest.AppendChild(d.CreateDocumentType(child.Data))
		case html.CommentNode:
			dest.AppendChild(d.CreateComment(child.Data))
		default:
			panic(fmt.Sprintf("Node not yet supported: %v", child.Type))
		}
	}
}
