// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("Document", "Node", createDocumentPrototype)
}

type documentV8Wrapper struct {
	handleReffedObject[dom.Document, jsTypeParam]
	parentNode *parentNodeV8Wrapper
}

func newDocumentV8Wrapper(scriptHost *V8ScriptHost) *documentV8Wrapper {
	return &documentV8Wrapper{
		newHandleReffedObject[dom.Document](scriptHost),
		newParentNodeV8Wrapper(scriptHost),
	}
}

func createDocumentPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newDocumentV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w documentV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("getElementsByTagName", wrapV8Callback(w.scriptHost, w.getElementsByTagName))
	prototypeTmpl.Set("getElementsByTagNameNS", wrapV8Callback(w.scriptHost, w.getElementsByTagNameNS))
	prototypeTmpl.Set("getElementsByClassName", wrapV8Callback(w.scriptHost, w.getElementsByClassName))
	prototypeTmpl.Set("createElement", wrapV8Callback(w.scriptHost, w.createElement))
	prototypeTmpl.Set("createElementNS", wrapV8Callback(w.scriptHost, w.createElementNS))
	prototypeTmpl.Set("createDocumentFragment", wrapV8Callback(w.scriptHost, w.createDocumentFragment))
	prototypeTmpl.Set("createTextNode", wrapV8Callback(w.scriptHost, w.createTextNode))
	prototypeTmpl.Set("createCDATASection", wrapV8Callback(w.scriptHost, w.createCDATASection))
	prototypeTmpl.Set("createComment", wrapV8Callback(w.scriptHost, w.createComment))
	prototypeTmpl.Set("createProcessingInstruction", wrapV8Callback(w.scriptHost, w.createProcessingInstruction))
	prototypeTmpl.Set("importNode", wrapV8Callback(w.scriptHost, w.importNode))
	prototypeTmpl.Set("adoptNode", wrapV8Callback(w.scriptHost, w.adoptNode))
	prototypeTmpl.Set("createAttribute", wrapV8Callback(w.scriptHost, w.createAttribute))
	prototypeTmpl.Set("createAttributeNS", wrapV8Callback(w.scriptHost, w.createAttributeNS))
	prototypeTmpl.Set("createEvent", wrapV8Callback(w.scriptHost, w.createEvent))
	prototypeTmpl.Set("createRange", wrapV8Callback(w.scriptHost, w.createRange))
	prototypeTmpl.Set("createNodeIterator", wrapV8Callback(w.scriptHost, w.createNodeIterator))
	prototypeTmpl.Set("createTreeWalker", wrapV8Callback(w.scriptHost, w.createTreeWalker))

	prototypeTmpl.SetAccessorProperty("implementation",
		wrapV8Callback(w.scriptHost, w.implementation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("URL",
		wrapV8Callback(w.scriptHost, w.URL),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("documentURI",
		wrapV8Callback(w.scriptHost, w.documentURI),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("compatMode",
		wrapV8Callback(w.scriptHost, w.compatMode),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("characterSet",
		wrapV8Callback(w.scriptHost, w.characterSet),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("charset",
		wrapV8Callback(w.scriptHost, w.charset),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("inputEncoding",
		wrapV8Callback(w.scriptHost, w.inputEncoding),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("contentType",
		wrapV8Callback(w.scriptHost, w.contentType),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("doctype",
		wrapV8Callback(w.scriptHost, w.doctype),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("documentElement",
		wrapV8Callback(w.scriptHost, w.documentElement),
		nil,
		v8.None)
	w.parentNode.installPrototype(prototypeTmpl)
}

func (w documentV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w documentV8Wrapper) getElementsByTagName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.getElementsByTagName")
	return cbCtx.ReturnWithError(errors.New("Document.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) getElementsByTagNameNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.getElementsByTagNameNS")
	return cbCtx.ReturnWithError(errors.New("Document.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) getElementsByClassName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.getElementsByClassName")
	return cbCtx.ReturnWithError(errors.New("Document.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createElementNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createElementNS")
	return cbCtx.ReturnWithError(errors.New("Document.createElementNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createDocumentFragment(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createDocumentFragment")
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.CreateDocumentFragment()
	return encodeEntity(cbCtx, result)
}

func (w documentV8Wrapper) createCDATASection(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createCDATASection")
	return cbCtx.ReturnWithError(errors.New("Document.createCDATASection: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createComment(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createComment")
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	data, errArg1 := consumeArgument(cbCtx, "data", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateComment(data)
	return encodeEntity(cbCtx, result)
}

func (w documentV8Wrapper) createProcessingInstruction(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createProcessingInstruction")
	return cbCtx.ReturnWithError(errors.New("Document.createProcessingInstruction: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) importNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.importNode")
	return cbCtx.ReturnWithError(errors.New("Document.importNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) adoptNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.adoptNode")
	return cbCtx.ReturnWithError(errors.New("Document.adoptNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createAttribute")
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	localName, errArg1 := consumeArgument(cbCtx, "localName", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateAttribute(localName)
	return encodeEntity(cbCtx, result)
}

func (w documentV8Wrapper) createAttributeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Document.createAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createEvent(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createEvent")
	return cbCtx.ReturnWithError(errors.New("Document.createEvent: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createRange(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createRange")
	return cbCtx.ReturnWithError(errors.New("Document.createRange: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createNodeIterator(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createNodeIterator")
	return cbCtx.ReturnWithError(errors.New("Document.createNodeIterator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) createTreeWalker(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createTreeWalker")
	return cbCtx.ReturnWithError(errors.New("Document.createTreeWalker: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) implementation(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.implementation")
	return cbCtx.ReturnWithError(errors.New("Document.implementation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) URL(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.URL")
	return cbCtx.ReturnWithError(errors.New("Document.URL: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) documentURI(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.documentURI")
	return cbCtx.ReturnWithError(errors.New("Document.documentURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) compatMode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.compatMode")
	return cbCtx.ReturnWithError(errors.New("Document.compatMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) characterSet(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.characterSet")
	return cbCtx.ReturnWithError(errors.New("Document.characterSet: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) charset(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.charset")
	return cbCtx.ReturnWithError(errors.New("Document.charset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) inputEncoding(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.inputEncoding")
	return cbCtx.ReturnWithError(errors.New("Document.inputEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) contentType(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.contentType")
	return cbCtx.ReturnWithError(errors.New("Document.contentType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) doctype(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.doctype")
	return cbCtx.ReturnWithError(errors.New("Document.doctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper) documentElement(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Document.documentElement")
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.DocumentElement()
	return encodeEntity(cbCtx, result)
}
