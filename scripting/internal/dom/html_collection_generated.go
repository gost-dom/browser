// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeHTMLCollection[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("item", HTMLCollection_item)
	jsClass.CreateOperation("namedItem", HTMLCollection_namedItem)
	jsClass.CreateAttribute("length", HTMLCollection_length, nil)
	HTMLCollectionCustomInitializer(jsClass)
}

func HTMLCollectionConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func HTMLCollection_item[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func HTMLCollection_namedItem[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func HTMLCollection_length[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.HTMLCollection](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
