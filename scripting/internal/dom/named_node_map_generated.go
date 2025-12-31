// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeNamedNodeMap[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("item", NamedNodeMap_item)
	jsClass.CreateOperation("getNamedItem", NamedNodeMap_getNamedItem)
	jsClass.CreateOperation("getNamedItemNS", NamedNodeMap_getNamedItemNS)
	jsClass.CreateOperation("setNamedItem", NamedNodeMap_setNamedItem)
	jsClass.CreateOperation("setNamedItemNS", NamedNodeMap_setNamedItemNS)
	jsClass.CreateOperation("removeNamedItem", NamedNodeMap_removeNamedItem)
	jsClass.CreateOperation("removeNamedItemNS", NamedNodeMap_removeNamedItemNS)
	jsClass.CreateAttribute("length", NamedNodeMap_length, nil)
	NamedNodeMapCustomInitializer(jsClass)
}

func NamedNodeMap_item[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.NamedNodeMap](cbCtx.Instance())
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

func NamedNodeMap_getNamedItem[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.NamedNodeMap_getNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func NamedNodeMap_getNamedItemNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.NamedNodeMap_getNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func NamedNodeMap_setNamedItem[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.NamedNodeMap_setNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func NamedNodeMap_setNamedItemNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.NamedNodeMap_setNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func NamedNodeMap_removeNamedItem[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.NamedNodeMap_removeNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func NamedNodeMap_removeNamedItemNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.NamedNodeMap_removeNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func NamedNodeMap_length[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
