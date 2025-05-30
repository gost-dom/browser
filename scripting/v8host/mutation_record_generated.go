// This file is generated. Do not edit.

package v8host

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
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

func createMutationRecordPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newMutationRecordV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w mutationRecordV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {

	prototypeTmpl.SetAccessorProperty("type",
		wrapV8Callback(w.scriptHost, w.type_),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		wrapV8Callback(w.scriptHost, w.target),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("addedNodes",
		wrapV8Callback(w.scriptHost, w.addedNodes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("removedNodes",
		wrapV8Callback(w.scriptHost, w.removedNodes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("previousSibling",
		wrapV8Callback(w.scriptHost, w.previousSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nextSibling",
		wrapV8Callback(w.scriptHost, w.nextSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributeName",
		wrapV8Callback(w.scriptHost, w.attributeName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributeNamespace",
		wrapV8Callback(w.scriptHost, w.attributeNamespace),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("oldValue",
		wrapV8Callback(w.scriptHost, w.oldValue),
		nil,
		v8.None)
}

func (w mutationRecordV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w mutationRecordV8Wrapper) type_(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.type_")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Type
	return w.toString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper) target(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.target")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Target
	return w.toJSWrapper(cbCtx, result)
}

func (w mutationRecordV8Wrapper) addedNodes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.addedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AddedNodes
	return w.toJSWrapper(cbCtx, result)
}

func (w mutationRecordV8Wrapper) removedNodes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.removedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.RemovedNodes
	return w.toJSWrapper(cbCtx, result)
}

func (w mutationRecordV8Wrapper) previousSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.previousSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousSibling
	return w.toJSWrapper(cbCtx, result)
}

func (w mutationRecordV8Wrapper) nextSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.nextSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextSibling
	return w.toJSWrapper(cbCtx, result)
}

func (w mutationRecordV8Wrapper) attributeName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.attributeName")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeName
	return w.toNullableString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper) attributeNamespace(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.attributeNamespace")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.AttributeNamespace
	return w.toNullableString_(cbCtx, result)
}

func (w mutationRecordV8Wrapper) oldValue(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.oldValue")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OldValue
	return w.toNullableString_(cbCtx, result)
}
