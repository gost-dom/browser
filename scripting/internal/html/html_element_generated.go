// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLElement[T any] struct {
	htmlOrSVGElement HTMLOrSVGElement[T]
}

func NewHTMLElement[T any](scriptHost js.ScriptEngine[T]) HTMLElement[T] {
	return HTMLElement[T]{NewHTMLOrSVGElement(scriptHost)}
}

func (wrapper HTMLElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w HTMLElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("click", HTMLElement_click)
	w.htmlOrSVGElement.installPrototype(jsClass)
}

func HTMLElementConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func HTMLElement_click[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Click()
	return nil, nil
}
