// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("HTMLFormElement", "HTMLElement", func(engine *V8ScriptHost) jsInitializer {
		return newHTMLFormElementV8Wrapper(engine)
	})
}

type htmlFormElementV8Wrapper struct {
	handleReffedObject[html.HTMLFormElement, jsTypeParam]
}

func newHTMLFormElementV8Wrapper(scriptHost *V8ScriptHost) *htmlFormElementV8Wrapper {
	return &htmlFormElementV8Wrapper{newHandleReffedObject[html.HTMLFormElement](scriptHost)}
}

func createHTMLFormElementPrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newHTMLFormElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper htmlFormElementV8Wrapper) initialize(jsClass v8Class) {
	wrapper.installPrototype(jsClass)
}

func (w htmlFormElementV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeMethod("submit", w.submit)
	jsClass.CreatePrototypeMethod("requestSubmit", w.requestSubmit)
	jsClass.CreatePrototypeMethod("reset", w.reset)
	jsClass.CreatePrototypeMethod("checkValidity", w.checkValidity)
	jsClass.CreatePrototypeMethod("reportValidity", w.reportValidity)
	jsClass.CreatePrototypeAttribute("acceptCharset", w.acceptCharset, w.setAcceptCharset)
	jsClass.CreatePrototypeAttribute("action", w.action, w.setAction)
	jsClass.CreatePrototypeAttribute("autocomplete", w.autocomplete, w.setAutocomplete)
	jsClass.CreatePrototypeAttribute("enctype", w.enctype, w.setEnctype)
	jsClass.CreatePrototypeAttribute("encoding", w.encoding, w.setEncoding)
	jsClass.CreatePrototypeAttribute("method", w.method, w.setMethod)
	jsClass.CreatePrototypeAttribute("target", w.target, w.setTarget)
	jsClass.CreatePrototypeAttribute("rel", w.rel, w.setRel)
	jsClass.CreatePrototypeAttribute("relList", w.relList, nil)
	jsClass.CreatePrototypeAttribute("elements", w.elements, nil)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w htmlFormElementV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlFormElementV8Wrapper) submit(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.submit")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Submit()
	return nil, errCall
}

func (w htmlFormElementV8Wrapper) requestSubmit(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.requestSubmit")
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
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.reset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.reset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) checkValidity(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.checkValidity")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.checkValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) reportValidity(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.reportValidity")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.reportValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) acceptCharset(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.acceptCharset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.acceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setAcceptCharset(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setAcceptCharset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setAcceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) action(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.action")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Action()
	return w.toString_(cbCtx, result)
}

func (w htmlFormElementV8Wrapper) setAction(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setAction")
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
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.autocomplete")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.autocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setAutocomplete(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setAutocomplete")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setAutocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) enctype(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.enctype")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.enctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setEnctype(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setEnctype")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setEnctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) encoding(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.encoding")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.encoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setEncoding(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setEncoding")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) method(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.method")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Method()
	return w.toString_(cbCtx, result)
}

func (w htmlFormElementV8Wrapper) setMethod(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setMethod")
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
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.target")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.target: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setTarget(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setTarget")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) rel(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.rel")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.rel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) setRel(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setRel")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setRel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) relList(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.relList")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.relList: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper) elements(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.elements")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Elements()
	return encodeEntity(cbCtx, result)
}

func (w htmlFormElementV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.length")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
