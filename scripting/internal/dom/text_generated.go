// This file is generated. Do not edit.

package dom

import (
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeText[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("splitText", Text_splitText)
	jsClass.CreateAttribute("wholeText", Text_wholeText, nil)
}

func TextConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func Text_splitText[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Text.Text_splitText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Text_wholeText[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Text.Text_wholeText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
