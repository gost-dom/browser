package html

import (
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/dom"
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
	return result
}

func (e *htmlTemplateElement) Content() dom.DocumentFragment { return e.content }

func (e *htmlTemplateElement) RenderChildren(builder *strings.Builder) {
	if renderer, ok := e.content.(ChildrenRenderer); ok {
		renderer.RenderChildren(builder)
	}
}
