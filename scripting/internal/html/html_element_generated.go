// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLElement[T any] struct {
	htmlOrSVGElement *HTMLOrSVGElement[T]
}

func NewHTMLElement[T any](scriptHost js.ScriptEngine[T]) *HTMLElement[T] {
	return &HTMLElement[T]{NewHTMLOrSVGElement(scriptHost)}
}

func (wrapper HTMLElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("click", w.click)
	w.htmlOrSVGElement.installPrototype(jsClass)
}

func (w HTMLElement[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLElement.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLElement.Constructor", js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLElement[T]) click(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: HTMLElement.click - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: HTMLElement.click", js.LogAttr("res", res))
	}()
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Click()
	return nil, nil
}
