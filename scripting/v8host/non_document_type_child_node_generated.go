// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

type nonDocumentTypeChildNodeV8Wrapper struct {
	handleReffedObject[dom.NonDocumentTypeChildNode]
}

func newNonDocumentTypeChildNodeV8Wrapper(scriptHost *V8ScriptHost) *nonDocumentTypeChildNodeV8Wrapper {
	return &nonDocumentTypeChildNodeV8Wrapper{newHandleReffedObject[dom.NonDocumentTypeChildNode](scriptHost)}
}

func createNonDocumentTypeChildNodePrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newNonDocumentTypeChildNodeV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w nonDocumentTypeChildNodeV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("previousElementSibling",
		v8.NewFunctionTemplateWithError(iso, w.previousElementSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nextElementSibling",
		v8.NewFunctionTemplateWithError(iso, w.nextElementSibling),
		nil,
		v8.None)
}

func (w nonDocumentTypeChildNodeV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w nonDocumentTypeChildNodeV8Wrapper) previousElementSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: NonDocumentTypeChildNode.previousElementSibling")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousElementSibling()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w nonDocumentTypeChildNodeV8Wrapper) nextElementSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: NonDocumentTypeChildNode.nextElementSibling")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextElementSibling()
	return cbCtx.Context().getInstanceForNode(result)
}
