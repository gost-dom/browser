// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeCharacterData[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("substringData", CharacterData_substringData)
	jsClass.CreateOperation("appendData", CharacterData_appendData)
	jsClass.CreateOperation("insertData", CharacterData_insertData)
	jsClass.CreateOperation("deleteData", CharacterData_deleteData)
	jsClass.CreateOperation("replaceData", CharacterData_replaceData)
	jsClass.CreateAttribute("data", CharacterData_data, CharacterData_setData)
	jsClass.CreateAttribute("length", CharacterData_length, nil)
	InitializeNonDocumentTypeChildNode(jsClass)
	InitializeChildNode(jsClass)
}

func CharacterDataConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func CharacterData_substringData[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.CharacterData_substringData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func CharacterData_appendData[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.CharacterData_appendData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func CharacterData_insertData[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.CharacterData_insertData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func CharacterData_deleteData[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.CharacterData_deleteData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func CharacterData_replaceData[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.CharacterData_replaceData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func CharacterData_data[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.CharacterData](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Data()
	return codec.EncodeString(cbCtx, result)
}

func CharacterData_setData[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.CharacterData](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetData(val)
	return nil, nil
}

func CharacterData_length[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.CharacterData](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
