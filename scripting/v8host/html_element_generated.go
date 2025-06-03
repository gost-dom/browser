// This file is generated. Do not edit.

package v8host

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("HTMLElement", "Element", createHTMLElementPrototype)
}

type htmlElementV8Wrapper struct {
	handleReffedObject[html.HTMLElement, jsTypeParam]
	htmlOrSVGElement *htmlOrSVGElementV8Wrapper
}

func newHTMLElementV8Wrapper(scriptHost *V8ScriptHost) *htmlElementV8Wrapper {
	return &htmlElementV8Wrapper{
		newHandleReffedObject[html.HTMLElement](scriptHost),
		newHTMLOrSVGElementV8Wrapper(scriptHost),
	}
}

func createHTMLElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newHTMLElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor)

	return constructor
}
func (w htmlElementV8Wrapper) installPrototype(ft *v8.FunctionTemplate) {
	jsClass := newV8Class(w.scriptHost, ft)
	jsClass.CreatePrototypeMethod("click", w.click)
	w.htmlOrSVGElement.installPrototype(ft)
}

func (w htmlElementV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlElementV8Wrapper) click(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.click")
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Click()
	return nil, nil
}
