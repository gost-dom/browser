// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
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
	jsClass.CreatePrototypeMethod("hasAttributes", w.hasAttributes)
	jsClass.CreatePrototypeMethod("getAttributeNames", w.getAttributeNames)
	jsClass.CreatePrototypeMethod("getAttribute", w.getAttribute)
	jsClass.CreatePrototypeMethod("getAttributeNS", w.getAttributeNS)
	jsClass.CreatePrototypeMethod("setAttribute", w.setAttribute)
	jsClass.CreatePrototypeMethod("setAttributeNS", w.setAttributeNS)
	jsClass.CreatePrototypeMethod("removeAttribute", w.removeAttribute)
	jsClass.CreatePrototypeMethod("removeAttributeNS", w.removeAttributeNS)
	jsClass.CreatePrototypeMethod("toggleAttribute", w.toggleAttribute)
	jsClass.CreatePrototypeMethod("hasAttribute", w.hasAttribute)
	jsClass.CreatePrototypeMethod("hasAttributeNS", w.hasAttributeNS)
	jsClass.CreatePrototypeMethod("getAttributeNode", w.getAttributeNode)
	jsClass.CreatePrototypeMethod("getAttributeNodeNS", w.getAttributeNodeNS)
	jsClass.CreatePrototypeMethod("setAttributeNode", w.setAttributeNode)
	jsClass.CreatePrototypeMethod("setAttributeNodeNS", w.setAttributeNodeNS)
	jsClass.CreatePrototypeMethod("removeAttributeNode", w.removeAttributeNode)
	jsClass.CreatePrototypeMethod("attachShadow", w.attachShadow)
	jsClass.CreatePrototypeMethod("closest", w.closest)
	jsClass.CreatePrototypeMethod("matches", w.matches)
	jsClass.CreatePrototypeMethod("getElementsByTagName", w.getElementsByTagName)
	jsClass.CreatePrototypeMethod("getElementsByTagNameNS", w.getElementsByTagNameNS)
	jsClass.CreatePrototypeMethod("getElementsByClassName", w.getElementsByClassName)
	jsClass.CreatePrototypeMethod("insertAdjacentElement", w.insertAdjacentElement)
	jsClass.CreatePrototypeMethod("insertAdjacentText", w.insertAdjacentText)
	jsClass.CreatePrototypeAttribute("namespaceURI", w.namespaceURI, nil)
	jsClass.CreatePrototypeAttribute("prefix", w.prefix, nil)
	jsClass.CreatePrototypeAttribute("localName", w.localName, nil)
	jsClass.CreatePrototypeAttribute("tagName", w.tagName, nil)
	jsClass.CreatePrototypeAttribute("id", w.id, w.setID)
	jsClass.CreatePrototypeAttribute("className", w.className, w.setClassName)
	jsClass.CreatePrototypeAttribute("classList", w.classList, nil)
	jsClass.CreatePrototypeAttribute("slot", w.slot, w.setSlot)
	jsClass.CreatePrototypeAttribute("attributes", w.attributes, nil)
	jsClass.CreatePrototypeAttribute("shadowRoot", w.shadowRoot, nil)
	w.parentNode.installPrototype(jsClass)
	w.nonDocumentTypeChildNode.installPrototype(jsClass)
	w.childNode.installPrototype(jsClass)
}

func (w Element[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.Constructor", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Element[T]) hasAttributes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.hasAttributes", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttributeNames(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getAttributeNames", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getAttribute", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
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
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getAttributeNS", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.setAttribute", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetAttribute(qualifiedName, value)
	return nil, nil
}

func (w Element[T]) setAttributeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.setAttributeNS", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) removeAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.removeAttribute", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
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
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.removeAttributeNS", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) toggleAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.toggleAttribute", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) hasAttribute(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.hasAttribute", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
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
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.hasAttributeNS", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttributeNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getAttributeNode", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getAttributeNodeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getAttributeNodeNS", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setAttributeNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.setAttributeNode", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setAttributeNodeNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.setAttributeNodeNS", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) removeAttributeNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.removeAttributeNode", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) attachShadow(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.attachShadow", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) closest(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.closest", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
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
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.matches", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
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
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getElementsByTagName", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getElementsByTagNameNS(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getElementsByTagNameNS", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) getElementsByClassName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.getElementsByClassName", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) insertAdjacentElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.insertAdjacentElement", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	where, errArg1 := js.ConsumeArgument(cbCtx, "where", nil, codec.DecodeString)
	element, errArg2 := js.ConsumeArgument(cbCtx, "element", nil, w.decodeElement)
	err = errors.Join(errArg1, errArg2)
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
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.insertAdjacentText", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.insertAdjacentText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) namespaceURI(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.namespaceURI", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) prefix(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.prefix", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) localName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.localName", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.localName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) tagName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.tagName", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TagName()
	return codec.EncodeString(cbCtx, result)
}

func (w Element[T]) id(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.id", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ID()
	return codec.EncodeString(cbCtx, result)
}

func (w Element[T]) setID(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.setID", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetID(val)
	return nil, nil
}

func (w Element[T]) className(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.className", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setClassName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.setClassName", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) slot(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.slot", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) setSlot(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.setSlot", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Element[T]) attributes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.attributes", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Attributes()
	return w.toNamedNodeMap(cbCtx, result)
}

func (w Element[T]) shadowRoot(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Element.shadowRoot", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Element.shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
