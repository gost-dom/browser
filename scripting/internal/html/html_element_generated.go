// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLElementV8Wrapper[T any] struct {
	htmlOrSVGElement *HTMLOrSVGElementV8Wrapper[T]
}

func NewHTMLElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *HTMLElementV8Wrapper[T] {
	return &HTMLElementV8Wrapper[T]{NewHTMLOrSVGElementV8Wrapper(scriptHost)}
}

func (wrapper HTMLElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("click", w.click)
	w.htmlOrSVGElement.installPrototype(jsClass)
}

func (w HTMLElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLElementV8Wrapper[T]) click(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.click")
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Click()
	return nil, nil
}
