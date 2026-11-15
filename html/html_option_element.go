package html

import "fmt"

type HTMLOptionElement interface {
	HTMLElement
	Value() string
	selectedness() bool
}

type htmlOptionElement struct {
	htmlElement
}

func NewHTMLOptionElement(ownerDocument HTMLDocument) HTMLOptionElement {
	result := &htmlOptionElement{
		htmlElement: newHTMLElement("option", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlOptionElement) selectedness() bool {
	v, ok := e.GetAttribute("selected")
	fmt.Println("SELECTED", v, ok)
	return ok
}

func (e *htmlOptionElement) Value() string {
	val, _ := e.GetAttribute("value")
	return val
}
