package html

import (
	"strings"

	"github.com/gost-dom/browser/dom"
)

type HTMLDocument interface {
	dom.Document
	// unexported
	getWindow() Window
	setActiveElement(e dom.Element)
}

type htmlDocument struct {
	dom.Document
	window Window
}

func mustAppendChild(p, c dom.Node) dom.Node {
	_, err := p.AppendChild(c)
	if err != nil {
		panic(err)
	}
	return c
}

// NewHTMLDocument creates an HTML document for an about:blank page.
//
// The resulting document has an outer HTML similar to this, but there are no
// guarantees about the actual content.
//
//	<html><head></head><body><h1>Gost-DOM</h1></body></html>
func NewHTMLDocument(window Window) HTMLDocument {
	doc := newHTMLDocument(window)
	body := doc.CreateElement("body")
	docEl := doc.CreateElement("html")
	h1 := mustAppendChild(body, doc.CreateElement("h1"))
	h1.SetTextContent("Gost-DOM")
	docEl.Append(
		doc.CreateElement("head"),
		body,
	)
	doc.AppendChild(docEl)
	return doc
}

// newHTMLDocument is used internally to create an empty HTML when parsing an
// HTML input.
func newHTMLDocument(window Window) HTMLDocument {
	var parent dom.DocumentParentWindow
	if window != nil {
		parent = window
	}
	var result HTMLDocument = &htmlDocument{dom.NewDocument(parent), window}
	result.SetSelf(result)
	return result
}

func (d *htmlDocument) CreateElementNS(namespace string, name string) dom.Element {
	if namespace == "http://www.w3.org/1999/xhtml" {
		return d.CreateElement(name)
	}
	return d.Document.CreateElementNS(namespace, name)
}

func (d *htmlDocument) CreateElement(name string) dom.Element {
	switch strings.ToLower(name) {
	case "template":
		return NewHTMLTemplateElement(d)
	case "form":
		return NewHtmlFormElement(d)
	case "input":
		return NewHTMLInputElement(d)
	case "button":
		return NewHTMLButtonElement(d)
	case "script":
		return NewHTMLScriptElement(d)
	case "a":
		return NewHTMLAnchorElement(d)
	}
	return NewHTMLElement(name, d)
}

func (d *htmlDocument) getWindow() Window { return d.window }

type SetActiveElementer interface {
	SetActiveElement(e dom.Element)
}

func (d *htmlDocument) setActiveElement(e dom.Element) {
	d.Document.(SetActiveElementer).SetActiveElement(e)
}
