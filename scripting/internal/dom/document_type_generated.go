// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeDocumentType[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("name", DocumentType_name, nil)
	jsClass.CreateAttribute("publicId", DocumentType_publicId, nil)
	jsClass.CreateAttribute("systemId", DocumentType_systemId, nil)
	InitializeChildNode(jsClass)
}

func DocumentType_name[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.DocumentType](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Name()
	return codec.EncodeString(cbCtx, result)
}

func DocumentType_publicId[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.DocumentType](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PublicId()
	return codec.EncodeString(cbCtx, result)
}

func DocumentType_systemId[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.DocumentType](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.SystemId()
	return codec.EncodeString(cbCtx, result)
}
