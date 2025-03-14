package html

import (
	"strconv"
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/dom"
)

type HTMLElement interface {
	dom.Element
	Renderer
	ChildrenRenderer
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

func (e *htmlElement) getHTMLDocument() HTMLDocument { return e.htmlDocument }

func (e *htmlElement) window() Window { return e.getHTMLDocument().getWindow() }

func (e *htmlElement) TagName() string {
	return strings.ToUpper(e.Element.TagName())
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
