package html

import (
	"errors"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func Initialize[T any](host js.ScriptEngine[T]) {
	installEventLoopGlobals(host)
}

func InitBuilder[T any](reg js.ClassBuilder[T]) {
	Bootstrap(reg)
	js.RegisterClass(reg, "DOMStringMap", "", NewDOMStringMap)
}

func installEventLoopGlobals[T any](host js.ScriptEngine[T]) {
	host.CreateFunction(
		"queueMicrotask",
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
			f, err := js.ConsumeArgument(cbCtx, "callback", nil, codec.DecodeFunction)
			if err == nil {
				clock := cbCtx.Scope().Clock()
				clock.AddSafeMicrotask(func() {
					if _, err := f.Call(cbCtx.Scope().GlobalThis()); err != nil {
						js.UnhandledError(cbCtx.Scope(), err)
					}
				})
			}
			return nil, err
		})
	host.CreateFunction(
		"setTimeout",
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
			f, err1 := js.ConsumeArgument(cbCtx, "callback", nil, codec.DecodeFunction)
			delay, err2 := js.ConsumeArgument(cbCtx, "delay", nil, codec.DecodeInt)
			err := errors.Join(err1, err2)
			if err != nil {
				return nil, err
			}
			clock := cbCtx.Scope().Clock()
			handle := clock.AddSafeTask(
				func() {
					if _, err := f.Call(cbCtx.Scope().GlobalThis()); err != nil {
						js.UnhandledError(cbCtx.Scope(), err)
					}
				},
				time.Duration(delay)*time.Millisecond,
			)
			return cbCtx.ValueFactory().NewUint32(uint32(handle)), nil
		})
	host.CreateFunction(
		"clearTimeout",
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
			handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
			if err == nil {
				cbCtx.Scope().Clock().Cancel(clock.TaskHandle(handle))
			}
			return nil, nil
		})
	host.CreateFunction(
		"setInterval",
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
			f, err1 := js.ConsumeArgument(cbCtx, "callback", nil, codec.DecodeFunction)
			delay, err2 := js.ConsumeArgument(cbCtx, "delay", nil, codec.DecodeInt)
			err := errors.Join(err1, err2)
			if err != nil {
				return nil, err
			}
			handle := cbCtx.Scope().Clock().SetInterval(
				func() {
					if _, err := f.Call(cbCtx.Scope().GlobalThis()); err != nil {
						js.UnhandledError(cbCtx.Scope(), err)
					}
				},
				time.Duration(delay)*time.Millisecond,
			)
			return codec.EncodeInt(cbCtx, int(handle))
		})
	host.CreateFunction(
		"clearInterval",
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
			handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
			if err == nil {
				cbCtx.Scope().Clock().Cancel(clock.TaskHandle(handle))
			}
			return nil, err
		})
}
