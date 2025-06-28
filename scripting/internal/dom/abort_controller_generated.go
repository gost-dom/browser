// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
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
	instance, errInst := js.As[dominterfaces.AbortController](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	reason, errArg1 := js.ConsumeArgument(cbCtx, "reason", codec.ZeroValue, w.decodeAny)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Abort(reason)
	return nil, nil
}

func (w AbortController[T]) signal(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: AbortController.signal")
	instance, err := js.As[dominterfaces.AbortController](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Signal()
	return w.toAbortSignal(cbCtx, result)
}
