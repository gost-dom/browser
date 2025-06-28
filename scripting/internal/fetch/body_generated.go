// This file is generated. Do not edit.

package fetch

import (
	"errors"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Body[T any] struct{}

func NewBody[T any](scriptHost js.ScriptEngine[T]) *Body[T] {
	return &Body[T]{}
}

func (wrapper Body[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Body[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("arrayBuffer", w.arrayBuffer)
	jsClass.CreatePrototypeMethod("blob", w.blob)
	jsClass.CreatePrototypeMethod("bytes", w.bytes)
	jsClass.CreatePrototypeMethod("formData", w.formData)
	jsClass.CreatePrototypeMethod("json", w.json)
	jsClass.CreatePrototypeMethod("text", w.text)
	jsClass.CreatePrototypeAttribute("body", w.body, nil)
	jsClass.CreatePrototypeAttribute("bodyUsed", w.bodyUsed, nil)
}

func (w Body[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Body[T]) arrayBuffer(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.arrayBuffer")
	return nil, errors.New("Body.arrayBuffer: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) blob(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.blob")
	return nil, errors.New("Body.blob: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) bytes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.bytes")
	return nil, errors.New("Body.bytes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) formData(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.formData")
	return nil, errors.New("Body.formData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) text(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.text")
	return nil, errors.New("Body.text: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) body(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.body")
	return nil, errors.New("Body.body: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Body[T]) bodyUsed(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: Body.bodyUsed")
	return nil, errors.New("Body.bodyUsed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
