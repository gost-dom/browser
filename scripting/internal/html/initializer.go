package html

import (
	"errors"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func Initialize[T any](host js.ScriptEngine[T]) {
	installEventLoopGlobals(host)
}

func InitBuilder[T any](e js.ScriptEngine[T]) {
	Bootstrap(e)
	js.RegisterClass(e, "DOMStringMap", "", NewDOMStringMap)
}

func QueueMicrotask[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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
func SetTimeout[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: setTimeout", js.ThisLogAttr(cbCtx))
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

func SetInterval[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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

func ClearTimeout[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: clearTimeout", js.ThisLogAttr(cbCtx))
	handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
	if err == nil {
		cbCtx.Clock().Cancel(clock.TaskHandle(handle))
	}
	return nil, nil
}

func ClearInterval[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: clearInterval", js.ThisLogAttr(cbCtx))
	handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
	if err == nil {
		cbCtx.Clock().Cancel(clock.TaskHandle(handle))
	}
	return nil, err
}

func installEventLoopGlobals[T any](host js.ScriptEngine[T]) {
	host.CreateFunction("queueMicrotask", QueueMicrotask[T])
	host.CreateFunction("setTimeout", SetTimeout[T])
	host.CreateFunction("clearTimeout", ClearTimeout[T])
	host.CreateFunction("setInterval", SetInterval[T])
	host.CreateFunction("clearInterval", ClearInterval[T])
}
