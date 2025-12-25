// This file is generated. Do not edit.

package dom

import (
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Text[T any] struct{}

func NewText[T any](scriptHost js.ScriptEngine[T]) *Text[T] {
	return &Text[T]{}
}

func (wrapper Text[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Text[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("splitText", w.splitText)
	jsClass.CreateAttribute("wholeText", w.wholeText, nil)
}

func (w Text[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Text[T]) splitText(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Text.splitText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Text[T]) wholeText(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Text.wholeText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
