// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type NonDocumentTypeChildNodeV8Wrapper[T any] struct{}

func NewNonDocumentTypeChildNodeV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *NonDocumentTypeChildNodeV8Wrapper[T] {
	return &NonDocumentTypeChildNodeV8Wrapper[T]{}
}

func (wrapper NonDocumentTypeChildNodeV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w NonDocumentTypeChildNodeV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("previousElementSibling", w.previousElementSibling, nil)
	jsClass.CreatePrototypeAttribute("nextElementSibling", w.nextElementSibling, nil)
}

func (w NonDocumentTypeChildNodeV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w NonDocumentTypeChildNodeV8Wrapper[T]) previousElementSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.previousElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousElementSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func (w NonDocumentTypeChildNodeV8Wrapper[T]) nextElementSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.nextElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextElementSibling()
	return codec.EncodeEntity(cbCtx, result)
}
