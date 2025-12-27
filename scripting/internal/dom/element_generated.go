// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Element[T any] struct {
	parentNode               *ParentNode[T]
	nonDocumentTypeChildNode *NonDocumentTypeChildNode[T]
	childNode                *ChildNode[T]
}

func NewElement[T any](scriptHost js.ScriptEngine[T]) *Element[T] {
	return &Element[T]{
		NewParentNode(scriptHost),
		NewNonDocumentTypeChildNode(scriptHost),
		NewChildNode(scriptHost),
	}
}

func (wrapper Element[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w Element[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("hasAttributes", w.hasAttributes)
	jsClass.CreateOperation("getAttributeNames", w.getAttributeNames)
	jsClass.CreateOperation("getAttribute", w.getAttribute)
	jsClass.CreateOperation("getAttributeNS", w.getAttributeNS)
	jsClass.CreateOperation("setAttribute", w.setAttribute)
	jsClass.CreateOperation("setAttributeNS", w.setAttributeNS)
	jsClass.CreateOperation("removeAttribute", w.removeAttribute)
	jsClass.CreateOperation("removeAttributeNS", w.removeAttributeNS)
	jsClass.CreateOperation("toggleAttribute", w.toggleAttribute)
	jsClass.CreateOperation("hasAttribute", w.hasAttribute)
	jsClass.CreateOperation("hasAttributeNS", w.hasAttributeNS)
	jsClass.CreateOperation("getAttributeNode", w.getAttributeNode)
	jsClass.CreateOperation("getAttributeNodeNS", w.getAttributeNodeNS)
	jsClass.CreateOperation("setAttributeNode", w.setAttributeNode)
	jsClass.CreateOperation("setAttributeNodeNS", w.setAttributeNodeNS)
	jsClass.CreateOperation("removeAttributeNode", w.removeAttributeNode)
	jsClass.CreateOperation("attachShadow", w.attachShadow)
	jsClass.CreateOperation("closest", w.closest)
	jsClass.CreateOperation("matches", w.matches)
	jsClass.CreateOperation("getElementsByTagName", w.getElementsByTagName)
	jsClass.CreateOperation("getElementsByTagNameNS", w.getElementsByTagNameNS)
	jsClass.CreateOperation("getElementsByClassName", w.getElementsByClassName)
	jsClass.CreateOperation("insertAdjacentElement", w.insertAdjacentElement)
	jsClass.CreateOperation("insertAdjacentText", w.insertAdjacentText)
	jsClass.CreateOperation("insertAdjacentHTML", w.insertAdjacentHTML)
	jsClass.CreateAttribute("namespaceURI", w.namespaceURI, nil)
	jsClass.CreateAttribute("prefix", w.prefix, nil)
	jsClass.CreateAttribute("localName", w.localName, nil)
	jsClass.CreateAttribute("tagName", w.tagName, nil)
	jsClass.CreateAttribute("id", w.id, w.setID)
	jsClass.CreateAttribute("className", w.className, w.setClassName)
	jsClass.CreateAttribute("classList", w.classList, nil)
	jsClass.CreateAttribute("slot", w.slot, w.setSlot)
	jsClass.CreateAttribute("attributes", w.attributes, nil)
	jsClass.CreateAttribute("shadowRoot", w.shadowRoot, nil)
	jsClass.CreateAttribute("innerHTML", w.innerHTML, w.setInnerHTML)
	jsClass.CreateAttribute("outerHTML", w.outerHTML, w.setOuterHTML)
	w.parentNode.installPrototype(jsClass)
	w.nonDocumentTypeChildNode.installPrototype(jsClass)
	w.childNode.installPrototype(jsClass)
}

func (w Element[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Element[T]) hasAttributes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttributeNames(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.GetAttribute(qualifiedName)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w Element[T]) getAttributeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetAttribute(qualifiedName, value)
	return nil, nil
}

func (w Element[T]) setAttributeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) removeAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.RemoveAttribute(qualifiedName)
	return nil, nil
}

func (w Element[T]) removeAttributeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) toggleAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) hasAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.HasAttribute(qualifiedName)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Element[T]) hasAttributeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttributeNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttributeNodeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setAttributeNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setAttributeNodeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) removeAttributeNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) attachShadow(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) closest(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	selectors, errArg1 := js.ConsumeArgument(cbCtx, "selectors", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.Closest(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w Element[T]) matches(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	selectors, errArg1 := js.ConsumeArgument(cbCtx, "selectors", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.Matches(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Element[T]) getElementsByTagName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
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

func (w Element[T]) getElementsByTagNameNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
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

func (w Element[T]) getElementsByClassName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) insertAdjacentElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	where, errArg1 := js.ConsumeArgument(cbCtx, "where", nil, codec.DecodeString)
	element, errArg2 := js.ConsumeArgument(cbCtx, "element", nil, decodeElement)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result, errCall := instance.InsertAdjacentElement(where, element)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w Element[T]) insertAdjacentText(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	where, errArg1 := js.ConsumeArgument(cbCtx, "where", nil, codec.DecodeString)
	data, errArg2 := js.ConsumeArgument(cbCtx, "data", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.InsertAdjacentText(where, data)
	return nil, errCall
}

func (w Element[T]) insertAdjacentHTML(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	position, errArg1 := js.ConsumeArgument(cbCtx, "position", nil, codec.DecodeString)
	string, errArg2 := js.ConsumeArgument(cbCtx, "string", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	errCall := instance.InsertAdjacentHTML(position, string)
	return nil, errCall
}

func (w Element[T]) namespaceURI(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) prefix(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) localName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LocalName()
	return codec.EncodeString(cbCtx, result)
}

func (w Element[T]) tagName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TagName()
	return codec.EncodeString(cbCtx, result)
}

func (w Element[T]) id(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ID()
	return codec.EncodeString(cbCtx, result)
}

func (w Element[T]) setID(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetID(val)
	return nil, nil
}

func (w Element[T]) className(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setClassName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) slot(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setSlot(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) attributes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Attributes()
	return encodeNamedNodeMap(cbCtx, result)
}

func (w Element[T]) shadowRoot(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) innerHTML(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.InnerHTML()
	return codec.EncodeString(cbCtx, result)
}

func (w Element[T]) setInnerHTML(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetInnerHTML(val)
}

func (w Element[T]) outerHTML(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OuterHTML()
	return codec.EncodeString(cbCtx, result)
}

func (w Element[T]) setOuterHTML(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetOuterHTML(val)
}
