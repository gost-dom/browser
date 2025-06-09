// This file is generated. Do not edit.

package xhr

import js "github.com/gost-dom/browser/scripting/internal/js"

type XMLHttpRequestEventTargetV8Wrapper[T any] struct{}

func NewXMLHttpRequestEventTargetV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *XMLHttpRequestEventTargetV8Wrapper[T] {
	return &XMLHttpRequestEventTargetV8Wrapper[T]{}
}

func (wrapper XMLHttpRequestEventTargetV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w XMLHttpRequestEventTargetV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {}

func (w XMLHttpRequestEventTargetV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: XMLHttpRequestEventTarget.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}
