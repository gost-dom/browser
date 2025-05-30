// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("HTMLFormElement", "HTMLElement", createHTMLFormElementPrototype)
}

type htmlFormElementV8Wrapper struct {
	handleReffedObject[html.HTMLFormElement]
}

func newHTMLFormElementV8Wrapper(scriptHost *V8ScriptHost) *htmlFormElementV8Wrapper {
	return &htmlFormElementV8Wrapper{newHandleReffedObject[html.HTMLFormElement](scriptHost)}
}

func createHTMLFormElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newHTMLFormElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlFormElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("submit", wrapV8Callback(w.scriptHost, w.submit))
	prototypeTmpl.Set("requestSubmit", wrapV8Callback(w.scriptHost, w.requestSubmit))
	prototypeTmpl.Set("reset", wrapV8Callback(w.scriptHost, w.reset))
	prototypeTmpl.Set("checkValidity", wrapV8Callback(w.scriptHost, w.checkValidity))
	prototypeTmpl.Set("reportValidity", wrapV8Callback(w.scriptHost, w.reportValidity))

	prototypeTmpl.SetAccessorProperty("acceptCharset",
		wrapV8Callback(w.scriptHost, w.acceptCharset),
		wrapV8Callback(w.scriptHost, w.setAcceptCharset),
		v8.None)
	prototypeTmpl.SetAccessorProperty("action",
		wrapV8Callback(w.scriptHost, w.action),
		wrapV8Callback(w.scriptHost, w.setAction),
		v8.None)
	prototypeTmpl.SetAccessorProperty("autocomplete",
		wrapV8Callback(w.scriptHost, w.autocomplete),
		wrapV8Callback(w.scriptHost, w.setAutocomplete),
		v8.None)
	prototypeTmpl.SetAccessorProperty("enctype",
		wrapV8Callback(w.scriptHost, w.enctype),
		wrapV8Callback(w.scriptHost, w.setEnctype),
		v8.None)
	prototypeTmpl.SetAccessorProperty("encoding",
		wrapV8Callback(w.scriptHost, w.encoding),
		wrapV8Callback(w.scriptHost, w.setEncoding),
		v8.None)
	prototypeTmpl.SetAccessorProperty("method",
		wrapV8Callback(w.scriptHost, w.method),
		wrapV8Callback(w.scriptHost, w.setMethod),
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		wrapV8Callback(w.scriptHost, w.target),
		wrapV8Callback(w.scriptHost, w.setTarget),
		v8.None)
	prototypeTmpl.SetAccessorProperty("rel",
		wrapV8Callback(w.scriptHost, w.rel),
		wrapV8Callback(w.scriptHost, w.setRel),
		v8.None)
	prototypeTmpl.SetAccessorProperty("relList",
		wrapV8Callback(w.scriptHost, w.relList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("elements",
		wrapV8Callback(w.scriptHost, w.elements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
}

func (w htmlFormElementV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlFormElementV8Wrapper) submit(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.submit")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Submit()
	return nil, errCall
}

func (w htmlFormElementV8Wrapper) requestSubmit(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.requestSubmit")
	instance, errInst := js.As[html.HTMLFormElement](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	submitter, errArg1 := consumeArgument(cbCtx, "submitter", w.defaultHTMLElement, w.decodeHTMLElement)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.RequestSubmit(submitter)
	return nil, errCall
}

func (w htmlFormElementV8Wrapper) reset(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.reset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.reset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) checkValidity(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.checkValidity")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.checkValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) reportValidity(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.reportValidity")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.reportValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) acceptCharset(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.acceptCharset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.acceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setAcceptCharset(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setAcceptCharset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setAcceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) action(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.action")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Action()
	return w.toString_(cbCtx, result)
}

func (w htmlFormElementV8Wrapper) setAction(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setAction")
	instance, err0 := js.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetAction(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlFormElementV8Wrapper) autocomplete(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.autocomplete")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.autocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setAutocomplete(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setAutocomplete")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setAutocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) enctype(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.enctype")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.enctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setEnctype(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setEnctype")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setEnctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) encoding(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.encoding")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.encoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setEncoding(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setEncoding")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) method(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.method")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Method()
	return w.toString_(cbCtx, result)
}

func (w htmlFormElementV8Wrapper) setMethod(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setMethod")
	instance, err0 := js.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetMethod(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlFormElementV8Wrapper) target(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.target")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.target: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setTarget(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setTarget")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) rel(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.rel")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.rel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setRel(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.setRel")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setRel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) relList(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.relList")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.relList: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) elements(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.elements")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Elements()
	return w.toJSWrapper(cbCtx, result)
}

func (w htmlFormElementV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: HTMLFormElement.length")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
