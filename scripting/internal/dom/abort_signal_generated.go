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
	jsClass.CreatePrototypeMethod("throwIfAborted", w.throwIfAborted)
	jsClass.CreatePrototypeAttribute("aborted", w.aborted, nil)
	jsClass.CreatePrototypeAttribute("reason", w.reason, nil)
}

func (w AbortSignal[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: AbortSignal.Constructor", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w AbortSignal[T]) throwIfAborted(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: AbortSignal.throwIfAborted", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dominterfaces.AbortSignal](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.ThrowIfAborted()
	return nil, errCall
}

func (w AbortSignal[T]) aborted(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: AbortSignal.aborted", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dominterfaces.AbortSignal](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Aborted()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w AbortSignal[T]) reason(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: AbortSignal.reason", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "AbortSignal.reason: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
