package v8host

import (
	"errors"
	"time"

	"github.com/gost-dom/browser/internal/clock"
	v8 "github.com/gost-dom/v8go"
)

type workItem struct {
	fn *v8.Function
}

type eventLoop struct {
	ctx     *V8ScriptContext
	errorCb func(error)
}

func (l *eventLoop) tick() error {
	return l.ctx.clock.Tick()
}

func newEventLoop(context *V8ScriptContext, cb func(error)) *eventLoop {
	return &eventLoop{context, cb}
}

func installEventLoopGlobals(host *V8ScriptHost, globalObjectTemplate *v8.ObjectTemplate) {
	globalObjectTemplate.Set(
		"queueMicrotask",
		wrapV8Callback(host, func(cbCtx jsCallbackContext) (jsValue, error) {
			f, err := consumeArgument(cbCtx, "callback", nil, decodeFunction)
			if err == nil {
				ctx := cbCtx.ScriptCtx()
				clock := ctx.clock
				clock.AddSafeMicrotask(func() {
					if _, err := f.Call(cbCtx.Global()); err != nil {
						ctx.eventLoop.errorCb(err)
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
			ctx := cbCtx.ScriptCtx()
			handle := ctx.clock.AddSafeTask(
				func() {
					if _, err := f.Call(cbCtx.Global()); err != nil {
						ctx.eventLoop.errorCb(err)
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
			handle := cbCtx.consumeValue()
			ctx := cbCtx.ScriptCtx()
			ctx.clock.Cancel(clock.TaskHandle(handle.Uint32()))
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
			ctx := cbCtx.ScriptCtx()
			handle := ctx.clock.SetInterval(
				func() {
					if _, err := f.Call(cbCtx.Global()); err != nil {
						ctx.eventLoop.errorCb(err)
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
			ctx := cbCtx.ScriptCtx()
			if err == nil {
				ctx.clock.Cancel(clock.TaskHandle(handle))
			}
			return nil, err
		}),
	)
}
