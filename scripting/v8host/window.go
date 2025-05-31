package v8host

import (
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
	windowTemplate.Set("location", location.ft.InstanceTemplate())
}

func (w *windowV8Wrapper) window(cbCtx jsCallbackContext) (jsValue, error) {
	return cbCtx.This(), nil
}

func (w *windowV8Wrapper) history(cbCtx *v8CallbackContext) (jsValue, error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	ctx := cbCtx.ScriptCtx()
	return ctx.getConstructor("History").NewInstance(ctx, win.History())
}
