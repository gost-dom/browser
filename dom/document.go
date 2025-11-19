package dom

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/log"
	"golang.org/x/net/html"
)

type DocumentEvent = string

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
	ActiveElement() Element
	Body() Element
	Head() Element
	CreateDocumentFragment() DocumentFragment
	CreateAttribute(string) Attr
	CreateTextNode(data string) Text
	CreateComment(data string) Comment
	CreateDocumentType(name string) DocumentType
	CreateElementNS(string, string) Element
	CreateElement(string) Element
	DocumentElement() Element
	GetElementsByTagName(string) NodeList
	ImportNode(Node, bool) Node
	parseFragment(reader io.Reader) (DocumentFragment, error)

	window() DocumentParentWindow
}

type document struct {
	rootNode
	logger        *slog.Logger
	activeElement Element
	ownerWindow   DocumentParentWindow
}

func NewDocument(window DocumentParentWindow) Document {
	result := &document{
		rootNode:    newRootNode(nil),
		ownerWindow: window,
	}
	// Hmmm, can document be replaced; and now the old doc's event goes to a
	// window they shouldn't?
	// What about disconnected documents, e.g. `new Document()` in the browser?
	result.SetParentTarget(window)
	result.SetSelf(result)
	if logger, isLogSource := window.(log.LogSource); isLogSource {
		result.logger = logger.Logger()
	}
	return result
}

func (d document) Logger() *slog.Logger         { return d.logger }
func (d document) window() DocumentParentWindow { return d.ownerWindow }

func (d document) ActiveElement() Element {
	if d.activeElement == nil {
		return d.Body()
	}
	return d.activeElement
}

func (d *document) cloneNode(doc Document, deep bool) Node {
	result := NewDocument(doc.window())
	if deep {
		result.Append(d.cloneChildren()...)
	}
	return result
}

func (d *document) parseFragment(reader io.Reader) (DocumentFragment, error) {
	return d.ownerWindow.ParseFragment(d.nodeDocument(), reader)
}

func (d *document) ImportNode(n Node, deep bool) Node {
	return n.cloneNode(d.getSelf().(Document), deep)
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

func (d *document) CreateAttribute(name string) Attr  { return newAttr(name, "", d.document) }
func (d *document) CreateElement(name string) Element { return NewElement(name, d.document) }
func (d *document) CreateText(data string) Text       { return d.CreateTextNode(data) }
func (d *document) CreateTextNode(data string) Text   { return NewText(data, d.document) }
func (d *document) CreateComment(data string) Comment { return NewComment(data, d.document) }
func (d *document) CreateElementNS(_ string, name string) Element {
	return NewElement(name, d.document)
}
func (d *document) CreateDocumentType(name string) DocumentType {
	return NewDocumentType(name, d.document)
}

func (d *document) CreateDocumentFragment() DocumentFragment {
	return NewDocumentFragment(d.document)
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

func (d *document) SetActiveElement(e Element) {
	d.activeElement = e
}

func (n *document) GetElementsByTagName(qualifiedName string) NodeList {
	res, err := n.QuerySelectorAll(qualifiedName)
	if err != nil {
		panic(fmt.Sprintf("document.GetElementsByTagName: %v", err))
	}
	return res
}
