// This file is generated. Do not edit.

package v8host

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
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
	iso := scriptHost.iso
	wrapper := newMutationRecordV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

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

func (w mutationRecordV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w mutationRecordV8Wrapper) type_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.type_")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return w.toString_(cbCtx.Context(), result)
}

func (w mutationRecordV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.target")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return cbCtx.Context().getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) addedNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.addedNodes")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AddedNodes
	return w.toNodeList(cbCtx.Context(), result)
}

func (w mutationRecordV8Wrapper) removedNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.removedNodes")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.RemovedNodes
	return w.toNodeList(cbCtx.Context(), result)
}

func (w mutationRecordV8Wrapper) previousSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.previousSibling")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling
	return cbCtx.Context().getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) nextSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.nextSibling")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling
	return cbCtx.Context().getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) attributeName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.attributeName")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeName
	return w.toNullableString_(cbCtx.Context(), result)
}

func (w mutationRecordV8Wrapper) attributeNamespace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.attributeNamespace")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.AttributeNamespace
	return w.toNullableString_(cbCtx.Context(), result)
}

func (w mutationRecordV8Wrapper) oldValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.oldValue")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[*dominterfaces.MutationRecord](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OldValue
	return w.toNullableString_(cbCtx.Context(), result)
}
