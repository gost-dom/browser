// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerJSClass("HTMLAnchorElement", "HTMLElement", createHTMLAnchorElementPrototype)
}

type htmlAnchorElementV8Wrapper struct {
	handleReffedObject[html.HTMLAnchorElement, jsTypeParam]
	htmlHyperlinkElementUtils *htmlHyperlinkElementUtilsV8Wrapper
}

func newHTMLAnchorElementV8Wrapper(scriptHost *V8ScriptHost) *htmlAnchorElementV8Wrapper {
	return &htmlAnchorElementV8Wrapper{
		newHandleReffedObject[html.HTMLAnchorElement](scriptHost),
		newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost),
	}
}

func createHTMLAnchorElementPrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newHTMLAnchorElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}

func (w htmlAnchorElementV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeAttribute("target", w.target, w.setTarget)
	w.htmlHyperlinkElementUtils.installPrototype(jsClass)
}

func (w htmlAnchorElementV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlAnchorElementV8Wrapper) target(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.target")
	instance, err := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target()
	return w.toString_(cbCtx, result)
}

func (w htmlAnchorElementV8Wrapper) setTarget(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.setTarget")
	instance, err0 := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetTarget(val)
	return cbCtx.ReturnWithValue(nil)
}
