// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("Element", "Node", createElementPrototype)
}

func createElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newElementV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w elementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("hasAttributes", wrapV8Callback(w.scriptHost, w.hasAttributes))
	prototypeTmpl.Set("getAttributeNames", wrapV8Callback(w.scriptHost, w.getAttributeNames))
	prototypeTmpl.Set("getAttribute", wrapV8Callback(w.scriptHost, w.getAttribute))
	prototypeTmpl.Set("getAttributeNS", wrapV8Callback(w.scriptHost, w.getAttributeNS))
	prototypeTmpl.Set("setAttribute", wrapV8Callback(w.scriptHost, w.setAttribute))
	prototypeTmpl.Set("setAttributeNS", wrapV8Callback(w.scriptHost, w.setAttributeNS))
	prototypeTmpl.Set("removeAttribute", wrapV8Callback(w.scriptHost, w.removeAttribute))
	prototypeTmpl.Set("removeAttributeNS", wrapV8Callback(w.scriptHost, w.removeAttributeNS))
	prototypeTmpl.Set("toggleAttribute", wrapV8Callback(w.scriptHost, w.toggleAttribute))
	prototypeTmpl.Set("hasAttribute", wrapV8Callback(w.scriptHost, w.hasAttribute))
	prototypeTmpl.Set("hasAttributeNS", wrapV8Callback(w.scriptHost, w.hasAttributeNS))
	prototypeTmpl.Set("getAttributeNode", wrapV8Callback(w.scriptHost, w.getAttributeNode))
	prototypeTmpl.Set("getAttributeNodeNS", wrapV8Callback(w.scriptHost, w.getAttributeNodeNS))
	prototypeTmpl.Set("setAttributeNode", wrapV8Callback(w.scriptHost, w.setAttributeNode))
	prototypeTmpl.Set("setAttributeNodeNS", wrapV8Callback(w.scriptHost, w.setAttributeNodeNS))
	prototypeTmpl.Set("removeAttributeNode", wrapV8Callback(w.scriptHost, w.removeAttributeNode))
	prototypeTmpl.Set("attachShadow", wrapV8Callback(w.scriptHost, w.attachShadow))
	prototypeTmpl.Set("matches", wrapV8Callback(w.scriptHost, w.matches))
	prototypeTmpl.Set("getElementsByTagName", wrapV8Callback(w.scriptHost, w.getElementsByTagName))
	prototypeTmpl.Set("getElementsByTagNameNS", wrapV8Callback(w.scriptHost, w.getElementsByTagNameNS))
	prototypeTmpl.Set("getElementsByClassName", wrapV8Callback(w.scriptHost, w.getElementsByClassName))
	prototypeTmpl.Set("insertAdjacentElement", wrapV8Callback(w.scriptHost, w.insertAdjacentElement))
	prototypeTmpl.Set("insertAdjacentText", wrapV8Callback(w.scriptHost, w.insertAdjacentText))

	prototypeTmpl.SetAccessorProperty("namespaceURI",
		wrapV8Callback(w.scriptHost, w.namespaceURI),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("prefix",
		wrapV8Callback(w.scriptHost, w.prefix),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("localName",
		wrapV8Callback(w.scriptHost, w.localName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("tagName",
		wrapV8Callback(w.scriptHost, w.tagName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("id",
		wrapV8Callback(w.scriptHost, w.id),
		wrapV8Callback(w.scriptHost, w.setID),
		v8.None)
	prototypeTmpl.SetAccessorProperty("className",
		wrapV8Callback(w.scriptHost, w.className),
		wrapV8Callback(w.scriptHost, w.setClassName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("classList",
		wrapV8Callback(w.scriptHost, w.classList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("slot",
		wrapV8Callback(w.scriptHost, w.slot),
		wrapV8Callback(w.scriptHost, w.setSlot),
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributes",
		wrapV8Callback(w.scriptHost, w.attributes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRoot",
		wrapV8Callback(w.scriptHost, w.shadowRoot),
		nil,
		v8.None)
	w.parentNode.installPrototype(prototypeTmpl)
	w.nonDocumentTypeChildNode.installPrototype(prototypeTmpl)
}

func (w elementV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w elementV8Wrapper) hasAttributes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.hasAttributes")
	return cbCtx.ReturnWithError(errors.New("Element.hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttributeNames(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.getAttributeNames")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.getAttribute")
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
	cbCtx.logger().Debug("V8 Function call: Element.getAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.setAttribute")
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
	cbCtx.logger().Debug("V8 Function call: Element.setAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) removeAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.removeAttribute")
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
	cbCtx.logger().Debug("V8 Function call: Element.removeAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) toggleAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.toggleAttribute")
	return cbCtx.ReturnWithError(errors.New("Element.toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) hasAttribute(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.hasAttribute")
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
	cbCtx.logger().Debug("V8 Function call: Element.hasAttributeNS")
	return cbCtx.ReturnWithError(errors.New("Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttributeNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.getAttributeNode")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getAttributeNodeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.getAttributeNodeNS")
	return cbCtx.ReturnWithError(errors.New("Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setAttributeNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.setAttributeNode")
	return cbCtx.ReturnWithError(errors.New("Element.setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setAttributeNodeNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.setAttributeNodeNS")
	return cbCtx.ReturnWithError(errors.New("Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) removeAttributeNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.removeAttributeNode")
	return cbCtx.ReturnWithError(errors.New("Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) attachShadow(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.attachShadow")
	return cbCtx.ReturnWithError(errors.New("Element.attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) matches(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.matches")
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
	cbCtx.logger().Debug("V8 Function call: Element.getElementsByTagName")
	return cbCtx.ReturnWithError(errors.New("Element.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getElementsByTagNameNS(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.getElementsByTagNameNS")
	return cbCtx.ReturnWithError(errors.New("Element.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) getElementsByClassName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.getElementsByClassName")
	return cbCtx.ReturnWithError(errors.New("Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) insertAdjacentElement(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.insertAdjacentElement")
	return cbCtx.ReturnWithError(errors.New("Element.insertAdjacentElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) insertAdjacentText(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.insertAdjacentText")
	return cbCtx.ReturnWithError(errors.New("Element.insertAdjacentText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) namespaceURI(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.namespaceURI")
	return cbCtx.ReturnWithError(errors.New("Element.namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) prefix(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.prefix")
	return cbCtx.ReturnWithError(errors.New("Element.prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) localName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.localName")
	return cbCtx.ReturnWithError(errors.New("Element.localName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) tagName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.tagName")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.TagName()
	return w.toString_(cbCtx, result)
}

func (w elementV8Wrapper) id(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.id")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ID()
	return w.toString_(cbCtx, result)
}

func (w elementV8Wrapper) setID(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.setID")
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
	cbCtx.logger().Debug("V8 Function call: Element.className")
	return cbCtx.ReturnWithError(errors.New("Element.className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setClassName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.setClassName")
	return cbCtx.ReturnWithError(errors.New("Element.setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) slot(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.slot")
	return cbCtx.ReturnWithError(errors.New("Element.slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) setSlot(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.setSlot")
	return cbCtx.ReturnWithError(errors.New("Element.setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w elementV8Wrapper) attributes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.attributes")
	instance, err := js.As[dom.Element](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Attributes()
	return w.toNamedNodeMap(cbCtx, result)
}

func (w elementV8Wrapper) shadowRoot(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Element.shadowRoot")
	return cbCtx.ReturnWithError(errors.New("Element.shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
