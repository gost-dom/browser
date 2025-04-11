package html

type HTMLLabelElement interface {
	HTMLElement
	HTMLFor() string
	SetHTMLFor(string)
}

type htmlLabelElement struct {
	htmlElement
}

func NewHTMLLabelElement(ownerDoc HTMLDocument) HTMLLabelElement {
	result := &htmlLabelElement{
		htmlElement: newHTMLElement("label", ownerDoc),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlLabelElement) HTMLFor() string {
	f, _ := e.GetAttribute("for")
	return f
}

func (e *htmlLabelElement) SetHTMLFor(f string) {
	e.SetAttribute("for", f)
}

func (e *htmlLabelElement) Click() {
	if ok := e.htmlElement.click(); !ok {
		return
	}
	id := e.HTMLFor()
	if input, _ := e.OwnerDocument().GetElementById(id).(HTMLElement); input != nil {
		input.Click()
	}
}
