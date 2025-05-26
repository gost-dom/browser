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
	iso := host.iso

	globalObjectTemplate.Set(
		"queueMicrotask",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				f, err := helper.consumeFunction()
				if err == nil {
					ctx.clock.AddSafeMicrotask(func() {
						if _, err := f.Call(info.Context().Global()); err != nil {
							ctx.eventLoop.errorCb(err)
						}
					})
				}
				return nil, err
			},
		),
	)
	globalObjectTemplate.Set(
		"setTimeout",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				f, err1 := helper.consumeFunction()
				delay, err2 := helper.consumeInt32()
				err := errors.Join(err1, err2)
				if err != nil {
					return v8.Undefined(iso), err
				}
				handle := ctx.clock.AddSafeTask(
					func() {
						if _, err := f.Call(info.Context().Global()); err != nil {
							ctx.eventLoop.errorCb(err)
						}
					},
					time.Duration(delay)*time.Millisecond,
				)
				return v8.NewValue(iso, uint32(handle))
			},
		),
	)
	globalObjectTemplate.Set(
		"clearTimeout",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				handle := helper.consumeValue()
				ctx.clock.Cancel(clock.TaskHandle(handle.Uint32()))
				return nil, nil
			},
		),
	)
	globalObjectTemplate.Set(
		"setInterval",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				f, err1 := helper.consumeFunction()
				delay, err2 := consumeArgument(helper, "delay", nil, decodeInt32)
				err := errors.Join(err1, err2)
				if err != nil {
					return v8.Undefined(iso), err
				}
				handle := ctx.clock.SetInterval(
					func() {
						if _, err := f.Call(info.Context().Global()); err != nil {
							ctx.eventLoop.errorCb(err)
						}
					},
					time.Duration(delay)*time.Millisecond,
				)
				return v8.NewValue(iso, uint32(handle))
			},
		),
	)
	globalObjectTemplate.Set(
		"clearInterval",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				handle := helper.consumeValue()
				ctx.clock.Cancel(clock.TaskHandle(handle.Uint32()))
				return nil, nil
			},
		),
	)
}
