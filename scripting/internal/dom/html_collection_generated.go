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
	jsClass.CreateOperation("item", w.item)
	jsClass.CreateOperation("namedItem", w.namedItem)
	jsClass.CreateAttribute("length", w.length, nil)
}

func HTMLCollectionConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLCollection[T]) item(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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
	instance, err := js.As[dom.HTMLCollection](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
