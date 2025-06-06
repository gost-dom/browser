// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type namedNodeMapV8Wrapper[T any] struct {
	handleReffedObject[dom.NamedNodeMap, T]
}

func newNamedNodeMapV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *namedNodeMapV8Wrapper[T] {
	return &namedNodeMapV8Wrapper[T]{newHandleReffedObject[dom.NamedNodeMap, T](scriptHost)}
}

func (wrapper namedNodeMapV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w namedNodeMapV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeMethod("getNamedItem", w.getNamedItem)
	jsClass.CreatePrototypeMethod("getNamedItemNS", w.getNamedItemNS)
	jsClass.CreatePrototypeMethod("setNamedItem", w.setNamedItem)
	jsClass.CreatePrototypeMethod("setNamedItemNS", w.setNamedItemNS)
	jsClass.CreatePrototypeMethod("removeNamedItem", w.removeNamedItem)
	jsClass.CreatePrototypeMethod("removeNamedItemNS", w.removeNamedItemNS)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w namedNodeMapV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w namedNodeMapV8Wrapper[T]) item(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.item")
	instance, errInst := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	index, errArg1 := consumeArgument(cbCtx, "index", nil, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Item(index)
	return encodeEntity(cbCtx, result)
}

func (w namedNodeMapV8Wrapper[T]) getNamedItem(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.getNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.getNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper[T]) getNamedItemNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.getNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.getNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper[T]) setNamedItem(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.setNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.setNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper[T]) setNamedItemNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.setNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.setNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper[T]) removeNamedItem(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.removeNamedItem")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.removeNamedItem: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper[T]) removeNamedItemNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.removeNamedItemNS")
	return cbCtx.ReturnWithError(errors.New("NamedNodeMap.removeNamedItemNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w namedNodeMapV8Wrapper[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NamedNodeMap.length")
	instance, err := js.As[dom.NamedNodeMap](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}
