// This file is generated. Do not edit.

package v8host

import (
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("HTMLElement", "Element", newHTMLElementV8Wrapper)
}

type htmlElementV8Wrapper struct {
	handleReffedObject[html.HTMLElement, jsTypeParam]
	htmlOrSVGElement *htmlOrSVGElementV8Wrapper
}

func newHTMLElementV8Wrapper(scriptHost jsScriptEngine) *htmlElementV8Wrapper {
	return &htmlElementV8Wrapper{
		newHandleReffedObject[html.HTMLElement](scriptHost),
		newHTMLOrSVGElementV8Wrapper(scriptHost),
	}
}

func (wrapper htmlElementV8Wrapper) initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
}

func (w htmlElementV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeMethod("click", w.click)
	w.htmlOrSVGElement.installPrototype(jsClass)
}

func (w htmlElementV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlElementV8Wrapper) click(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLElement.click")
	instance, err := js.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.Click()
	return nil, nil
}
