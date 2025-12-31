// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeChildNode[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("before", ChildNode_before)
	jsClass.CreateOperation("after", ChildNode_after)
	jsClass.CreateOperation("replaceWith", ChildNode_replaceWith)
	jsClass.CreateOperation("remove", ChildNode_remove)
}

func ChildNode_before[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ChildNode.ChildNode_before: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func ChildNode_after[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ChildNode.ChildNode_after: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func ChildNode_replaceWith[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "ChildNode.ChildNode_replaceWith: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func ChildNode_remove[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.ChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Remove()
	return nil, nil
}
