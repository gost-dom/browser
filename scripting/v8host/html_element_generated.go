// This file is generated. Do not edit.

package v8host

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("HTMLElement", "Element", newHTMLElementV8Wrapper)
}

type htmlElementV8Wrapper[T any] struct {
	handleReffedObject[html.HTMLElement, T]
	htmlOrSVGElement *htmlOrSVGElementV8Wrapper[T]
}

func newHTMLElementV8Wrapper(scriptHost jsScriptEngine) *htmlElementV8Wrapper[jsTypeParam] {
	return &htmlElementV8Wrapper[jsTypeParam]{
		newHandleReffedObject[html.HTMLElement, jsTypeParam](scriptHost),
		newHTMLOrSVGElementV8Wrapper(scriptHost),
	}
}

func (wrapper htmlElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w htmlElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("click", w.click)
	w.htmlOrSVGElement.installPrototype(jsClass)
}

func (w htmlElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlElementV8Wrapper[T]) click(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.click")
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Click()
	return nil, nil
}
