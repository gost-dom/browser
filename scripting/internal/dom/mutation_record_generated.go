// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MutationRecord[T any] struct{}

func NewMutationRecord[T any](scriptHost js.ScriptEngine[T]) *MutationRecord[T] {
	return &MutationRecord[T]{}
}

func (wrapper MutationRecord[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MutationRecord[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateAttribute("type", w.type_, nil)
	jsClass.CreateAttribute("target", w.target, nil)
	jsClass.CreateAttribute("addedNodes", w.addedNodes, nil)
	jsClass.CreateAttribute("removedNodes", w.removedNodes, nil)
	jsClass.CreateAttribute("previousSibling", w.previousSibling, nil)
	jsClass.CreateAttribute("nextSibling", w.nextSibling, nil)
	jsClass.CreateAttribute("attributeName", w.attributeName, nil)
	jsClass.CreateAttribute("attributeNamespace", w.attributeNamespace, nil)
	jsClass.CreateAttribute("oldValue", w.oldValue, nil)
}

func (w MutationRecord[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w MutationRecord[T]) type_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func (w MutationRecord[T]) target(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) addedNodes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AddedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) removedNodes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.RemovedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) previousSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) nextSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecord[T]) attributeName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeName
	return codec.EncodeNullableString(cbCtx, result)
}

func (w MutationRecord[T]) attributeNamespace(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeNamespace
	return codec.EncodeNullableString(cbCtx, result)
}

func (w MutationRecord[T]) oldValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OldValue
	return codec.EncodeNullableString(cbCtx, result)
}
