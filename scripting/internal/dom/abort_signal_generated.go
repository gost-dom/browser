// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeAbortSignal[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("throwIfAborted", AbortSignal_throwIfAborted)
	jsClass.CreateAttribute("aborted", AbortSignal_aborted, nil)
	jsClass.CreateAttribute("reason", AbortSignal_reason, nil)
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
