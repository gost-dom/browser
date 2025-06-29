// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type NamedNodeMap[T any] struct{}

func NewNamedNodeMap[T any](scriptHost js.ScriptEngine[T]) *NamedNodeMap[T] {
	return &NamedNodeMap[T]{}
}

func (wrapper NamedNodeMap[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w NamedNodeMap[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeMethod("getNamedItem", w.getNamedItem)
	jsClass.CreatePrototypeMethod("getNamedItemNS", w.getNamedItemNS)
	jsClass.CreatePrototypeMethod("setNamedItem", w.setNamedItem)
	jsClass.CreatePrototypeMethod("setNamedItemNS", w.setNamedItemNS)
	jsClass.CreatePrototypeMethod("removeNamedItem", w.removeNamedItem)
	jsClass.CreatePrototypeMethod("removeNamedItemNS", w.removeNamedItemNS)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w NamedNodeMap[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w NamedNodeMap[T]) item(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.item")
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

func (w NamedNodeMap[T]) getNamedItem(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.getNamedItem")
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.getNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w NamedNodeMap[T]) getNamedItemNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.getNamedItemNS")
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.getNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w NamedNodeMap[T]) setNamedItem(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.setNamedItem")
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.setNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w NamedNodeMap[T]) setNamedItemNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.setNamedItemNS")
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.setNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w NamedNodeMap[T]) removeNamedItem(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.removeNamedItem")
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.removeNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w NamedNodeMap[T]) removeNamedItemNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.removeNamedItemNS")
	return codec.EncodeCallbackErrorf(cbCtx, "NamedNodeMap.removeNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w NamedNodeMap[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: NamedNodeMap.length")
	instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
