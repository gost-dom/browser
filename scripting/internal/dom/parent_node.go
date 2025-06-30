package dom

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w ParentNode[T]) toHTMLCollection(
	ctx js.CallbackContext[T],
	c dom.HTMLCollection,
) (js.Value[T], error) {
	return ctx.Constructor("HTMLCollection").NewInstance(c)
}
