// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
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
	jsClass.CreateOperation("submit", HTMLFormElement_submit)
	jsClass.CreateOperation("requestSubmit", HTMLFormElement_requestSubmit)
	jsClass.CreateOperation("reset", HTMLFormElement_reset)
	jsClass.CreateOperation("checkValidity", HTMLFormElement_checkValidity)
	jsClass.CreateOperation("reportValidity", HTMLFormElement_reportValidity)
	jsClass.CreateAttribute("acceptCharset", HTMLFormElement_acceptCharset, HTMLFormElement_setAcceptCharset)
	jsClass.CreateAttribute("action", HTMLFormElement_action, HTMLFormElement_setAction)
	jsClass.CreateAttribute("autocomplete", HTMLFormElement_autocomplete, HTMLFormElement_setAutocomplete)
	jsClass.CreateAttribute("enctype", HTMLFormElement_enctype, HTMLFormElement_setEnctype)
	jsClass.CreateAttribute("encoding", HTMLFormElement_encoding, HTMLFormElement_setEncoding)
	jsClass.CreateAttribute("method", HTMLFormElement_method, HTMLFormElement_setMethod)
	jsClass.CreateAttribute("target", HTMLFormElement_target, HTMLFormElement_setTarget)
	jsClass.CreateAttribute("rel", HTMLFormElement_rel, HTMLFormElement_setRel)
	jsClass.CreateAttribute("relList", HTMLFormElement_relList, nil)
	jsClass.CreateAttribute("elements", HTMLFormElement_elements, nil)
	jsClass.CreateAttribute("length", HTMLFormElement_length, nil)
}

func HTMLFormElementConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func HTMLFormElement_submit[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	errCall := instance.Submit()
	return nil, errCall
}

func HTMLFormElement_requestSubmit[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func HTMLFormElement_reset[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_reset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_checkValidity[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_checkValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_reportValidity[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_reportValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_acceptCharset[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_acceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_setAcceptCharset[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_setAcceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_action[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Action()
	return codec.EncodeString(cbCtx, result)
}

func HTMLFormElement_setAction[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetAction(val)
	return nil, nil
}

func HTMLFormElement_autocomplete[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_autocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_setAutocomplete[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_setAutocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_enctype[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_enctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_setEnctype[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_setEnctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_encoding[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_encoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_setEncoding[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_setEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_method[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Method()
	return codec.EncodeString(cbCtx, result)
}

func HTMLFormElement_setMethod[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetMethod(val)
	return nil, nil
}

func HTMLFormElement_target[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_target: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_setTarget[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_setTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_rel[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_rel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_setRel[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_setRel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_relList[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_relList: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLFormElement_elements[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Elements()
	return codec.EncodeEntity(cbCtx, result)
}

func HTMLFormElement_length[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLFormElement.HTMLFormElement_length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
