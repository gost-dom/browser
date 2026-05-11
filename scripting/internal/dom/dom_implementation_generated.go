// This file is generated. Do not edit.

package dom

import (
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeDOMImplementation[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("createDocument", DOMImplementation_createDocument)
	jsClass.CreateOperation("createHTMLDocument", DOMImplementation_createHTMLDocument)
	jsClass.CreateOperation("hasFeature", DOMImplementation_hasFeature)
}

func DOMImplementation_hasFeature[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*html.DOMImplementation](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.HasFeature()
	return codec.EncodeBoolean(cbCtx, result)
}
