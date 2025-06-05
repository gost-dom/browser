// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("HTMLInputElement", "HTMLElement", newHTMLInputElementV8Wrapper)
}

type htmlInputElementV8Wrapper struct {
	handleReffedObject[html.HTMLInputElement, jsTypeParam]
}

func newHTMLInputElementV8Wrapper(scriptHost jsScriptEngine) *htmlInputElementV8Wrapper {
	return &htmlInputElementV8Wrapper{newHandleReffedObject[html.HTMLInputElement](scriptHost)}
}

func (wrapper htmlInputElementV8Wrapper) Initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
}

func (w htmlInputElementV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeMethod("checkValidity", w.checkValidity)
	jsClass.CreatePrototypeAttribute("type", w.type_, w.setType)
}

func (w htmlInputElementV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlInputElementV8Wrapper) checkValidity(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.checkValidity")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.CheckValidity()
	return w.toBoolean(cbCtx, result)
}

func (w htmlInputElementV8Wrapper) type_(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.type_")
	instance, err := js.As[html.HTMLInputElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type()
	return w.toString_(cbCtx, result)
}

func (w htmlInputElementV8Wrapper) setType(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLInputElement.setType")
	instance, err0 := js.As[html.HTMLInputElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetType(val)
	return cbCtx.ReturnWithValue(nil)
}
