// This file is generated. Do not edit.

package html

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeWindowOrWorkerGlobalScope[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("setTimeout", WindowOrWorkerGlobalScope_setTimeout)
	jsClass.CreateOperation("clearTimeout", WindowOrWorkerGlobalScope_clearTimeout)
	jsClass.CreateOperation("setInterval", WindowOrWorkerGlobalScope_setInterval)
	jsClass.CreateOperation("clearInterval", WindowOrWorkerGlobalScope_clearInterval)
	jsClass.CreateOperation("queueMicrotask", WindowOrWorkerGlobalScope_queueMicrotask)
}

func WindowOrWorkerGlobalScope_setTimeout[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.WindowOrWorkerGlobalScope](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	handler, errArg1 := js.ConsumeArgument(cbCtx, "handler", nil, decodeTimerHandler)
	timeout, errArg2 := js.ConsumeArgument(cbCtx, "timeout", codec.ZeroValue, codec.DecodeDuration)
	cbCtx.ConsumeArg()
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.SetTimeout(handler, timeout)
	return encodeTaskHandle(cbCtx, result)
}

func WindowOrWorkerGlobalScope_clearTimeout[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.WindowOrWorkerGlobalScope](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	id, errArg1 := js.ConsumeArgument(cbCtx, "id", codec.ZeroValue, decodeTaskHandle)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.ClearTimeout(id)
	return nil, nil
}

func WindowOrWorkerGlobalScope_setInterval[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.WindowOrWorkerGlobalScope](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	handler, errArg1 := js.ConsumeArgument(cbCtx, "handler", nil, decodeTimerHandler)
	timeout, errArg2 := js.ConsumeArgument(cbCtx, "timeout", codec.ZeroValue, codec.DecodeDuration)
	cbCtx.ConsumeArg()
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.SetInterval(handler, timeout)
	return encodeTaskHandle(cbCtx, result)
}

func WindowOrWorkerGlobalScope_clearInterval[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.WindowOrWorkerGlobalScope](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	id, errArg1 := js.ConsumeArgument(cbCtx, "id", codec.ZeroValue, decodeTaskHandle)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.ClearInterval(id)
	return nil, nil
}

func WindowOrWorkerGlobalScope_queueMicrotask[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[htmlinterfaces.WindowOrWorkerGlobalScope](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	callback, errArg1 := js.ConsumeArgument(cbCtx, "callback", nil, decodeVoidFunction)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.QueueMicrotask(callback)
	return nil, nil
}
