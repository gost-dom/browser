// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type htmlTemplateElementV8Wrapper[T any] struct{}

func newHTMLTemplateElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *htmlTemplateElementV8Wrapper[T] {
	return &htmlTemplateElementV8Wrapper[T]{}
}

func (wrapper htmlTemplateElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w htmlTemplateElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("content", w.content, nil)
	jsClass.CreatePrototypeAttribute("shadowRootMode", w.shadowRootMode, w.setShadowRootMode)
	jsClass.CreatePrototypeAttribute("shadowRootDelegatesFocus", w.shadowRootDelegatesFocus, w.setShadowRootDelegatesFocus)
	jsClass.CreatePrototypeAttribute("shadowRootClonable", w.shadowRootClonable, w.setShadowRootClonable)
	jsClass.CreatePrototypeAttribute("shadowRootSerializable", w.shadowRootSerializable, w.setShadowRootSerializable)
}

func (w htmlTemplateElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlTemplateElementV8Wrapper[T]) content(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.content")
	instance, err := js.As[html.HTMLTemplateElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Content()
	return encodeEntity(cbCtx, result)
}

func (w htmlTemplateElementV8Wrapper[T]) shadowRootMode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootMode")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper[T]) setShadowRootMode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootMode")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper[T]) shadowRootDelegatesFocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootDelegatesFocus")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper[T]) setShadowRootDelegatesFocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootDelegatesFocus")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper[T]) shadowRootClonable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootClonable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper[T]) setShadowRootClonable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootClonable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper[T]) shadowRootSerializable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootSerializable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.shadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w htmlTemplateElementV8Wrapper[T]) setShadowRootSerializable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootSerializable")
	return cbCtx.ReturnWithError(errors.New("HTMLTemplateElement.setShadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
