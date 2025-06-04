// This file is generated. Do not edit.

package v8host

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
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

func createHTMLElementPrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newHTMLElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}

func (w htmlElementV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeMethod("click", w.click)
	w.htmlOrSVGElement.installPrototype(jsClass)
}

func (w htmlElementV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.constructor")
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
