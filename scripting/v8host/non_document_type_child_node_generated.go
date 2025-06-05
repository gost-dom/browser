// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type nonDocumentTypeChildNodeV8Wrapper[T any] struct {
	handleReffedObject[dom.NonDocumentTypeChildNode, T]
}

func newNonDocumentTypeChildNodeV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *nonDocumentTypeChildNodeV8Wrapper[T] {
	return &nonDocumentTypeChildNodeV8Wrapper[T]{newHandleReffedObject[dom.NonDocumentTypeChildNode, T](scriptHost)}
}

func (wrapper nonDocumentTypeChildNodeV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w nonDocumentTypeChildNodeV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("previousElementSibling", w.previousElementSibling, nil)
	jsClass.CreatePrototypeAttribute("nextElementSibling", w.nextElementSibling, nil)
}

func (w nonDocumentTypeChildNodeV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nonDocumentTypeChildNodeV8Wrapper[T]) previousElementSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.previousElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousElementSibling()
	return encodeEntity(cbCtx, result)
}

func (w nonDocumentTypeChildNodeV8Wrapper[T]) nextElementSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.nextElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextElementSibling()
	return encodeEntity(cbCtx, result)
}
