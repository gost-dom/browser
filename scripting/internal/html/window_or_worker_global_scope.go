package html

import (
	"errors"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func decodeTimerHandler[T any](
	scope js.Scope[T], v js.Value[T],
) (res htmlinterfaces.TimerHandler, err error) {
	f, err := codec.DecodeFunction(scope, v)
	if err != nil {
		return nil, err
	}
	res = func() {
		if _, err := f.Call(scope.GlobalThis()); err != nil {
			dom.HandleJSCallbackError(scope, "TimerHandler", err)
		}
	}
	return res, nil
}

func WindowOrWorkerGlobalScope_decodeVoidFunction[T any](
	scope js.Scope[T], v js.Value[T],
) (res htmlinterfaces.TimerHandler, err error) {
	return decodeTimerHandler(scope, v)
}

func WindowOrWorkerGlobalScope_setTimeout[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	f, err1 := js.ConsumeArgument(cbCtx, "callback", nil, codec.DecodeFunction)
	delay, err2 := js.ConsumeArgument(cbCtx, "delay", codec.ZeroValue, codec.DecodeInt)
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	clock := cbCtx.Clock()
	handle := clock.AddSafeTask(
		func() {
			if _, err := f.Call(cbCtx.GlobalThis()); err != nil {
				dom.HandleJSCallbackError(cbCtx, "setTimeout", err)
			}
		},
		time.Duration(delay)*time.Millisecond,
	)
	return cbCtx.NewUint32(uint32(handle)), nil
}

func WindowOrWorkerGlobalScope_setInterval[T any](
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	f, err1 := js.ConsumeArgument(cbCtx, "callback", nil, codec.DecodeFunction)
	delay, err2 := js.ConsumeArgument(cbCtx, "delay", nil, codec.DecodeInt)
	err := errors.Join(err1, err2)
	if err != nil {
		return nil, err
	}
	handle := cbCtx.Clock().SetInterval(
		func() {
			if _, err := f.Call(cbCtx.GlobalThis()); err != nil {
				dom.HandleJSCallbackError(cbCtx, "SetInterval", err)
			}
		},
		time.Duration(delay)*time.Millisecond,
	)
	return codec.EncodeInt(cbCtx, int(handle))
}

func WindowOrWorkerGlobalScope_clearTimeout[T any](
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
	if err == nil {
		cbCtx.Clock().Cancel(clock.TaskHandle(handle))
	}
	return nil, nil
}

func WindowOrWorkerGlobalScope_clearInterval[T any](
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
	if err == nil {
		cbCtx.Clock().Cancel(clock.TaskHandle(handle))
	}
	return nil, err
}

func WindowOrWorkerGlobalScope_queueMicrotask[T any](
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	f, err := js.ConsumeArgument(cbCtx, "callback", nil, codec.DecodeFunction)
	if err == nil {
		clock := cbCtx.Clock()
		clock.AddSafeMicrotask(func() {
			if _, err := f.Call(cbCtx.GlobalThis()); err != nil {
				dom.HandleJSCallbackError(cbCtx, "Microtask", err)
			}
		})
	}
	return nil, err
}
