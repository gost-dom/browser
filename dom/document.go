package dom

import (
	"io"

	"github.com/gost-dom/browser/dom/event"
	"golang.org/x/net/html"
)

type DocumentEvent = string

type staticNodeList NodeList

const (
	DocumentEventDOMContentLoaded DocumentEvent = "DOMContentLoaded"
	DocumentEventLoad             DocumentEvent = "load"
)

// Deprecated: This interface is part of an implementation details, and it was
// an oversight that it wasn't placed in an internal package. This will be
// removed from the public API in a future version
type DocumentParentWindow interface {
	event.EventTarget
	ParseFragment(ownerDocument Document, reader io.Reader) (DocumentFragment, error)
}

type Document interface {
	RootNode
	Body() Element
	Head() Element
	CreateDocumentFragment() DocumentFragment
	CreateAttribute(string) Attr
	CreateText(data string) Text
	CreateComment(data string) Comment
	CreateDocumentType(name string) DocumentType
	CreateElementNS(string, string) Element
	CreateElement(string) Element
	DocumentElement() Element
	parseFragment(reader io.Reader) (DocumentFragment, error)
}

type elementConstructor func(doc *document) Element

type document struct {
	rootNode
	ownerWindow DocumentParentWindow
}

func NewDocument(window DocumentParentWindow) Document {
	result := &document{newRootNode(nil), window}
	// Hmmm, can document be replaced; and now the old doc's event goes to a
	// window they shouldn't?
	// What about disconnected documents, e.g. `new Document()` in the browser?
	result.SetParentTarget(window)
	result.SetSelf(result)
	return result
}

func (d *document) CloneNode(deep bool) Node {
	result := NewDocument(d.ownerWindow)
	if deep {
		result.Append(d.cloneChildren()...)
	}
	return result
}

func (d *document) parseFragment(reader io.Reader) (DocumentFragment, error) {
	return d.ownerWindow.ParseFragment(d, reader)
}

func (d *document) Body() Element {
	root := d.DocumentElement()
	if root != nil {
		for _, child := range root.ChildNodes().All() {
			if e, ok := child.(Element); ok {
				if e.TagName() == "BODY" {
					return e
				}
			}
		}
	}
	return nil
}

func (d *document) Head() Element {
	root := d.DocumentElement()
	if root != nil {
		for _, child := range root.ChildNodes().All() {
			if e, ok := child.(Element); ok {
				if e.TagName() == "HEAD" {
					return e
				}
			}
		}
	}
	return nil
}

func (d *document) CreateAttribute(name string) Attr {
	return newAttr(name, "", d)
}

func (d *document) CreateElement(name string) Element {
	return NewElement(name, d)
}

func (d *document) CreateText(data string) Text                 { return NewText(data, d) }
func (d *document) CreateComment(data string) Comment           { return NewComment(data, d) }
func (d *document) CreateDocumentType(name string) DocumentType { return NewDocumentType(name, d) }

func (d *document) CreateElementNS(_ string, name string) Element {
	return NewElement(name, d)
}

func (d *document) CreateDocumentFragment() DocumentFragment {
	return NewDocumentFragment(d)
}

func (d *document) DocumentElement() Element {
	for _, c := range d.ChildNodes().All() {
		if e, ok := c.(Element); ok {
			return e
		}
	}
	return nil
}

func (d *document) NodeName() string  { return "#document" }
func (d *document) IsConnected() bool { return true }

func (d *document) GetElementById(id string) Element {
	return rootNodeHelper{d}.GetElementById(id)
}

func (d *document) createHtmlNode() *html.Node {
	return &html.Node{
		Type: html.DocumentNode,
	}
}

func (d *document) NodeType() NodeType { return NodeTypeDocument }
