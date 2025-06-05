package v8host

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func installGlobals(windowTemplate *v8.ObjectTemplate, host *V8ScriptHost) {
	// location := host.globals.namedGlobals["Location"]
	// windowTemplate.Set("location", location.ft.InstanceTemplate())
}

func (w *windowV8Wrapper) window(cbCtx jsCallbackContext) (jsValue, error) {
	return cbCtx.This(), nil
}

func (w *windowV8Wrapper) history(cbCtx jsCallbackContext) (jsValue, error) {
	win, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return cbCtx.Scope().Constructor("History").NewInstance(win.History())
}
