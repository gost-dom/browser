package html

import (
	"github.com/gost-dom/browser/dom"
)

type HTMLOptionElement interface {
	HTMLElement
	Value() string
	Selected() bool
}

type htmlOptionElement struct {
	htmlElement

	selectedness bool
}

func NewHTMLOptionElement(ownerDocument HTMLDocument) HTMLOptionElement {
	result := &htmlOptionElement{
		htmlElement: newHTMLElement("option", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlOptionElement) Selected() bool {
	return e.Selectedness()
}

func (e *htmlOptionElement) Selectedness() bool {
	return e.selectedness
}

func (e *htmlOptionElement) SetSelectedness(s bool) {
	e.selectedness = s
}

func (e *htmlOptionElement) Value() string {
	val, _ := e.GetAttribute("value")
	return val
}

func (e *htmlOptionElement) SetAttributeNode(a dom.Attr) (dom.Attr, error) {
	if a.Name() == "selected" {
		e.selectedness = true
	}
	return e.htmlElement.SetAttributeNode(a)
}
