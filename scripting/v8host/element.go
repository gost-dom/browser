package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (e *elementV8Wrapper[T]) CustomInitializer(class js.Class[T]) {
	class.CreatePrototypeMethod("insertAdjacentHTML", e.insertAdjacentHTML)
	class.CreatePrototypeAttribute("outerHTML", e.outerHTML, nil)
}

func (e *elementV8Wrapper[T]) insertAdjacentHTML(
	cbCtx js.CallbackContext[T],
) (val js.Value[T], err error) {
	element, e0 := js.As[dom.Element](cbCtx.Instance())
	position, e1 := consumeArgument(cbCtx, "position", nil, decodeString)
	html, e2 := consumeArgument(cbCtx, "html", nil, decodeString)
	err = errors.Join(e0, e1, e2)
	if err == nil {
		err = element.InsertAdjacentHTML(position, html)
	}
	return nil, err
}

func (e *elementV8Wrapper[T]) outerHTML(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	if element, err := js.As[dom.Element](cbCtx.Instance()); err == nil {
		return codec.EncodeString(cbCtx, element.OuterHTML())
	} else {
		return nil, err
	}
}

func (e elementV8Wrapper[T]) classList(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	cl := instance.ClassList()
	tokenList := cbCtx.Scope().Constructor("DOMTokenList")
	return tokenList.NewInstance(cl)
}

func (e *elementV8Wrapper[T]) toNamedNodeMap(
	cbCtx js.CallbackContext[T],
	n dom.NamedNodeMap,
) (js.Value[T], error) {

	return codec.EncodeEntity(cbCtx, n)
}
