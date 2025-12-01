// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLOrSVGElement[T any] struct{}

func NewHTMLOrSVGElement[T any](scriptHost js.ScriptEngine[T]) *HTMLOrSVGElement[T] {
	return &HTMLOrSVGElement[T]{}
}

func (wrapper HTMLOrSVGElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLOrSVGElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("focus", w.focus)
	jsClass.CreatePrototypeMethod("blur", w.blur)
	jsClass.CreatePrototypeAttribute("dataset", w.dataset, nil)
	jsClass.CreatePrototypeAttribute("nonce", w.nonce, w.setNonce)
	jsClass.CreatePrototypeAttribute("autofocus", w.autofocus, w.setAutofocus)
	jsClass.CreatePrototypeAttribute("tabIndex", w.tabIndex, w.setTabIndex)
}

func (w HTMLOrSVGElement[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLOrSVGElement[T]) blur(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Blur()
	return nil, nil
}

func (w HTMLOrSVGElement[T]) dataset(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Dataset()
	return codec.EncodeEntity(cbCtx, result)
}

func (w HTMLOrSVGElement[T]) nonce(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Nonce()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLOrSVGElement[T]) setNonce(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetNonce(val)
	return nil, nil
}

func (w HTMLOrSVGElement[T]) autofocus(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Autofocus()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w HTMLOrSVGElement[T]) setAutofocus(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeBoolean)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetAutofocus(val)
	return nil, nil
}

func (w HTMLOrSVGElement[T]) tabIndex(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TabIndex()
	return codec.EncodeInt(cbCtx, result)
}

func (w HTMLOrSVGElement[T]) setTabIndex(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeInt)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTabIndex(val)
	return nil, nil
}
