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
	handleReffedObject[*dominterfaces.MutationRecord]
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

func (w mutationRecordV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w mutationRecordV8Wrapper) type_(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.type_")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return w.toString_(cbCtx.ScriptCtx(), result)
}

func (w mutationRecordV8Wrapper) target(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.target")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return cbCtx.ScriptCtx().getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) addedNodes(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.addedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AddedNodes
	return w.toNodeList(cbCtx.ScriptCtx(), result)
}

func (w mutationRecordV8Wrapper) removedNodes(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.removedNodes")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.RemovedNodes
	return w.toNodeList(cbCtx.ScriptCtx(), result)
}

func (w mutationRecordV8Wrapper) previousSibling(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.previousSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling
	return cbCtx.ScriptCtx().getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) nextSibling(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.nextSibling")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling
	return cbCtx.ScriptCtx().getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) attributeName(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.attributeName")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeName
	return w.toNullableString_(cbCtx.ScriptCtx(), result)
}

func (w mutationRecordV8Wrapper) attributeNamespace(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.attributeNamespace")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeNamespace
	return w.toNullableString_(cbCtx.ScriptCtx(), result)
}

func (w mutationRecordV8Wrapper) oldValue(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MutationRecord.oldValue")
	instance, err := js.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OldValue
	return w.toNullableString_(cbCtx.ScriptCtx(), result)
}
