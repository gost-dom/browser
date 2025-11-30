// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLFormElement[T any] struct{}

func NewHTMLFormElement[T any](scriptHost js.ScriptEngine[T]) *HTMLFormElement[T] {
	return &HTMLFormElement[T]{}
}

func (wrapper HTMLFormElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLFormElement[T]) installPrototype(jsClass js.Class[T]) {
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

func (w HTMLFormElement[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLFormElement[T]) submit(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Submit()
	return nil, errCall
}

func (w HTMLFormElement[T]) requestSubmit(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLFormElement](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	submitter, errArg1 := js.ConsumeArgument(cbCtx, "submitter", codec.ZeroValue, codec.DecodeHTMLElement)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.RequestSubmit(submitter)
	return nil, errCall
}

func (w HTMLFormElement[T]) reset(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.reset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) checkValidity(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.checkValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) reportValidity(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.reportValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) acceptCharset(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.acceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) setAcceptCharset(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.setAcceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) action(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Action()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLFormElement[T]) setAction(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetAction(val)
	return nil, nil
}

func (w HTMLFormElement[T]) autocomplete(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.autocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) setAutocomplete(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.setAutocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) enctype(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.enctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) setEnctype(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.setEnctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) encoding(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.encoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) setEncoding(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.setEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) method(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Method()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLFormElement[T]) setMethod(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetMethod(val)
	return nil, nil
}

func (w HTMLFormElement[T]) target(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.target: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) setTarget(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.setTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) rel(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.rel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) setRel(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.setRel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) relList(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.relList: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLFormElement[T]) elements(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Elements()
	return codec.EncodeEntity(cbCtx, result)
}

func (w HTMLFormElement[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
