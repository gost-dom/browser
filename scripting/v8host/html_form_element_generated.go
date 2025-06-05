// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("HTMLFormElement", "HTMLElement", newHTMLFormElementV8Wrapper)
}

type htmlFormElementV8Wrapper[T any] struct {
	handleReffedObject[html.HTMLFormElement, T]
}

func newHTMLFormElementV8Wrapper(scriptHost jsScriptEngine) *htmlFormElementV8Wrapper[jsTypeParam] {
	return &htmlFormElementV8Wrapper[jsTypeParam]{newHandleReffedObject[html.HTMLFormElement, jsTypeParam](scriptHost)}
}

func (wrapper htmlFormElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w htmlFormElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w htmlFormElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlFormElementV8Wrapper[T]) submit(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.submit")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	errCall := instance.Submit()
	return nil, errCall
}

func (w htmlFormElementV8Wrapper[T]) requestSubmit(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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

func (w htmlFormElementV8Wrapper[T]) reset(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.reset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.reset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) checkValidity(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.checkValidity")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.checkValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) reportValidity(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.reportValidity")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.reportValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) acceptCharset(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.acceptCharset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.acceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) setAcceptCharset(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setAcceptCharset")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setAcceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) action(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.action")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Action()
	return w.toString_(cbCtx, result)
}

func (w htmlFormElementV8Wrapper[T]) setAction(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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

func (w htmlFormElementV8Wrapper[T]) autocomplete(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.autocomplete")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.autocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) setAutocomplete(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setAutocomplete")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setAutocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) enctype(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.enctype")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.enctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) setEnctype(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setEnctype")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setEnctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) encoding(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.encoding")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.encoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) setEncoding(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setEncoding")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) method(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.method")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Method()
	return w.toString_(cbCtx, result)
}

func (w htmlFormElementV8Wrapper[T]) setMethod(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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

func (w htmlFormElementV8Wrapper[T]) target(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.target")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.target: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) setTarget(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setTarget")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) rel(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.rel")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.rel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) setRel(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.setRel")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.setRel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) relList(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.relList")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.relList: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlFormElementV8Wrapper[T]) elements(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.elements")
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Elements()
	return encodeEntity(cbCtx, result)
}

func (w htmlFormElementV8Wrapper[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLFormElement.length")
	return cbCtx.ReturnWithError(errors.New("HTMLFormElement.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
