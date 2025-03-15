package html

import (
	"strings"
)

type HTMLInputElement interface {
	HTMLElement
	Type() string
	SetType(value string)
	Name() string
	CheckValidity() bool
}

type htmlInputElement struct{ htmlElement }

func NewHTMLInputElement(ownerDocument HTMLDocument) HTMLInputElement {
	result := &htmlInputElement{newHTMLElement("input", ownerDocument)}
	result.SetSelf(result)
	return result
}

func (e *htmlInputElement) Name() string        { return e.GetAttributeNode("name").Value() }
func (e *htmlInputElement) CheckValidity() bool { return true }

func (e *htmlInputElement) Type() string {
	t, _ := e.GetAttribute("type")
	if t == "" {
		return "text"
	}
	return strings.ToLower(t)
}

func (e *htmlInputElement) SetType(val string) {
	e.SetAttribute("type", val)
}

func (e *htmlInputElement) Click() {
	ok := e.htmlElement.click()
	if ok && e.Type() == "submit" {
		e.trySubmitForm()
	}
}

func (e *htmlInputElement) trySubmitForm() {
	var form HTMLFormElement
	parent := e.Parent()
	for {
		if parent == nil {
			break
		}
		if f, ok := parent.(HTMLFormElement); ok {
			form = f
			break
		}
		parent = parent.Parent()
	}
	if form != nil {
		form.RequestSubmit(e)
	}
}
