package html

import "iter"

type HTMLSelectElement interface {
	HTMLElement
	Name() string
	SetName(string)
	Value() string
	SetValue(string)
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

func (e *htmlSelectElement) options() iter.Seq[*htmlOptionElement] {
	return func(yield func(*htmlOptionElement) bool) {
		opts, _ := e.QuerySelectorAll("option")
		for opt := range opts.All() {
			if optElement, ok := opt.(*htmlOptionElement); ok {
				if !yield(optElement) {
					return
				}
			}
		}
	}
}

func (e *htmlSelectElement) Value() string {
	// https://html.spec.whatwg.org/multipage/form-elements.html#dom-select-value

	for opt := range e.options() {
		if opt.Selectedness() {
			return opt.Value()
		}
	}
	return ""
}

func (e *htmlSelectElement) SetValue(v string) {
	var firstMatchingOption *htmlOptionElement
	for opt := range e.options() {
		opt.SetSelectedness(false)
		if opt.Value() == v && firstMatchingOption == nil {
			firstMatchingOption = opt
		}
	}
	if firstMatchingOption != nil {
		firstMatchingOption.SetSelectedness(true)
	}
}
