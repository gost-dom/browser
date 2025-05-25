// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type nonDocumentTypeChildNodeV8Wrapper struct {
	handleReffedObject[dom.NonDocumentTypeChildNode]
}

func newNonDocumentTypeChildNodeV8Wrapper(scriptHost *V8ScriptHost) *nonDocumentTypeChildNodeV8Wrapper {
	return &nonDocumentTypeChildNodeV8Wrapper{newHandleReffedObject[dom.NonDocumentTypeChildNode](scriptHost)}
}

func createNonDocumentTypeChildNodePrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newNonDocumentTypeChildNodeV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w nonDocumentTypeChildNodeV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {

	prototypeTmpl.SetAccessorProperty("previousElementSibling",
		wrapV8Callback(w.scriptHost, w.previousElementSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nextElementSibling",
		wrapV8Callback(w.scriptHost, w.nextElementSibling),
		nil,
		v8.None)
}

func (w nonDocumentTypeChildNodeV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: NonDocumentTypeChildNode.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nonDocumentTypeChildNodeV8Wrapper) previousElementSibling(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: NonDocumentTypeChildNode.previousElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousElementSibling()
	return cbCtx.getInstanceForNode(result)
}

func (w nonDocumentTypeChildNodeV8Wrapper) nextElementSibling(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: NonDocumentTypeChildNode.nextElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextElementSibling()
	return cbCtx.getInstanceForNode(result)
}
