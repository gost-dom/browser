// This file is generated. Do not edit.

package xhr

import js "github.com/gost-dom/browser/scripting/internal/js"

type XMLHttpRequestEventTarget[T any] struct{}

func NewXMLHttpRequestEventTarget[T any](scriptHost js.ScriptEngine[T]) *XMLHttpRequestEventTarget[T] {
	return &XMLHttpRequestEventTarget[T]{}
}

func (wrapper XMLHttpRequestEventTarget[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w XMLHttpRequestEventTarget[T]) installPrototype(jsClass js.Class[T]) {}

func (w XMLHttpRequestEventTarget[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: XMLHttpRequestEventTarget.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: XMLHttpRequestEventTarget.Constructor", js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}
