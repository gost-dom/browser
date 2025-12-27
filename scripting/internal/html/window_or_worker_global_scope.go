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

func (w WindowOrWorkerGlobalScope[T]) decodeTimerHandler(
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

func (w WindowOrWorkerGlobalScope[T]) decodeVoidFunction(
	scope js.Scope[T], v js.Value[T],
) (res htmlinterfaces.TimerHandler, err error) {
	return w.decodeTimerHandler(scope, v)
}

func (w WindowOrWorkerGlobalScope[T]) setTimeout(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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

func (w WindowOrWorkerGlobalScope[T]) setInterval(
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: queueMicrotask", js.ThisLogAttr(cbCtx))
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

func (w WindowOrWorkerGlobalScope[T]) clearTimeout(
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: clearTimeout", js.ThisLogAttr(cbCtx))
	handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
	if err == nil {
		cbCtx.Clock().Cancel(clock.TaskHandle(handle))
	}
	return nil, nil
}

func (w WindowOrWorkerGlobalScope[T]) clearInterval(
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: clearInterval", js.ThisLogAttr(cbCtx))
	handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
	if err == nil {
		cbCtx.Clock().Cancel(clock.TaskHandle(handle))
	}
	return nil, err
}

func (w WindowOrWorkerGlobalScope[T]) queueMicrotask(
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: queueMicrotask", js.ThisLogAttr(cbCtx))
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
