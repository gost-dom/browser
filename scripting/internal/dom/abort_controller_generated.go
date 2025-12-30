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
	jsClass.CreateOperation("abort", AbortController_abort)
	jsClass.CreateAttribute("signal", AbortController_signal, nil)
}

func AbortControllerConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return CreateAbortController(cbCtx)
}

func AbortController_abort[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dominterfaces.AbortController](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	reason, errArg1 := js.ConsumeArgument(cbCtx, "reason", codec.ZeroValue, decodeAny)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Abort(reason)
	return nil, nil
}

func AbortController_signal[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbortController](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Signal()
	return encodeAbortSignal(cbCtx, result)
}
