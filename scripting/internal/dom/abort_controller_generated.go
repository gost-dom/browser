// This file is generated. Do not edit.

package dom

import (
	"errors"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type AbortController[T any] struct{}

func NewAbortController[T any](scriptHost js.ScriptEngine[T]) *AbortController[T] {
	return &AbortController[T]{}
}

func (wrapper AbortController[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w AbortController[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("abort", w.abort)
	jsClass.CreatePrototypeAttribute("signal", w.signal, nil)
}

func (w AbortController[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: AbortController.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w AbortController[T]) abort(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: AbortController.abort")
	return nil, errors.New("AbortController.abort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w AbortController[T]) signal(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: AbortController.signal")
	return nil, errors.New("AbortController.signal: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
