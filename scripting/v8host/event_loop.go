package v8host

import (
	"errors"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func installEventLoopGlobals(host *V8ScriptHost, globalObjectTemplate *v8.ObjectTemplate) {
	globalObjectTemplate.Set(
		"queueMicrotask",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
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
		}),
	)
	globalObjectTemplate.Set(
		"setTimeout",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
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
		}),
	)
	globalObjectTemplate.Set(
		"clearTimeout",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
			if err == nil {
				cbCtx.Scope().Clock().Cancel(clock.TaskHandle(handle))
			}
			return nil, nil
		}),
	)
	globalObjectTemplate.Set(
		"setInterval",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
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
		}),
	)
	globalObjectTemplate.Set(
		"clearInterval",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			handle, err := js.ConsumeArgument(cbCtx, "handle", nil, codec.DecodeInt)
			if err == nil {
				cbCtx.Scope().Clock().Cancel(clock.TaskHandle(handle))
			}
			return nil, err
		}),
	)
}
