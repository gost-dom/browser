// This file is generated. Do not edit.

package v8host

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("MutationRecord", "", createMutationRecordPrototype)
}

type mutationRecordV8Wrapper struct {
	handleReffedObject[dominterfaces.MutationRecord]
}

func newMutationRecordV8Wrapper(scriptHost *V8ScriptHost) *mutationRecordV8Wrapper {
	return &mutationRecordV8Wrapper{newHandleReffedObject[dominterfaces.MutationRecord](scriptHost)}
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
	iso := w.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("type",
		v8.NewFunctionTemplateWithError(iso, w.type_),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, w.target),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("addedNodes",
		v8.NewFunctionTemplateWithError(iso, w.addedNodes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("removedNodes",
		v8.NewFunctionTemplateWithError(iso, w.removedNodes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("previousSibling",
		v8.NewFunctionTemplateWithError(iso, w.previousSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nextSibling",
		v8.NewFunctionTemplateWithError(iso, w.nextSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributeName",
		v8.NewFunctionTemplateWithError(iso, w.attributeName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributeNamespace",
		v8.NewFunctionTemplateWithError(iso, w.attributeNamespace),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("oldValue",
		v8.NewFunctionTemplateWithError(iso, w.oldValue),
		nil,
		v8.None)
}

func (w mutationRecordV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w mutationRecordV8Wrapper) type_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.type")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return w.toString_(ctx, result)
}

func (w mutationRecordV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.target")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return ctx.getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) addedNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.addedNodes")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.AddedNodes
	return w.toNodeList(ctx, result)
}

func (w mutationRecordV8Wrapper) removedNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.removedNodes")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.RemovedNodes
	return w.toNodeList(ctx, result)
}

func (w mutationRecordV8Wrapper) previousSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.previousSibling")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling
	return ctx.getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) nextSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.nextSibling")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling
	return ctx.getInstanceForNode(result)
}

func (w mutationRecordV8Wrapper) attributeName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.attributeName")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.AttributeName
	return w.toNullableString_(ctx, result)
}

func (w mutationRecordV8Wrapper) attributeNamespace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.attributeNamespace")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.AttributeNamespace
	return w.toNullableString_(ctx, result)
}

func (w mutationRecordV8Wrapper) oldValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MutationRecord.oldValue")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.OldValue
	return w.toNullableString_(ctx, result)
}
