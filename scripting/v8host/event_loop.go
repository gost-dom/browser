package v8host

import (
	"errors"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	v8 "github.com/gost-dom/v8go"
)

func installEventLoopGlobals(host *V8ScriptHost, globalObjectTemplate *v8.ObjectTemplate) {
	globalObjectTemplate.Set(
		"queueMicrotask",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			f, err := consumeArgument(cbCtx, "callback", nil, decodeFunction)
			if err == nil {
				clock := cbCtx.Scope().Clock()
				clock.AddSafeMicrotask(func() {
					if _, err := f.Call(cbCtx.Scope().GlobalThis()); err != nil {
						UnhandledError(cbCtx.Scope(), err)
					}
				})
			}
			return nil, err
		}),
	)
	globalObjectTemplate.Set(
		"setTimeout",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			f, err1 := consumeArgument(cbCtx, "callback", nil, decodeFunction)
			delay, err2 := consumeArgument(cbCtx, "delay", nil, decodeInt32)
			err := errors.Join(err1, err2)
			if err != nil {
				return nil, err
			}
			clock := cbCtx.Scope().Clock()
			handle := clock.AddSafeTask(
				func() {
					if _, err := f.Call(cbCtx.Scope().GlobalThis()); err != nil {
						UnhandledError(cbCtx.Scope(), err)
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
			handle, err := consumeArgument(cbCtx, "handle", nil, decodeUint32)
			if err == nil {
				cbCtx.Scope().Clock().Cancel(clock.TaskHandle(handle))
			}
			return nil, nil
		}),
	)
	globalObjectTemplate.Set(
		"setInterval",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			f, err1 := consumeArgument(cbCtx, "callback", nil, decodeFunction)
			delay, err2 := consumeArgument(cbCtx, "delay", nil, decodeInt32)
			err := errors.Join(err1, err2)
			if err != nil {
				return nil, err
			}
			handle := cbCtx.Scope().Clock().SetInterval(
				func() {
					if _, err := f.Call(cbCtx.Scope().GlobalThis()); err != nil {
						UnhandledError(cbCtx.Scope(), err)
					}
				},
				time.Duration(delay)*time.Millisecond,
			)
			return encodeUint32(cbCtx, uint32(handle))
		}),
	)
	globalObjectTemplate.Set(
		"clearInterval",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			handle, err := consumeArgument(cbCtx, "handle", nil, decodeUint32)
			if err == nil {
				cbCtx.Scope().Clock().Cancel(clock.TaskHandle(handle))
			}
			return nil, err
		}),
	)
}
