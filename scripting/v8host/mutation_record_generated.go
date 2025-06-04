// This file is generated. Do not edit.

package v8host

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerJSClass("MutationRecord", "", createMutationRecordPrototype)
}

type mutationRecordV8Wrapper struct {
	handleReffedObject[*dominterfaces.MutationRecord, jsTypeParam]
}

func newMutationRecordV8Wrapper(scriptHost *V8ScriptHost) *mutationRecordV8Wrapper {
	return &mutationRecordV8Wrapper{newHandleReffedObject[*dominterfaces.MutationRecord](scriptHost)}
}

func createMutationRecordPrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newMutationRecordV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper mutationRecordV8Wrapper) initialize(jsClass v8Class) {
	wrapper.installPrototype(jsClass)
}

func (w mutationRecordV8Wrapper) installPrototype(jsClass v8Class) {
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

func (w mutationRecordV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w mutationRecordV8Wrapper) type_(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.type_")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type
	return w.toString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper) target(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.target")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper) addedNodes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.addedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AddedNodes
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper) removedNodes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.removedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.RemovedNodes
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper) previousSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.previousSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousSibling
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper) nextSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.nextSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextSibling
	return encodeEntity(cbCtx, result)
}

func (w mutationRecordV8Wrapper) attributeName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.attributeName")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeName
	return w.toNullableString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper) attributeNamespace(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.attributeNamespace")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeNamespace
	return w.toNullableString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper) oldValue(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MutationRecord.oldValue")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OldValue
	return w.toNullableString_(cbCtx, result)
}
