// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type ElementV8Wrapper[T any] struct {
	parentNode               *ParentNodeV8Wrapper[T]
	nonDocumentTypeChildNode *NonDocumentTypeChildNodeV8Wrapper[T]
}

func NewElementV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *ElementV8Wrapper[T] {
	return &ElementV8Wrapper[T]{
		NewParentNodeV8Wrapper(scriptHost),
		NewNonDocumentTypeChildNodeV8Wrapper(scriptHost),
	}
}

func (wrapper ElementV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w ElementV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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
}

func (w ElementV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w ElementV8Wrapper[T]) hasAttributes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.hasAttributes")
	return nil, errors.New("Element.hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) getAttributeNames(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNames")
	return nil, errors.New("Element.getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) getAttribute(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttribute")
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

func (w ElementV8Wrapper[T]) getAttributeNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNS")
	return nil, errors.New("Element.getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) setAttribute(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttribute")
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	qualifiedName, errArg1 := js.ConsumeArgument(cbCtx, "qualifiedName", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetAttribute(qualifiedName, value)
	return nil, nil
}

func (w ElementV8Wrapper[T]) setAttributeNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttributeNS")
	return nil, errors.New("Element.setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) removeAttribute(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.removeAttribute")
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

func (w ElementV8Wrapper[T]) removeAttributeNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.removeAttributeNS")
	return nil, errors.New("Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) toggleAttribute(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.toggleAttribute")
	return nil, errors.New("Element.toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) hasAttribute(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.hasAttribute")
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

func (w ElementV8Wrapper[T]) hasAttributeNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.hasAttributeNS")
	return nil, errors.New("Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) getAttributeNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNode")
	return nil, errors.New("Element.getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) getAttributeNodeNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNodeNS")
	return nil, errors.New("Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) setAttributeNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttributeNode")
	return nil, errors.New("Element.setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) setAttributeNodeNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttributeNodeNS")
	return nil, errors.New("Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) removeAttributeNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.removeAttributeNode")
	return nil, errors.New("Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) attachShadow(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.attachShadow")
	return nil, errors.New("Element.attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) matches(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.matches")
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

func (w ElementV8Wrapper[T]) getElementsByTagName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getElementsByTagName")
	return nil, errors.New("Element.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) getElementsByTagNameNS(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getElementsByTagNameNS")
	return nil, errors.New("Element.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) getElementsByClassName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getElementsByClassName")
	return nil, errors.New("Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) insertAdjacentElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.insertAdjacentElement")
	return nil, errors.New("Element.insertAdjacentElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) insertAdjacentText(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.insertAdjacentText")
	return nil, errors.New("Element.insertAdjacentText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) namespaceURI(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.namespaceURI")
	return nil, errors.New("Element.namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) prefix(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.prefix")
	return nil, errors.New("Element.prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) localName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.localName")
	return nil, errors.New("Element.localName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) tagName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.tagName")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TagName()
	return codec.EncodeString(cbCtx, result)
}

func (w ElementV8Wrapper[T]) id(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.id")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ID()
	return codec.EncodeString(cbCtx, result)
}

func (w ElementV8Wrapper[T]) setID(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setID")
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetID(val)
	return nil, nil
}

func (w ElementV8Wrapper[T]) className(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.className")
	return nil, errors.New("Element.className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) setClassName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setClassName")
	return nil, errors.New("Element.setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) slot(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.slot")
	return nil, errors.New("Element.slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) setSlot(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setSlot")
	return nil, errors.New("Element.setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w ElementV8Wrapper[T]) attributes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.attributes")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Attributes()
	return w.toNamedNodeMap(cbCtx, result)
}

func (w ElementV8Wrapper[T]) shadowRoot(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Element.shadowRoot")
	return nil, errors.New("Element.shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
