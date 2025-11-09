package dom

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (e *Element[T]) CustomInitializer(class js.Class[T]) {
	class.CreatePrototypeMethod("insertAdjacentHTML", e.insertAdjacentHTML)
	class.CreatePrototypeAttribute("outerHTML", e.outerHTML, nil)
	class.CreatePrototypeAttribute("innerHTML", e.innerHTML, e.setInnerHTML)
}

func (e *Element[T]) insertAdjacentHTML(
	cbCtx js.CallbackContext[T],
) (val js.Value[T], err error) {
	element, e0 := js.As[dom.Element](cbCtx.Instance())
	position, e1 := js.ConsumeArgument(cbCtx, "position", nil, codec.DecodeString)
	html, e2 := js.ConsumeArgument(cbCtx, "html", nil, codec.DecodeString)
	err = errors.Join(e0, e1, e2)
	if err == nil {
		err = element.InsertAdjacentHTML(position, html)
	}
	return nil, err
}

func (e *Element[T]) outerHTML(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	if element, err := js.As[dom.Element](cbCtx.Instance()); err == nil {
		return codec.EncodeString(cbCtx, element.OuterHTML())
	} else {
		return nil, err
	}
}

func (e *Element[T]) innerHTML(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	if element, err := js.As[dom.Element](cbCtx.Instance()); err == nil {
		return codec.EncodeString(cbCtx, element.InnerHTML())
	} else {
		return nil, err
	}
}

func (e *Element[T]) setInnerHTML(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	html, err1 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeString)
	element, errInstance := js.As[dom.Element](cbCtx.Instance())

	err := errors.Join(err1, errInstance)
	if err == nil {
		err = element.SetInnerHTML(html)
	}
	return nil, err
}

func (e Element[T]) classList(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	cl := instance.ClassList()
	tokenList := cbCtx.Constructor("DOMTokenList")
	return tokenList.NewInstance(cl)
}

func (e *Element[T]) toNamedNodeMap(
	cbCtx js.CallbackContext[T],
	n dom.NamedNodeMap,
) (js.Value[T], error) {
	return codec.EncodeEntity(cbCtx, n)
}

func (e *Element[T]) decodeElement(
	ctx js.Scope[T],
	val js.Value[T],
) (dom.Element, error) {
	return codec.DecodeAs[dom.Element](ctx, val)
}
