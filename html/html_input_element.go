package html

import (
	"strings"
)

type HTMLInputElement interface {
	HTMLElement
	Type() string
	SetType(value string)
	Name() string
	SetName(string)
	Value() string
	SetValue(string)
	CheckValidity() bool
	Checked() bool
	SetChecked(bool)
}

type htmlInputElement struct {
	htmlElement
	checked bool
	value   string
}

func NewHTMLInputElement(ownerDocument HTMLDocument) HTMLInputElement {
	result := &htmlInputElement{
		htmlElement: newHTMLElement("input", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlInputElement) Name() string         { return e.GetAttributeNode("name").Value() }
func (e *htmlInputElement) SetName(value string) { e.SetAttribute("name", value) }
func (e *htmlInputElement) CheckValidity() bool  { return true }
func (e *htmlInputElement) Checked() bool        { return e.checked }
func (e *htmlInputElement) SetChecked(b bool)    { e.checked = b }
func (e *htmlInputElement) Value() string {
	value := e.value
	if value == "" {
		value, _ = e.GetAttribute("value")
	}
	return value
}

func (e *htmlInputElement) SetValue(value string) { e.value = value }

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
	if ok := e.htmlElement.click(); !ok {
		return
	}
	switch e.Type() {
	case "submit":
		e.trySubmitForm()
	case "checkbox":
		e.SetChecked(!e.checked)
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
