// This file is generated. Do not edit.

package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MutationRecordV8Wrapper[T any] struct{}

func NewMutationRecordV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *MutationRecordV8Wrapper[T] {
	return &MutationRecordV8Wrapper[T]{}
}

func (wrapper MutationRecordV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MutationRecordV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeAttribute("type", w.type_, nil)
	jsClass.CreatePrototypeAttribute("target", w.target, nil)
	jsClass.CreatePrototypeAttribute("addedNodes", w.addedNodes, nil)
	jsClass.CreatePrototypeAttribute("removedNodes", w.removedNodes, nil)
	jsClass.CreatePrototypeAttribute("previousSibling", w.previousSibling, nil)
	jsClass.CreatePrototypeAttribute("nextSibling", w.nextSibling, nil)
	jsClass.CreatePrototypeAttribute("attributeName", w.attributeName, nil)
	jsClass.CreatePrototypeAttribute("attributeNamespace", w.attributeNamespace, nil)
	jsClass.CreatePrototypeAttribute("oldValue", w.oldValue, nil)
}

func (w MutationRecordV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w MutationRecordV8Wrapper[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.type_")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) target(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.target")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) addedNodes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.addedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AddedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) removedNodes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.removedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.RemovedNodes
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) previousSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.previousSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousSibling
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) nextSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.nextSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextSibling
	return codec.EncodeEntity(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) attributeName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.attributeName")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeName
	return codec.EncodeNullableString(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) attributeNamespace(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.attributeNamespace")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeNamespace
	return codec.EncodeNullableString(cbCtx, result)
}

func (w MutationRecordV8Wrapper[T]) oldValue(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.oldValue")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OldValue
	return codec.EncodeNullableString(cbCtx, result)
}
