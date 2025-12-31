// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeHTMLTemplateElement[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("content", HTMLTemplateElement_content, nil)
	jsClass.CreateAttribute("shadowRootMode", HTMLTemplateElement_shadowRootMode, HTMLTemplateElement_setShadowRootMode)
	jsClass.CreateAttribute("shadowRootDelegatesFocus", HTMLTemplateElement_shadowRootDelegatesFocus, HTMLTemplateElement_setShadowRootDelegatesFocus)
	jsClass.CreateAttribute("shadowRootClonable", HTMLTemplateElement_shadowRootClonable, HTMLTemplateElement_setShadowRootClonable)
	jsClass.CreateAttribute("shadowRootSerializable", HTMLTemplateElement_shadowRootSerializable, HTMLTemplateElement_setShadowRootSerializable)
}

func HTMLTemplateElement_content[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLTemplateElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Content()
	return codec.EncodeEntity(cbCtx, result)
}

func HTMLTemplateElement_shadowRootMode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_shadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLTemplateElement_setShadowRootMode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_setShadowRootMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLTemplateElement_shadowRootDelegatesFocus[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_shadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLTemplateElement_setShadowRootDelegatesFocus[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_setShadowRootDelegatesFocus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLTemplateElement_shadowRootClonable[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_shadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLTemplateElement_setShadowRootClonable[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_setShadowRootClonable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLTemplateElement_shadowRootSerializable[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_shadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func HTMLTemplateElement_setShadowRootSerializable[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "HTMLTemplateElement.HTMLTemplateElement_setShadowRootSerializable: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
