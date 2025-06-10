// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type HTMLTemplateElement[T any] struct{}

func NewHTMLTemplateElement[T any](scriptHost js.ScriptEngine[T]) *HTMLTemplateElement[T] {
	return &HTMLTemplateElement[T]{}
}

func (wrapper HTMLTemplateElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w HTMLTemplateElement[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("content", w.content, nil)
	jsClass.CreatePrototypeAttribute("shadowRootMode", w.shadowRootMode, w.setShadowRootMode)
	jsClass.CreatePrototypeAttribute("shadowRootDelegatesFocus", w.shadowRootDelegatesFocus, w.setShadowRootDelegatesFocus)
	jsClass.CreatePrototypeAttribute("shadowRootClonable", w.shadowRootClonable, w.setShadowRootClonable)
	jsClass.CreatePrototypeAttribute("shadowRootSerializable", w.shadowRootSerializable, w.setShadowRootSerializable)
}

func (w HTMLTemplateElement[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w HTMLTemplateElement[T]) content(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.content")
	instance, err := js.As[html.HTMLTemplateElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return codec.EncodeEntity(cbCtx, result)
}

func (w HTMLTemplateElement[T]) shadowRootMode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootMode")
	return nil, errors.New("HTMLTemplateElement.shadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLTemplateElement[T]) setShadowRootMode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootMode")
	return nil, errors.New("HTMLTemplateElement.setShadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLTemplateElement[T]) shadowRootDelegatesFocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootDelegatesFocus")
	return nil, errors.New("HTMLTemplateElement.shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLTemplateElement[T]) setShadowRootDelegatesFocus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootDelegatesFocus")
	return nil, errors.New("HTMLTemplateElement.setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLTemplateElement[T]) shadowRootClonable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootClonable")
	return nil, errors.New("HTMLTemplateElement.shadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLTemplateElement[T]) setShadowRootClonable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootClonable")
	return nil, errors.New("HTMLTemplateElement.setShadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLTemplateElement[T]) shadowRootSerializable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.shadowRootSerializable")
	return nil, errors.New("HTMLTemplateElement.shadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w HTMLTemplateElement[T]) setShadowRootSerializable(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: HTMLTemplateElement.setShadowRootSerializable")
	return nil, errors.New("HTMLTemplateElement.setShadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
