// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
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
	jsClass.CreatePrototypeAttribute("target", w.target, w.setTarget)
	w.htmlHyperlinkElementUtils.installPrototype(jsClass)
}

func (w HTMLAnchorElement[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLAnchorElement[T]) target(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.target")
	instance, err := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return codec.EncodeString(cbCtx, result)
}

func (w HTMLAnchorElement[T]) setTarget(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.setTarget")
	instance, err0 := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTarget(val)
	return nil, nil
}
