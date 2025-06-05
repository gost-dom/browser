// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("NodeList", "", newNodeListV8Wrapper)
}

type nodeListV8Wrapper[T any] struct {
	handleReffedObject[dom.NodeList, T]
}

func newNodeListV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *nodeListV8Wrapper[T] {
	return &nodeListV8Wrapper[T]{newHandleReffedObject[dom.NodeList, T](scriptHost)}
}

func (wrapper nodeListV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w nodeListV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w nodeListV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NodeList.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeListV8Wrapper[T]) item(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NodeList.item")
	instance, errInst := js.As[dom.NodeList](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	index, errArg1 := consumeArgument(cbCtx, "index", nil, w.decodeUnsignedLong)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Item(index)
	return encodeEntity(cbCtx, result)
}

func (w nodeListV8Wrapper[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NodeList.length")
	instance, err := js.As[dom.NodeList](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx, result)
}
