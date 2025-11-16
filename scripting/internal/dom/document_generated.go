// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Document[T any] struct {
	parentNode *ParentNode[T]
}

func NewDocument[T any](scriptHost js.ScriptEngine[T]) *Document[T] {
	return &Document[T]{NewParentNode(scriptHost)}
}

func (wrapper Document[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w Document[T]) installPrototype(jsClass js.Class[T]) {
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

func (w Document[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.Constructor", js.LogAttr("res", res))
	}()
	return w.CreateInstance(cbCtx)
}

func (w Document[T]) getElementsByTagName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.getElementsByTagName - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.getElementsByTagName", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetElementsByTagName(qualifiedName)
	return w.toHTMLCollection(cbCtx, result)
}

func (w Document[T]) getElementsByTagNameNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.getElementsByTagNameNS - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.getElementsByTagNameNS", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) getElementsByClassName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.getElementsByClassName - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.getElementsByClassName", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) createElementNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createElementNS - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createElementNS", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	namespace, errArg1 := js.ConsumeArgument(cbCtx, "namespace", codec.ZeroValue, codec.DecodeString)
	qualifiedName, errArg2 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.CreateElementNS(namespace, qualifiedName)
	return codec.EncodeEntity(cbCtx, result)
}

func (w Document[T]) createDocumentFragment(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createDocumentFragment - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createDocumentFragment", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CreateDocumentFragment()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Document[T]) createTextNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createTextNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createTextNode", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateTextNode(data)
	return codec.EncodeEntity(cbCtx, result)
}

func (w Document[T]) createCDATASection(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createCDATASection - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createCDATASection", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.createCDATASection: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) createComment(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createComment - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createComment", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	data, errArg1 := js.ConsumeArgument(cbCtx, "data", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateComment(data)
	return codec.EncodeEntity(cbCtx, result)
}

func (w Document[T]) createProcessingInstruction(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createProcessingInstruction - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createProcessingInstruction", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.createProcessingInstruction: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) importNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.importNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.importNode", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	subtree, errArg2 := js.ConsumeArgument(cbCtx, "subtree", codec.ZeroValue, codec.DecodeBoolean)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.ImportNode(node, subtree)
	return codec.EncodeEntity(cbCtx, result)
}

func (w Document[T]) adoptNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.adoptNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.adoptNode", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.adoptNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) createAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createAttribute - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createAttribute", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Document](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	localName, errArg1 := js.ConsumeArgument(cbCtx, "localName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateAttribute(localName)
	return codec.EncodeEntity(cbCtx, result)
}

func (w Document[T]) createAttributeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createAttributeNS - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createAttributeNS", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.createAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) createEvent(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createEvent - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createEvent", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.createEvent: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) createRange(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createRange - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createRange", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.createRange: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) createNodeIterator(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createNodeIterator - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createNodeIterator", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.createNodeIterator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) createTreeWalker(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.createTreeWalker - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.createTreeWalker", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.createTreeWalker: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) implementation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.implementation - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.implementation", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.implementation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) URL(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.URL - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.URL", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.URL: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) documentURI(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.documentURI - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.documentURI", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.documentURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) compatMode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.compatMode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.compatMode", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.compatMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) characterSet(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.characterSet - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.characterSet", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.characterSet: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) charset(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.charset - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.charset", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.charset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) inputEncoding(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.inputEncoding - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.inputEncoding", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.inputEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) contentType(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.contentType - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.contentType", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.contentType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) doctype(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.doctype - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.doctype", js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Document.doctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Document[T]) documentElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Document.documentElement - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Document.documentElement", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Document](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DocumentElement()
	return codec.EncodeEntity(cbCtx, result)
}
