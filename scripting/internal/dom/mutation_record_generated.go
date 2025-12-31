// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MutationRecord[T any] struct{}

func NewMutationRecord[T any](scriptHost js.ScriptEngine[T]) MutationRecord[T] {
	return MutationRecord[T]{}
}

func (wrapper MutationRecord[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MutationRecord[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateAttribute("type", MutationRecord_type, nil)
	jsClass.CreateAttribute("target", MutationRecord_target, nil)
	jsClass.CreateAttribute("addedNodes", MutationRecord_addedNodes, nil)
	jsClass.CreateAttribute("removedNodes", MutationRecord_removedNodes, nil)
	jsClass.CreateAttribute("previousSibling", MutationRecord_previousSibling, nil)
	jsClass.CreateAttribute("nextSibling", MutationRecord_nextSibling, nil)
	jsClass.CreateAttribute("attributeName", MutationRecord_attributeName, nil)
	jsClass.CreateAttribute("attributeNamespace", MutationRecord_attributeNamespace, nil)
	jsClass.CreateAttribute("oldValue", MutationRecord_oldValue, nil)
}

func MutationRecordConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func MutationRecord_type[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func MutationRecord_target[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return codec.EncodeEntity(cbCtx, result)
}

func MutationRecord_addedNodes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AddedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func MutationRecord_removedNodes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.RemovedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func MutationRecord_previousSibling[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling
	return codec.EncodeEntity(cbCtx, result)
}

func MutationRecord_nextSibling[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling
	return codec.EncodeEntity(cbCtx, result)
}

func MutationRecord_attributeName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeName
	return codec.EncodeNullableString(cbCtx, result)
}

func MutationRecord_attributeNamespace[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeNamespace
	return codec.EncodeNullableString(cbCtx, result)
}

func MutationRecord_oldValue[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OldValue
	return codec.EncodeNullableString(cbCtx, result)
}
