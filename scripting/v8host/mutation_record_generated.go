// This file is generated. Do not edit.

package v8host

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type mutationRecordV8Wrapper[T any] struct {
	handleReffedObject[*dominterfaces.MutationRecord, T]
}

func newMutationRecordV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *mutationRecordV8Wrapper[T] {
	return &mutationRecordV8Wrapper[T]{newHandleReffedObject[*dominterfaces.MutationRecord, T](scriptHost)}
}

func (wrapper mutationRecordV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w mutationRecordV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w mutationRecordV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w mutationRecordV8Wrapper[T]) type_(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.type_")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type
	return w.toString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) target(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.target")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) addedNodes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.addedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AddedNodes
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) removedNodes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.removedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.RemovedNodes
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) previousSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.previousSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousSibling
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) nextSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.nextSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextSibling
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) attributeName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.attributeName")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeName
	return w.toNullableString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) attributeNamespace(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.attributeNamespace")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeNamespace
	return w.toNullableString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper[T]) oldValue(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.oldValue")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OldValue
	return w.toNullableString_(cbCtx, result)
}
