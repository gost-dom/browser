package v8host

import (
	"runtime/cgo"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func installGlobals(
	windowFnTemplate *v8.FunctionTemplate,
	host *V8ScriptHost,
	globalInstalls []globalInstall,
) {
	windowTemplate := windowFnTemplate.InstanceTemplate()
	for _, globalInstall := range globalInstalls {
		windowTemplate.Set(globalInstall.name, globalInstall.constructor)
	}
	location := host.globals.namedGlobals["Location"]
	windowTemplate.Set("location", location.InstanceTemplate())
}

func (w *windowV8Wrapper) window(cbCtx *argumentHelper) (*v8.Value, error) {
	return cbCtx.This().Value, nil
}

func (w *windowV8Wrapper) history(cbCtx *argumentHelper) (*v8.Value, error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	ctx := cbCtx.ScriptCtx()
	history, err := w.scriptHost.globals.namedGlobals["History"].InstanceTemplate().
		NewInstance(ctx.v8ctx)
	if err != nil {
		return nil, err
	}
	handle := cgo.NewHandle(win.History())
	ctx.addDisposer(handleDisposable(handle))
	internal := v8.NewValueExternalHandle(w.iso(), handle)
	history.SetInternalField(0, internal)
	return history.Value, nil
}
