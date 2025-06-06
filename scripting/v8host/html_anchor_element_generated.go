// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type htmlAnchorElementV8Wrapper[T any] struct {
	htmlHyperlinkElementUtils *htmlHyperlinkElementUtilsV8Wrapper[T]
}

func newHTMLAnchorElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *htmlAnchorElementV8Wrapper[T] {
	return &htmlAnchorElementV8Wrapper[T]{newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost)}
}

func (wrapper htmlAnchorElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w htmlAnchorElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("target", w.target, w.setTarget)
	w.htmlHyperlinkElementUtils.installPrototype(jsClass)
}

func (w htmlAnchorElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlAnchorElementV8Wrapper[T]) target(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.target")
	instance, err := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target()
	return codec.EncodeString(cbCtx, result)
}

func (w htmlAnchorElementV8Wrapper[T]) setTarget(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLAnchorElement.setTarget")
	instance, err0 := js.As[html.HTMLAnchorElement](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetTarget(val)
	return cbCtx.ReturnWithValue(nil)
}
