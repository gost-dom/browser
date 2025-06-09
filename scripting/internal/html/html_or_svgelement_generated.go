// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLOrSVGElementV8Wrapper[T any] struct{}

func NewHTMLOrSVGElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *HTMLOrSVGElementV8Wrapper[T] {
	return &HTMLOrSVGElementV8Wrapper[T]{}
}

func (wrapper HTMLOrSVGElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLOrSVGElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("focus", w.focus)
	jsClass.CreatePrototypeMethod("blur", w.blur)
	jsClass.CreatePrototypeAttribute("dataset", w.dataset, nil)
	jsClass.CreatePrototypeAttribute("nonce", w.nonce, w.setNonce)
	jsClass.CreatePrototypeAttribute("autofocus", w.autofocus, w.setAutofocus)
	jsClass.CreatePrototypeAttribute("tabIndex", w.tabIndex, w.setTabIndex)
}

func (w HTMLOrSVGElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLOrSVGElementV8Wrapper[T]) blur(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.blur")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Blur()
	return nil, nil
}

func (w HTMLOrSVGElementV8Wrapper[T]) dataset(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.dataset")
	return nil, errors.New("HTMLOrSVGElement.dataset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLOrSVGElementV8Wrapper[T]) nonce(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.nonce")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Nonce()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLOrSVGElementV8Wrapper[T]) setNonce(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.setNonce")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetNonce(val)
	return nil, nil
}

func (w HTMLOrSVGElementV8Wrapper[T]) autofocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.autofocus")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Autofocus()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w HTMLOrSVGElementV8Wrapper[T]) setAutofocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.setAutofocus")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetAutofocus(val)
	return nil, nil
}

func (w HTMLOrSVGElementV8Wrapper[T]) tabIndex(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.tabIndex")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TabIndex()
	return codec.EncodeInt(cbCtx, result)
}

func (w HTMLOrSVGElementV8Wrapper[T]) setTabIndex(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.setTabIndex")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeInt)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTabIndex(val)
	return nil, nil
}
