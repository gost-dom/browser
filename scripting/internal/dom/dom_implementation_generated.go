// This file is generated. Do not edit.

package dom

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeDOMImplementation[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("createDocumentType", DOMImplementation_createDocumentType)
	jsClass.CreateOperation("createDocument", DOMImplementation_createDocument)
	jsClass.CreateOperation("createHTMLDocument", DOMImplementation_createHTMLDocument)
	jsClass.CreateOperation("hasFeature", DOMImplementation_hasFeature)
}

func DOMImplementation_createDocumentType[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*html.DOMImplementation](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	publicId, errArg2 := js.ConsumeArgument(cbCtx, "publicId", nil, codec.DecodeString)
	systemId, errArg3 := js.ConsumeArgument(cbCtx, "systemId", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2, errArg3)
	if err != nil {
		return nil, err
	}
	result := instance.CreateDocumentType(qualifiedName, publicId, systemId)
	return codec.EncodeEntity(cbCtx, result)
}

func DOMImplementation_hasFeature[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*html.DOMImplementation](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.HasFeature()
	return codec.EncodeBoolean(cbCtx, result)
}
