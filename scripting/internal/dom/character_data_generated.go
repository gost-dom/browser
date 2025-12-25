// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type CharacterData[T any] struct {
	nonDocumentTypeChildNode *NonDocumentTypeChildNode[T]
	childNode                *ChildNode[T]
}

func NewCharacterData[T any](scriptHost js.ScriptEngine[T]) *CharacterData[T] {
	return &CharacterData[T]{
		NewNonDocumentTypeChildNode(scriptHost),
		NewChildNode(scriptHost),
	}
}

func (wrapper CharacterData[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w CharacterData[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("substringData", w.substringData)
	jsClass.CreateOperation("appendData", w.appendData)
	jsClass.CreateOperation("insertData", w.insertData)
	jsClass.CreateOperation("deleteData", w.deleteData)
	jsClass.CreateOperation("replaceData", w.replaceData)
	jsClass.CreateAttribute("data", w.data, w.setData)
	jsClass.CreateAttribute("length", w.length, nil)
	w.nonDocumentTypeChildNode.installPrototype(jsClass)
	w.childNode.installPrototype(jsClass)
}

func (w CharacterData[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w CharacterData[T]) substringData(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.substringData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) appendData(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.appendData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) insertData(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.insertData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) deleteData(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.deleteData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) replaceData(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.replaceData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) data(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.CharacterData](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Data()
	return codec.EncodeString(cbCtx, result)
}

func (w CharacterData[T]) setData(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.CharacterData](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetData(val)
	return nil, nil
}

func (w CharacterData[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.CharacterData](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
