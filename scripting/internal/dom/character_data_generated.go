// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
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
	jsClass.CreatePrototypeMethod("substringData", w.substringData)
	jsClass.CreatePrototypeMethod("appendData", w.appendData)
	jsClass.CreatePrototypeMethod("insertData", w.insertData)
	jsClass.CreatePrototypeMethod("deleteData", w.deleteData)
	jsClass.CreatePrototypeMethod("replaceData", w.replaceData)
	jsClass.CreatePrototypeAttribute("data", w.data, w.setData)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
	w.nonDocumentTypeChildNode.installPrototype(jsClass)
	w.childNode.installPrototype(jsClass)
}

func (w CharacterData[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w CharacterData[T]) substringData(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.substringData")
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.substringData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) appendData(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.appendData")
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.appendData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) insertData(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.insertData")
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.insertData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) deleteData(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.deleteData")
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.deleteData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) replaceData(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.replaceData")
	return codec.EncodeCallbackErrorf(cbCtx, "CharacterData.replaceData: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w CharacterData[T]) data(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.data")
	instance, err := js.As[dom.CharacterData](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Data()
	return codec.EncodeString(cbCtx, result)
}

func (w CharacterData[T]) setData(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.setData")
	instance, err0 := js.As[dom.CharacterData](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetData(val)
	return nil, nil
}

func (w CharacterData[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: CharacterData.length")
	instance, err := js.As[dom.CharacterData](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
