package html

import (
	"regexp"
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/dom"
)

// DOMStringMap provides access to data-* attributes of an HTML or SVG element.
// In JavaScript, it is a dictionary-like object wrapping content with a "data-"
// prefix, converting kebab-case to camel-case, and stripping the prefix.
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset
type DOMStringMap struct{ Element dom.Element }

var camelCaseDetector = regexp.MustCompile("[a-z][A-Z]")

// toKebab converts a camelCase string to a kebabString.
//
// This is intended for [HTMLElement.Dataset] that provides a camelCase API over
// the kebab-case data- content attribute names.
func toKebab(str string) string {
	return camelCaseDetector.ReplaceAllStringFunc(str, func(match string) string {
		lower := []rune(strings.ToLower(match))
		return string([]rune{lower[0], '-', lower[1]})
	})
}

func (m DOMStringMap) Get(key string) (val string, exists bool) {
	key = toKebab(key)
	return m.Element.GetAttribute("data-" + key)
}

type HTMLElement interface {
	dom.Element
	Renderer
	ChildrenRenderer
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
