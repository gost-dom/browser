// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeNonElementParentNode[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("getElementById", NonElementParentNode_getElementById)
}

func NonElementParentNode_getElementById[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.NonElementParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	elementId, errArg1 := js.ConsumeArgument(cbCtx, "elementId", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetElementById(elementId)
	return codec.EncodeEntity(cbCtx, result)
}
