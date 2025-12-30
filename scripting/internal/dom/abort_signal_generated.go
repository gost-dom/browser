// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type AbortSignal[T any] struct{}

func NewAbortSignal[T any](scriptHost js.ScriptEngine[T]) *AbortSignal[T] {
	return &AbortSignal[T]{}
}

func (wrapper AbortSignal[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w AbortSignal[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("throwIfAborted", AbortSignal_throwIfAborted)
	jsClass.CreateAttribute("aborted", AbortSignal_aborted, nil)
	jsClass.CreateAttribute("reason", AbortSignal_reason, nil)
}

func AbortSignalConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func AbortSignal_throwIfAborted[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbortSignal](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.ThrowIfAborted()
	return nil, errCall
}

func AbortSignal_aborted[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dominterfaces.AbortSignal](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Aborted()
	return codec.EncodeBoolean(cbCtx, result)
}

func AbortSignal_reason[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "AbortSignal.AbortSignal_reason: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
