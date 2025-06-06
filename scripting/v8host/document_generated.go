// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type documentV8Wrapper[T any] struct {
	parentNode *parentNodeV8Wrapper[T]
}

func newDocumentV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *documentV8Wrapper[T] {
	return &documentV8Wrapper[T]{newParentNodeV8Wrapper(scriptHost)}
}

func (wrapper documentV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w documentV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("getElementsByTagName", w.getElementsByTagName)
	jsClass.CreatePrototypeMethod("getElementsByTagNameNS", w.getElementsByTagNameNS)
	jsClass.CreatePrototypeMethod("getElementsByClassName", w.getElementsByClassName)
	jsClass.CreatePrototypeMethod("createElement", w.createElement)
	jsClass.CreatePrototypeMethod("createElementNS", w.createElementNS)
	jsClass.CreatePrototypeMethod("createDocumentFragment", w.createDocumentFragment)
	jsClass.CreatePrototypeMethod("createTextNode", w.createTextNode)
	jsClass.CreatePrototypeMethod("createCDATASection", w.createCDATASection)
	jsClass.CreatePrototypeMethod("createComment", w.createComment)
	jsClass.CreatePrototypeMethod("createProcessingInstruction", w.createProcessingInstruction)
	jsClass.CreatePrototypeMethod("importNode", w.importNode)
	jsClass.CreatePrototypeMethod("adoptNode", w.adoptNode)
	jsClass.CreatePrototypeMethod("createAttribute", w.createAttribute)
	jsClass.CreatePrototypeMethod("createAttributeNS", w.createAttributeNS)
	jsClass.CreatePrototypeMethod("createEvent", w.createEvent)
	jsClass.CreatePrototypeMethod("createRange", w.createRange)
	jsClass.CreatePrototypeMethod("createNodeIterator", w.createNodeIterator)
	jsClass.CreatePrototypeMethod("createTreeWalker", w.createTreeWalker)
	jsClass.CreatePrototypeAttribute("implementation", w.implementation, nil)
	jsClass.CreatePrototypeAttribute("URL", w.URL, nil)
	jsClass.CreatePrototypeAttribute("documentURI", w.documentURI, nil)
	jsClass.CreatePrototypeAttribute("compatMode", w.compatMode, nil)
	jsClass.CreatePrototypeAttribute("characterSet", w.characterSet, nil)
	jsClass.CreatePrototypeAttribute("charset", w.charset, nil)
	jsClass.CreatePrototypeAttribute("inputEncoding", w.inputEncoding, nil)
	jsClass.CreatePrototypeAttribute("contentType", w.contentType, nil)
	jsClass.CreatePrototypeAttribute("doctype", w.doctype, nil)
	jsClass.CreatePrototypeAttribute("documentElement", w.documentElement, nil)
	w.parentNode.installPrototype(jsClass)
}

func (w documentV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.Constructor")
	return w.CreateInstance(cbCtx)
}

func (w documentV8Wrapper[T]) getElementsByTagName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.getElementsByTagName")
	return cbCtx.ReturnWithError(errors.New("Document.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) getElementsByTagNameNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.getElementsByTagNameNS")
	return cbCtx.ReturnWithError(errors.New("Document.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) getElementsByClassName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.getElementsByClassName")
	return cbCtx.ReturnWithError(errors.New("Document.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createElementNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createElementNS")
	return cbCtx.ReturnWithError(errors.New("Document.createElementNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createDocumentFragment(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createDocumentFragment")
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.CreateDocumentFragment()
	return encodeEntity(cbCtx, result)
}

func (w documentV8Wrapper[T]) createCDATASection(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createCDATASection")
	return cbCtx.ReturnWithError(errors.New("Document.createCDATASection: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createComment(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createComment")
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	data, errArg1 := consumeArgument(cbCtx, "data", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateComment(data)
	return encodeEntity(cbCtx, result)
}

func (w documentV8Wrapper[T]) createProcessingInstruction(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createProcessingInstruction")
	return cbCtx.ReturnWithError(errors.New("Document.createProcessingInstruction: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) importNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.importNode")
	return cbCtx.ReturnWithError(errors.New("Document.importNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) adoptNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.adoptNode")
	return cbCtx.ReturnWithError(errors.New("Document.adoptNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createAttribute(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createAttribute")
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	localName, errArg1 := consumeArgument(cbCtx, "localName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateAttribute(localName)
	return encodeEntity(cbCtx, result)
}

func (w documentV8Wrapper[T]) createAttributeNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Document.createAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createEvent(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createEvent")
	return cbCtx.ReturnWithError(errors.New("Document.createEvent: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createRange(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createRange")
	return cbCtx.ReturnWithError(errors.New("Document.createRange: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createNodeIterator(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createNodeIterator")
	return cbCtx.ReturnWithError(errors.New("Document.createNodeIterator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) createTreeWalker(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.createTreeWalker")
	return cbCtx.ReturnWithError(errors.New("Document.createTreeWalker: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) implementation(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.implementation")
	return cbCtx.ReturnWithError(errors.New("Document.implementation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) URL(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.URL")
	return cbCtx.ReturnWithError(errors.New("Document.URL: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) documentURI(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.documentURI")
	return cbCtx.ReturnWithError(errors.New("Document.documentURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) compatMode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.compatMode")
	return cbCtx.ReturnWithError(errors.New("Document.compatMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) characterSet(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.characterSet")
	return cbCtx.ReturnWithError(errors.New("Document.characterSet: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) charset(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.charset")
	return cbCtx.ReturnWithError(errors.New("Document.charset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) inputEncoding(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.inputEncoding")
	return cbCtx.ReturnWithError(errors.New("Document.inputEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) contentType(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.contentType")
	return cbCtx.ReturnWithError(errors.New("Document.contentType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) doctype(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.doctype")
	return cbCtx.ReturnWithError(errors.New("Document.doctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w documentV8Wrapper[T]) documentElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Document.documentElement")
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.DocumentElement()
	return encodeEntity(cbCtx, result)
}
