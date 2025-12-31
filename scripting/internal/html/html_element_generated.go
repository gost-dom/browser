// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeHTMLElement[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("click", HTMLElement_click)
	HTMLElementCustomInitializer(jsClass)
	InitializeHTMLOrSVGElement(jsClass)
}

func HTMLElement_click[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Click()
	return nil, nil
}
