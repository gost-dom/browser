// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeNonDocumentTypeChildNode[T any](jsClass js.Class[T]) {
	jsClass.CreateAttribute("previousElementSibling", NonDocumentTypeChildNode_previousElementSibling, nil)
	jsClass.CreateAttribute("nextElementSibling", NonDocumentTypeChildNode_nextElementSibling, nil)
}

func NonDocumentTypeChildNode_previousElementSibling[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousElementSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func NonDocumentTypeChildNode_nextElementSibling[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextElementSibling()
	return codec.EncodeEntity(cbCtx, result)
}
