// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeElement[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("hasAttributes", Element_hasAttributes)
	jsClass.CreateOperation("getAttributeNames", Element_getAttributeNames)
	jsClass.CreateOperation("getAttribute", Element_getAttribute)
	jsClass.CreateOperation("getAttributeNS", Element_getAttributeNS)
	jsClass.CreateOperation("setAttribute", Element_setAttribute)
	jsClass.CreateOperation("setAttributeNS", Element_setAttributeNS)
	jsClass.CreateOperation("removeAttribute", Element_removeAttribute)
	jsClass.CreateOperation("removeAttributeNS", Element_removeAttributeNS)
	jsClass.CreateOperation("toggleAttribute", Element_toggleAttribute)
	jsClass.CreateOperation("hasAttribute", Element_hasAttribute)
	jsClass.CreateOperation("hasAttributeNS", Element_hasAttributeNS)
	jsClass.CreateOperation("getAttributeNode", Element_getAttributeNode)
	jsClass.CreateOperation("getAttributeNodeNS", Element_getAttributeNodeNS)
	jsClass.CreateOperation("setAttributeNode", Element_setAttributeNode)
	jsClass.CreateOperation("setAttributeNodeNS", Element_setAttributeNodeNS)
	jsClass.CreateOperation("removeAttributeNode", Element_removeAttributeNode)
	jsClass.CreateOperation("attachShadow", Element_attachShadow)
	jsClass.CreateOperation("closest", Element_closest)
	jsClass.CreateOperation("matches", Element_matches)
	jsClass.CreateOperation("getElementsByTagName", Element_getElementsByTagName)
	jsClass.CreateOperation("getElementsByTagNameNS", Element_getElementsByTagNameNS)
	jsClass.CreateOperation("getElementsByClassName", Element_getElementsByClassName)
	jsClass.CreateOperation("insertAdjacentElement", Element_insertAdjacentElement)
	jsClass.CreateOperation("insertAdjacentText", Element_insertAdjacentText)
	jsClass.CreateOperation("insertAdjacentHTML", Element_insertAdjacentHTML)
	jsClass.CreateAttribute("namespaceURI", Element_namespaceURI, nil)
	jsClass.CreateAttribute("prefix", Element_prefix, nil)
	jsClass.CreateAttribute("localName", Element_localName, nil)
	jsClass.CreateAttribute("tagName", Element_tagName, nil)
	jsClass.CreateAttribute("id", Element_id, Element_setID)
	jsClass.CreateAttribute("className", Element_className, Element_setClassName)
	jsClass.CreateAttribute("classList", Element_classList, nil)
	jsClass.CreateAttribute("slot", Element_slot, Element_setSlot)
	jsClass.CreateAttribute("attributes", Element_attributes, nil)
	jsClass.CreateAttribute("shadowRoot", Element_shadowRoot, nil)
	jsClass.CreateAttribute("innerHTML", Element_innerHTML, Element_setInnerHTML)
	jsClass.CreateAttribute("outerHTML", Element_outerHTML, Element_setOuterHTML)
	InitializeParentNode(jsClass)
	InitializeNonDocumentTypeChildNode(jsClass)
	InitializeChildNode(jsClass)
}

func ElementConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func Element_hasAttributes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_getAttributeNames[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_getAttribute[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_getAttributeNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_setAttribute[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_setAttributeNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_removeAttribute[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_removeAttributeNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_toggleAttribute[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_hasAttribute[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_hasAttributeNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_getAttributeNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_getAttributeNodeNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_setAttributeNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_setAttributeNodeNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_removeAttributeNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_attachShadow[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_closest[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_matches[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_getElementsByTagName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_getElementsByTagNameNS[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_getElementsByClassName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_insertAdjacentElement[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_insertAdjacentText[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_insertAdjacentHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Element_namespaceURI[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_prefix[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_localName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LocalName()
	return codec.EncodeString(cbCtx, result)
}

func Element_tagName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TagName()
	return codec.EncodeString(cbCtx, result)
}

func Element_id[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ID()
	return codec.EncodeString(cbCtx, result)
}

func Element_setID[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetID(val)
	return nil, nil
}

func Element_className[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_setClassName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_classList[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ClassList()
	return encodeDOMTokenList(cbCtx, result)
}

func Element_slot[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_setSlot[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_attributes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Attributes()
	return encodeNamedNodeMap(cbCtx, result)
}

func Element_shadowRoot[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Element.Element_shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Element_innerHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.InnerHTML()
	return codec.EncodeString(cbCtx, result)
}

func Element_setInnerHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetInnerHTML(val)
}

func Element_outerHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OuterHTML()
	return codec.EncodeString(cbCtx, result)
}

func Element_setOuterHTML[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	return nil, instance.SetOuterHTML(val)
}
