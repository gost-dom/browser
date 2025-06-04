// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type nonDocumentTypeChildNodeV8Wrapper struct {
	handleReffedObject[dom.NonDocumentTypeChildNode, jsTypeParam]
}

func newNonDocumentTypeChildNodeV8Wrapper(scriptHost *V8ScriptHost) *nonDocumentTypeChildNodeV8Wrapper {
	return &nonDocumentTypeChildNodeV8Wrapper{newHandleReffedObject[dom.NonDocumentTypeChildNode](scriptHost)}
}

func createNonDocumentTypeChildNodePrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newNonDocumentTypeChildNodeV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper nonDocumentTypeChildNodeV8Wrapper) initialize(jsClass v8Class) {
	wrapper.installPrototype(jsClass)
}

func (w nonDocumentTypeChildNodeV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeAttribute("previousElementSibling", w.previousElementSibling, nil)
	jsClass.CreatePrototypeAttribute("nextElementSibling", w.nextElementSibling, nil)
}

func (w nonDocumentTypeChildNodeV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nonDocumentTypeChildNodeV8Wrapper) previousElementSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.previousElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousElementSibling()
	return encodeEntity(cbCtx, result)
}

func (w nonDocumentTypeChildNodeV8Wrapper) nextElementSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: NonDocumentTypeChildNode.nextElementSibling")
	instance, err := js.As[dom.NonDocumentTypeChildNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextElementSibling()
	return encodeEntity(cbCtx, result)
}
