// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerClass("Element", "Node", newElementV8Wrapper)
}

func createElementPrototype(scriptHost *V8ScriptHost) jsClass {
	wrapper := newElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	wrapper.CustomInitializer(jsClass)
	return jsClass
}
func (wrapper elementV8Wrapper) initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w elementV8Wrapper) installPrototype(jsClass jsClass) {
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

func (w elementV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w elementV8Wrapper) hasAttributes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.hasAttributes")
	return cbCtx.ReturnWithError(errors.New("Element.hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttributeNames(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNames")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttribute")
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	qualifiedName, errArg1 := consumeArgument(cbCtx, "qualifiedName", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.GetAttribute(qualifiedName)
	return w.toNillableString_(cbCtx, result, hasValue)
}

func (w elementV8Wrapper) getAttributeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttribute")
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	qualifiedName, errArg1 := consumeArgument(cbCtx, "qualifiedName", nil, w.decodeString)
	value, errArg2 := consumeArgument(cbCtx, "value", nil, w.decodeString)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.SetAttribute(qualifiedName, value)
	return nil, nil
}

func (w elementV8Wrapper) setAttributeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) removeAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.removeAttribute")
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	qualifiedName, errArg1 := consumeArgument(cbCtx, "qualifiedName", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.RemoveAttribute(qualifiedName)
	return nil, nil
}

func (w elementV8Wrapper) removeAttributeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.removeAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) toggleAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.toggleAttribute")
	return cbCtx.ReturnWithError(errors.New("Element.toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) hasAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.hasAttribute")
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	qualifiedName, errArg1 := consumeArgument(cbCtx, "qualifiedName", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.HasAttribute(qualifiedName)
	return w.toBoolean(cbCtx, result)
}

func (w elementV8Wrapper) hasAttributeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.hasAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttributeNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNode")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttributeNodeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getAttributeNodeNS")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setAttributeNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttributeNode")
	return cbCtx.ReturnWithError(errors.New("Element.setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setAttributeNodeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setAttributeNodeNS")
	return cbCtx.ReturnWithError(errors.New("Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) removeAttributeNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.removeAttributeNode")
	return cbCtx.ReturnWithError(errors.New("Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) attachShadow(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.attachShadow")
	return cbCtx.ReturnWithError(errors.New("Element.attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) matches(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.matches")
	instance, errInst := js.As[dom.Element](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	selectors, errArg1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.Matches(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return w.toBoolean(cbCtx, result)
}

func (w elementV8Wrapper) getElementsByTagName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getElementsByTagName")
	return cbCtx.ReturnWithError(errors.New("Element.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getElementsByTagNameNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getElementsByTagNameNS")
	return cbCtx.ReturnWithError(errors.New("Element.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getElementsByClassName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.getElementsByClassName")
	return cbCtx.ReturnWithError(errors.New("Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) insertAdjacentElement(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.insertAdjacentElement")
	return cbCtx.ReturnWithError(errors.New("Element.insertAdjacentElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) insertAdjacentText(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.insertAdjacentText")
	return cbCtx.ReturnWithError(errors.New("Element.insertAdjacentText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) namespaceURI(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.namespaceURI")
	return cbCtx.ReturnWithError(errors.New("Element.namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) prefix(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.prefix")
	return cbCtx.ReturnWithError(errors.New("Element.prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) localName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.localName")
	return cbCtx.ReturnWithError(errors.New("Element.localName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) tagName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.tagName")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.TagName()
	return w.toString_(cbCtx, result)
}

func (w elementV8Wrapper) id(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.id")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ID()
	return w.toString_(cbCtx, result)
}

func (w elementV8Wrapper) setID(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setID")
	instance, err0 := js.As[dom.Element](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	instance.SetID(val)
	return cbCtx.ReturnWithValue(nil)
}

func (w elementV8Wrapper) className(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.className")
	return cbCtx.ReturnWithError(errors.New("Element.className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setClassName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setClassName")
	return cbCtx.ReturnWithError(errors.New("Element.setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) slot(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.slot")
	return cbCtx.ReturnWithError(errors.New("Element.slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setSlot(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.setSlot")
	return cbCtx.ReturnWithError(errors.New("Element.setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) attributes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.attributes")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Attributes()
	return w.toNamedNodeMap(cbCtx, result)
}

func (w elementV8Wrapper) shadowRoot(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Element.shadowRoot")
	return cbCtx.ReturnWithError(errors.New("Element.shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
