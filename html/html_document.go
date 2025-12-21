package html

import (
	"strings"

	"github.com/gost-dom/browser/dom"
)

type HTMLDocument interface {
	dom.Document
	Location() Location
	// unexported
	window() *window
	setActiveElement(e dom.Element)
	location() *location
	setLocation(*location)
}

type htmlDocument struct {
	dom.Document
	win         Window
	docLocation *location
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
	doc := NewEmptyHtmlDocument(window)
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

// NewEmptyHtmlDocument creates an HTML document without any content.
func NewEmptyHtmlDocument(window Window) HTMLDocument {
	var result HTMLDocument = &htmlDocument{dom.NewDocument(window), window, nil}
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
	case "label":
		return NewHTMLLabelElement(d)
	}
	return NewHTMLElement(name, d)
}

func (d *htmlDocument) Location() Location {
	d.Logger().Info("Document Location", "loc", d.docLocation)
	if d.docLocation == nil {
		return nil
	}
	return d.docLocation
}

func (d *htmlDocument) location() *location     { return d.docLocation }
func (d *htmlDocument) setLocation(l *location) { d.docLocation = l }

func (d *htmlDocument) window() *window {
	if d.win == nil {
		return nil
	}
	return d.win.window()
}

type SetActiveElementer interface {
	SetActiveElement(e dom.Element)
}

func (d *htmlDocument) setActiveElement(e dom.Element) {
	d.Document.(SetActiveElementer).SetActiveElement(e)
}
