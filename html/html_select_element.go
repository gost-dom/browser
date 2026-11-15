package html

type HTMLSelectElement interface {
	HTMLElement
	Name() string
	SetName(string)
	Value() string
}

type htmlSelectElement struct {
	htmlElement
}

func NewHTMLSelectElement(ownerDocument HTMLDocument) HTMLSelectElement {
	result := &htmlSelectElement{
		htmlElement: newHTMLElement("select", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlSelectElement) Name() string {
	r, _ := e.GetAttribute("name")
	return r
}

func (e *htmlSelectElement) SetName(value string) { e.SetAttribute("name", value) }

func (e *htmlSelectElement) Value() string {
	// https://html.spec.whatwg.org/multipage/form-elements.html#dom-select-value

	opts, _ := e.QuerySelectorAll("option")
	for opt := range opts.All() {
		if optElement, ok := opt.(HTMLOptionElement); ok {
			if optElement.selectedness() {
				return optElement.Value()
			}
		}
	}
	return ""
}
