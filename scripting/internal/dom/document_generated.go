// This file is generated. Do not edit.

package dom

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Document[T any] struct {
	nonElementParentNode *NonElementParentNode[T]
	parentNode           *ParentNode[T]
}

func NewDocument[T any](scriptHost js.ScriptEngine[T]) *Document[T] {
	return &Document[T]{
		NewNonElementParentNode(scriptHost),
		NewParentNode(scriptHost),
	}
}

func (wrapper Document[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Document[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("getElementsByTagName", Document_getElementsByTagName)
	jsClass.CreateOperation("getElementsByTagNameNS", Document_getElementsByTagNameNS)
	jsClass.CreateOperation("getElementsByClassName", Document_getElementsByClassName)
	jsClass.CreateOperation("createElement", Document_createElement)
	jsClass.CreateOperation("createElementNS", Document_createElementNS)
	jsClass.CreateOperation("createDocumentFragment", Document_createDocumentFragment)
	jsClass.CreateOperation("createTextNode", Document_createTextNode)
	jsClass.CreateOperation("createCDATASection", Document_createCDATASection)
	jsClass.CreateOperation("createComment", Document_createComment)
	jsClass.CreateOperation("createProcessingInstruction", Document_createProcessingInstruction)
	jsClass.CreateOperation("importNode", Document_importNode)
	jsClass.CreateOperation("adoptNode", Document_adoptNode)
	jsClass.CreateOperation("createAttribute", Document_createAttribute)
	jsClass.CreateOperation("createAttributeNS", Document_createAttributeNS)
	jsClass.CreateOperation("createEvent", Document_createEvent)
	jsClass.CreateOperation("createRange", Document_createRange)
	jsClass.CreateOperation("createNodeIterator", Document_createNodeIterator)
	jsClass.CreateOperation("createTreeWalker", Document_createTreeWalker)
	jsClass.CreateAttribute("implementation", Document_implementation, nil)
	jsClass.CreateAttribute("URL", Document_URL, nil)
	jsClass.CreateAttribute("documentURI", Document_documentURI, nil)
	jsClass.CreateAttribute("compatMode", Document_compatMode, nil)
	jsClass.CreateAttribute("characterSet", Document_characterSet, nil)
	jsClass.CreateAttribute("charset", Document_charset, nil)
	jsClass.CreateAttribute("inputEncoding", Document_inputEncoding, nil)
	jsClass.CreateAttribute("contentType", Document_contentType, nil)
	jsClass.CreateAttribute("doctype", Document_doctype, nil)
	jsClass.CreateAttribute("documentElement", Document_documentElement, nil)
	jsClass.CreateAttribute("location", Document_location, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("body", Document_body, Document_setBody)
	jsClass.CreateAttribute("head", Document_head, nil)
	w.nonElementParentNode.installPrototype(jsClass)
	w.parentNode.installPrototype(jsClass)
}

func DocumentConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return CreateDocument(cbCtx)
}

func Document_getElementsByTagName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetElementsByTagName(qualifiedName)
	return encodeHTMLCollection(cbCtx, result)
}

func Document_getElementsByTagNameNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	namespace, errArg1 := js.ConsumeArgument(cbCtx, "namespace", codec.ZeroValue, codec.DecodeString)
	localName, errArg2 := js.ConsumeArgument(cbCtx, "localName", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.GetElementsByTagNameNS(namespace, localName)
	return encodeHTMLCollection(cbCtx, result)
}

func Document_getElementsByClassName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_createElement[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	localName, errArg1 := js.ConsumeArgument(cbCtx, "localName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CreateElement(localName)
	return codec.EncodeEntity(cbCtx, result)
}

func Document_createElementNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	namespace, errArg1 := js.ConsumeArgument(cbCtx, "namespace", codec.ZeroValue, codec.DecodeString)
	qualifiedName, errArg2 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.CreateElementNS(namespace, qualifiedName)
	return codec.EncodeEntity(cbCtx, result)
}

func Document_createDocumentFragment[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CreateDocumentFragment()
	return codec.EncodeEntity(cbCtx, result)
}

func Document_createTextNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
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

func Document_createCDATASection[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_createCDATASection: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_createComment[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
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

func Document_createProcessingInstruction[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	target, errArg1 := js.ConsumeArgument(cbCtx, "target", nil, codec.DecodeString)
	data, errArg2 := js.ConsumeArgument(cbCtx, "data", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.CreateProcessingInstruction(target, data)
	return codec.EncodeEntity(cbCtx, result)
}

func Document_importNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	subtree, errArg2 := js.ConsumeArgument(cbCtx, "subtree", codec.ZeroValue, codec.DecodeBoolean)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.ImportNode(node, subtree)
	return codec.EncodeEntity(cbCtx, result)
}

func Document_adoptNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_adoptNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_createAttribute[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
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

func Document_createAttributeNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[html.HTMLDocument](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	namespace, errArg1 := js.ConsumeArgument(cbCtx, "namespace", codec.ZeroValue, codec.DecodeString)
	qualifiedName, errArg2 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result := instance.CreateAttributeNS(namespace, qualifiedName)
	return codec.EncodeEntity(cbCtx, result)
}

func Document_createEvent[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_createEvent: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_createRange[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_createRange: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_createNodeIterator[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_createNodeIterator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_createTreeWalker[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_createTreeWalker: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_implementation[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_implementation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_URL[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_URL: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_documentURI[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_documentURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_compatMode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_compatMode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_characterSet[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_characterSet: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_charset[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_charset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_inputEncoding[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_inputEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_contentType[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_contentType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_doctype[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Document.Document_doctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Document_documentElement[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DocumentElement()
	return codec.EncodeEntity(cbCtx, result)
}

func Document_location[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Location()
	return codec.EncodeEntity(cbCtx, result)
}

func Document_body[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Body()
	return codec.EncodeEntity(cbCtx, result)
}

func Document_setBody[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[html.HTMLDocument](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeHTMLElement)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetBody(val)
}

func Document_head[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.HTMLDocument](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Head()
	return codec.EncodeEntity(cbCtx, result)
}
