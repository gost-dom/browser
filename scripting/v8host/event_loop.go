package v8host

import (
	"log/slog"
	"runtime/debug"

	"github.com/gost-dom/browser/clock"
	"github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

type workItem struct {
	fn *v8.Function
}

type eventLoop struct {
	clock   *clock.Clock
	ctx     *V8ScriptContext
	errorCb func(error)
}

func (l *eventLoop) tick() error {
	return l.clock.Tick()
}

func newWorkItem(fn *v8.Function) workItem {
	return workItem{fn}
}

// dispatch places an item on the event loop to be executed immediately
func (l *eventLoop) dispatch(task clock.TaskCallback) {
	l.clock.AddSafeTask(clock.Immediate, func() {
		if err := task(); err != nil { //w.fn.Call(l.globalObject); err != nil {
			l.errorCb(err)
		}
	})
}

func newEventLoop(context *V8ScriptContext, cb func(error)) *eventLoop {
	return &eventLoop{clock.New(), context, cb}
}

func installEventLoopGlobals(host *V8ScriptHost, globalObjectTemplate *v8.ObjectTemplate) {
	iso := host.iso

	globalObjectTemplate.Set(
		"setTimeout",
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				helper := newArgumentHelper(host, info)
				f, err1 := helper.getFunctionArg(0)
				// delay, err2 := helper.GetInt32Arg(1)
				if err1 == nil {
					ctx.eventLoop.dispatch(func() error {
						_, err := f.Call(info.Context().Global())
						if err != nil {
							log.Error(
								"EventLoop: Error",
								slog.String("script", f.String()),
								slog.String("error", err.Error()),
								slog.String("stack", string(debug.Stack())),
							)
						}

						return err

					})
				}
				// TODO: Return a cancel token
				return v8.Undefined(iso), err1
			},
		),
	)
}
