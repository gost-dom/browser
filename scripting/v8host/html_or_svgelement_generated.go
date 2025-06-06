// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type htmlOrSVGElementV8Wrapper[T any] struct{}

func newHTMLOrSVGElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *htmlOrSVGElementV8Wrapper[T] {
	return &htmlOrSVGElementV8Wrapper[T]{}
}

func (wrapper htmlOrSVGElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w htmlOrSVGElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("focus", w.focus)
	jsClass.CreatePrototypeMethod("blur", w.blur)
	jsClass.CreatePrototypeAttribute("dataset", w.dataset, nil)
	jsClass.CreatePrototypeAttribute("nonce", w.nonce, w.setNonce)
	jsClass.CreatePrototypeAttribute("autofocus", w.autofocus, w.setAutofocus)
	jsClass.CreatePrototypeAttribute("tabIndex", w.tabIndex, w.setTabIndex)
}

func (w htmlOrSVGElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlOrSVGElementV8Wrapper[T]) blur(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.blur")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Blur()
	return nil, nil
}

func (w htmlOrSVGElementV8Wrapper[T]) dataset(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.dataset")
	return cbCtx.ReturnWithError(errors.New("HTMLOrSVGElement.dataset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlOrSVGElementV8Wrapper[T]) nonce(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.nonce")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Nonce()
	return codec.EncodeString(cbCtx, result)
}

func (w htmlOrSVGElementV8Wrapper[T]) setNonce(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.setNonce")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetNonce(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlOrSVGElementV8Wrapper[T]) autofocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.autofocus")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Autofocus()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w htmlOrSVGElementV8Wrapper[T]) setAutofocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.setAutofocus")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeBoolean)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetAutofocus(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w htmlOrSVGElementV8Wrapper[T]) tabIndex(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.tabIndex")
	instance, err := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.TabIndex()
	return codec.EncodeInt(cbCtx, result)
}

func (w htmlOrSVGElementV8Wrapper[T]) setTabIndex(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLOrSVGElement.setTabIndex")
	instance, err0 := js.As[html.HTMLOrSVGElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeInt)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetTabIndex(val)
	return cbCtx.ReturnWithValue(nil)
}
