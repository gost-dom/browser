package html

type HTMLTextAreaElement interface {
	HTMLElement
	Value() string
	SetValue(string)
}

type htmlTextAreaElement struct {
	htmlElement
	value string
}

func NewHTMLTextAreaElement(ownerDocument HTMLDocument) HTMLTextAreaElement {
	result := &htmlTextAreaElement{
		htmlElement: newHTMLElement("textarea", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlTextAreaElement) Value() string {
	return e.value
}
func (e *htmlTextAreaElement) SetValue(v string) {
	e.value = v
}
