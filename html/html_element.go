package html

import (
	"strconv"
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/uievents"
)

type HTMLElement interface {
	dom.Element
	Renderer
	ChildrenRenderer
	Click() bool
	Blur() bool
	Focus() bool
	Tabindex() int
	SetTabindex(int)
	Dataset() DOMStringMap
	getHTMLDocument() HTMLDocument
	window() Window
}

type htmlElement struct {
	dom.Element
	Renderer
	ChildrenRenderer
	htmlDocument HTMLDocument
	dataset      DOMStringMap
	self         HTMLElement
}

func NewHTMLElement(tagName string, ownerDocument HTMLDocument) HTMLElement {
	return newHTMLElement(tagName, ownerDocument)
}

func newHTMLElement(tagName string, ownerDocument HTMLDocument) *htmlElement {
	element := dom.NewElement(tagName, ownerDocument)
	renderer, _ := element.(Renderer)
	childrenRenderer, _ := element.(ChildrenRenderer)
	result := &htmlElement{
		Element:          element,
		Renderer:         renderer,
		ChildrenRenderer: childrenRenderer,
		htmlDocument:     ownerDocument,
		dataset:          DOMStringMap{element},
	}
	result.SetSelf(result)
	return result
}

func (e *htmlElement) SetSelf(self dom.Node) {
	e.self = self.(HTMLElement)
	e.Element.SetSelf(self)
}

func (e *htmlElement) getHTMLDocument() HTMLDocument { return e.htmlDocument }

func (e *htmlElement) window() Window { return e.getHTMLDocument().getWindow() }

func (e *htmlElement) TagName() string {
	return strings.ToUpper(e.Element.TagName())
}

func (e *htmlElement) Click() bool { return uievents.Click(e.Element) }

func (e *htmlElement) Blur() bool {
	uievents.Blur(e.self)
	uievents.Focusout(e.self)
	e.htmlDocument.setActiveElement(nil)
	return true
}

func (e *htmlElement) Focus() bool {
	if oldTarget, ok := e.htmlDocument.ActiveElement().(HTMLElement); ok {
		oldTarget.Blur()
	}
	uievents.Focus(e.Element)
	uievents.Focusin(e.Element)
	e.htmlDocument.setActiveElement(e.self)
	return true
}

func (e *htmlElement) Dataset() DOMStringMap { return e.dataset }

func (e *htmlElement) Tabindex() int {
	val, ok := e.GetAttribute("tabindex")
	res, err := strconv.Atoi(val)
	if ok && err == nil {
		return res
	} else {
		return -1
	}
}

func (e *htmlElement) SetTabindex(i int) { e.SetAttribute("tabindex", strconv.Itoa(i)) }

func (e *htmlElement) Nonce() string         { v, _ := e.GetAttribute("nonce"); return v }
func (e *htmlElement) SetNonce(nonce string) { e.SetAttribute("nonce", nonce) }
