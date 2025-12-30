// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLAnchorElement[T any] struct {
	htmlHyperlinkElementUtils *HTMLHyperlinkElementUtils[T]
}

func NewHTMLAnchorElement[T any](scriptHost js.ScriptEngine[T]) *HTMLAnchorElement[T] {
	return &HTMLAnchorElement[T]{NewHTMLHyperlinkElementUtils(scriptHost)}
}

func (wrapper HTMLAnchorElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLAnchorElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateAttribute("target", HTMLAnchorElement_target, HTMLAnchorElement_setTarget)
	w.htmlHyperlinkElementUtils.installPrototype(jsClass)
}

func HTMLAnchorElementConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func HTMLAnchorElement_target[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return codec.EncodeString(cbCtx, result)
}

func HTMLAnchorElement_setTarget[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTarget(val)
	return nil, nil
}
