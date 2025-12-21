package html

import (
	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/entity"
)

type HTMLTemplateElement interface {
	HTMLElement
	Content() dom.DocumentFragment
}

type htmlTemplateElement struct {
	htmlElement
	content dom.DocumentFragment
}

func NewHTMLTemplateElement(ownerDocument HTMLDocument) HTMLTemplateElement {
	result := &htmlTemplateElement{
		newHTMLElement("template", ownerDocument),
		dom.NewDocumentFragment(ownerDocument),
	}
	result.SetSelf(result)
	entity.SetComponentType(result, result.content.(ChildrenRenderer))
	entity.SetComponentType(result, result.content.(dom.ContentContainer))
	return result
}

func (e *htmlTemplateElement) Content() dom.DocumentFragment { return e.content }
