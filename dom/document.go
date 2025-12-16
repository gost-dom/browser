package dom

import (
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
	CreateProcessingInstruction(string, string) ProcessingInstruction
	DocumentElement() Element
	GetElementsByTagName(string) NodeList
	ImportNode(Node, bool) Node
}

type document struct {
	rootNode
	logger        *slog.Logger
	activeElement Element
}

func NewDocument(parentEventTarget event.EventTarget) Document {
	result := &document{rootNode: newRootNode(nil)}
	// Hmmm, can document be replaced; and now the old doc's event goes to a
	// window they shouldn't?
	// What about disconnected documents, e.g. `new Document()` in the browser?
	result.SetParentTarget(parentEventTarget)
	result.SetSelf(result)
	if logger, ok := parentEventTarget.(log.LogSource); ok {
		result.logger = logger.Logger()
	}
	return result
}

func (d document) Logger() *slog.Logger {
	if d.logger != nil {
		return d.logger
	}
	return log.Default()
}

func (d document) ActiveElement() Element {
	if d.activeElement == nil {
		return d.Body()
	}
	return d.activeElement
}

func (d *document) cloneNode(doc Document, deep bool) Node {
	result := NewDocument(nil)
	if deep {
		result.Append(d.cloneChildren()...)
	}
	return result
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
		return &nodeList{}
	}
	return res
}

func (d *document) CreateProcessingInstruction(target string, data string) ProcessingInstruction {
	return NewProcessingInstruction(target, data, d.document)
}
