// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLCollection[T any] struct{}

func NewHTMLCollection[T any](scriptHost js.ScriptEngine[T]) *HTMLCollection[T] {
	return &HTMLCollection[T]{}
}

func (wrapper HTMLCollection[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w HTMLCollection[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeMethod("namedItem", w.namedItem)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w HTMLCollection[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLCollection.Constructor", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLCollection[T]) item(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLCollection.item", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.HTMLCollection](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	index, errArg1 := js.ConsumeArgument(cbCtx, "index", nil, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Item(index)
	return codec.EncodeEntity(cbCtx, result)
}

func (w HTMLCollection[T]) namedItem(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLCollection.namedItem", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.HTMLCollection](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.NamedItem(name)
	return codec.EncodeEntity(cbCtx, result)
}

func (w HTMLCollection[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLCollection.length", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.HTMLCollection](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
