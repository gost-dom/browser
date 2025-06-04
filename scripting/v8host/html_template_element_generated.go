// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("HTMLTemplateElement", "HTMLElement", newHTMLTemplateElementV8Wrapper)
}

type htmlTemplateElementV8Wrapper struct {
	handleReffedObject[html.HTMLTemplateElement, jsTypeParam]
}

func newHTMLTemplateElementV8Wrapper(scriptHost *V8ScriptHost) *htmlTemplateElementV8Wrapper {
	return &htmlTemplateElementV8Wrapper{newHandleReffedObject[html.HTMLTemplateElement](scriptHost)}
}

func createHTMLTemplateElementPrototype(scriptHost *V8ScriptHost) jsClass {
	wrapper := newHTMLTemplateElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper htmlTemplateElementV8Wrapper) initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
}

func (w htmlTemplateElementV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeAttribute("content", w.content, nil)
	jsClass.CreatePrototypeAttribute("shadowRootMode", w.shadowRootMode, w.setShadowRootMode)
	jsClass.CreatePrototypeAttribute("shadowRootDelegatesFocus", w.shadowRootDelegatesFocus, w.setShadowRootDelegatesFocus)
	jsClass.CreatePrototypeAttribute("shadowRootClonable", w.shadowRootClonable, w.setShadowRootClonable)
	jsClass.CreatePrototypeAttribute("shadowRootSerializable", w.shadowRootSerializable, w.setShadowRootSerializable)
}

func (w htmlTemplateElementV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlTemplateElementV8Wrapper) content(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.content")
	instance, err := js.As[html.HTMLTemplateElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Content()
	return encodeEntity(cbCtx, result)
}

func (w htmlTemplateElementV8Wrapper) shadowRootMode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootMode")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootMode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootMode")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) shadowRootDelegatesFocus(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootDelegatesFocus")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootDelegatesFocus(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootDelegatesFocus")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) shadowRootClonable(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootClonable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootClonable(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootClonable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) shadowRootSerializable(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootSerializable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper) setShadowRootSerializable(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootSerializable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
