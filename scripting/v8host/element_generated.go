// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("Element", "Node", createElementPrototype)
}

func createElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w elementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("hasAttributes", v8.NewFunctionTemplateWithError(iso, w.hasAttributes))
	prototypeTmpl.Set("getAttributeNames", v8.NewFunctionTemplateWithError(iso, w.getAttributeNames))
	prototypeTmpl.Set("getAttribute", v8.NewFunctionTemplateWithError(iso, w.getAttribute))
	prototypeTmpl.Set("getAttributeNS", v8.NewFunctionTemplateWithError(iso, w.getAttributeNS))
	prototypeTmpl.Set("setAttribute", v8.NewFunctionTemplateWithError(iso, w.setAttribute))
	prototypeTmpl.Set("setAttributeNS", v8.NewFunctionTemplateWithError(iso, w.setAttributeNS))
	prototypeTmpl.Set("removeAttribute", v8.NewFunctionTemplateWithError(iso, w.removeAttribute))
	prototypeTmpl.Set("removeAttributeNS", v8.NewFunctionTemplateWithError(iso, w.removeAttributeNS))
	prototypeTmpl.Set("toggleAttribute", v8.NewFunctionTemplateWithError(iso, w.toggleAttribute))
	prototypeTmpl.Set("hasAttribute", v8.NewFunctionTemplateWithError(iso, w.hasAttribute))
	prototypeTmpl.Set("hasAttributeNS", v8.NewFunctionTemplateWithError(iso, w.hasAttributeNS))
	prototypeTmpl.Set("getAttributeNode", v8.NewFunctionTemplateWithError(iso, w.getAttributeNode))
	prototypeTmpl.Set("getAttributeNodeNS", v8.NewFunctionTemplateWithError(iso, w.getAttributeNodeNS))
	prototypeTmpl.Set("setAttributeNode", v8.NewFunctionTemplateWithError(iso, w.setAttributeNode))
	prototypeTmpl.Set("setAttributeNodeNS", v8.NewFunctionTemplateWithError(iso, w.setAttributeNodeNS))
	prototypeTmpl.Set("removeAttributeNode", v8.NewFunctionTemplateWithError(iso, w.removeAttributeNode))
	prototypeTmpl.Set("attachShadow", v8.NewFunctionTemplateWithError(iso, w.attachShadow))
	prototypeTmpl.Set("matches", v8.NewFunctionTemplateWithError(iso, w.matches))
	prototypeTmpl.Set("getElementsByTagName", v8.NewFunctionTemplateWithError(iso, w.getElementsByTagName))
	prototypeTmpl.Set("getElementsByTagNameNS", v8.NewFunctionTemplateWithError(iso, w.getElementsByTagNameNS))
	prototypeTmpl.Set("getElementsByClassName", v8.NewFunctionTemplateWithError(iso, w.getElementsByClassName))
	prototypeTmpl.Set("insertAdjacentElement", v8.NewFunctionTemplateWithError(iso, w.insertAdjacentElement))
	prototypeTmpl.Set("insertAdjacentText", v8.NewFunctionTemplateWithError(iso, w.insertAdjacentText))

	prototypeTmpl.SetAccessorProperty("namespaceURI",
		v8.NewFunctionTemplateWithError(iso, w.namespaceURI),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("prefix",
		v8.NewFunctionTemplateWithError(iso, w.prefix),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("localName",
		v8.NewFunctionTemplateWithError(iso, w.localName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("tagName",
		v8.NewFunctionTemplateWithError(iso, w.tagName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("id",
		v8.NewFunctionTemplateWithError(iso, w.id),
		v8.NewFunctionTemplateWithError(iso, w.setID),
		v8.None)
	prototypeTmpl.SetAccessorProperty("className",
		v8.NewFunctionTemplateWithError(iso, w.className),
		v8.NewFunctionTemplateWithError(iso, w.setClassName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("classList",
		v8.NewFunctionTemplateWithError(iso, w.classList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("slot",
		v8.NewFunctionTemplateWithError(iso, w.slot),
		v8.NewFunctionTemplateWithError(iso, w.setSlot),
		v8.None)
	prototypeTmpl.SetAccessorProperty("attributes",
		v8.NewFunctionTemplateWithError(iso, w.attributes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("shadowRoot",
		v8.NewFunctionTemplateWithError(iso, w.shadowRoot),
		nil,
		v8.None)
	w.parentNode.installPrototype(prototypeTmpl)
	w.nonDocumentTypeChildNode.installPrototype(prototypeTmpl)
}

func (w elementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w elementV8Wrapper) hasAttributes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.hasAttributes")
	return nil, errors.New("Element.hasAttributes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) getAttributeNames(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.getAttributeNames")
	return nil, errors.New("Element.getAttributeNames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) getAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.getAttributeNS")
	return nil, errors.New("Element.getAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) setAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.setAttribute")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	qualifiedName, err1 := tryParseArg(args, 0, w.decodeString)
	value, err2 := tryParseArg(args, 1, w.decodeString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.SetAttribute(qualifiedName, value)
		return nil, nil
	}
	return nil, errors.New("Element.setAttribute: Missing arguments")
}

func (w elementV8Wrapper) setAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.setAttributeNS")
	return nil, errors.New("Element.setAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) removeAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.removeAttribute")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	qualifiedName, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.RemoveAttribute(qualifiedName)
		return nil, nil
	}
	return nil, errors.New("Element.removeAttribute: Missing arguments")
}

func (w elementV8Wrapper) removeAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.removeAttributeNS")
	return nil, errors.New("Element.removeAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) toggleAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.toggleAttribute")
	return nil, errors.New("Element.toggleAttribute: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) hasAttribute(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.hasAttribute")
	ctx := w.mustGetContext(info)
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	qualifiedName, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.HasAttribute(qualifiedName)
		return w.toBoolean(ctx, result)
	}
	return nil, errors.New("Element.hasAttribute: Missing arguments")
}

func (w elementV8Wrapper) hasAttributeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.hasAttributeNS")
	return nil, errors.New("Element.hasAttributeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) getAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.getAttributeNode")
	return nil, errors.New("Element.getAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) getAttributeNodeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.getAttributeNodeNS")
	return nil, errors.New("Element.getAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) setAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.setAttributeNode")
	return nil, errors.New("Element.setAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) setAttributeNodeNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.setAttributeNodeNS")
	return nil, errors.New("Element.setAttributeNodeNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) removeAttributeNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.removeAttributeNode")
	return nil, errors.New("Element.removeAttributeNode: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) attachShadow(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.attachShadow")
	return nil, errors.New("Element.attachShadow: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) matches(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.matches")
	ctx := w.mustGetContext(info)
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	selectors, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.Matches(selectors)
		if callErr != nil {
			return nil, callErr
		} else {
			return w.toBoolean(ctx, result)
		}
	}
	return nil, errors.New("Element.matches: Missing arguments")
}

func (w elementV8Wrapper) getElementsByTagName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.getElementsByTagName")
	return nil, errors.New("Element.getElementsByTagName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) getElementsByTagNameNS(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.getElementsByTagNameNS")
	return nil, errors.New("Element.getElementsByTagNameNS: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) getElementsByClassName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.getElementsByClassName")
	return nil, errors.New("Element.getElementsByClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) insertAdjacentElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.insertAdjacentElement")
	return nil, errors.New("Element.insertAdjacentElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) insertAdjacentText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.insertAdjacentText")
	return nil, errors.New("Element.insertAdjacentText: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) namespaceURI(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.namespaceURI")
	return nil, errors.New("Element.namespaceURI: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) prefix(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.prefix")
	return nil, errors.New("Element.prefix: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) localName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.localName")
	return nil, errors.New("Element.localName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) tagName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.tagName")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.TagName()
	return w.toString(ctx, result)
}

func (w elementV8Wrapper) id(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.id")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.ID()
	return w.toString(ctx, result)
}

func (w elementV8Wrapper) setID(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.setID")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetID(val)
	return nil, nil
}

func (w elementV8Wrapper) className(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.className")
	return nil, errors.New("Element.className: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) setClassName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.setClassName")
	return nil, errors.New("Element.setClassName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) slot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.slot")
	return nil, errors.New("Element.slot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) setSlot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.setSlot")
	return nil, errors.New("Element.setSlot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w elementV8Wrapper) attributes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.attributes")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Attributes()
	return w.toNamedNodeMap(ctx, result)
}

func (w elementV8Wrapper) shadowRoot(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Element.shadowRoot")
	return nil, errors.New("Element.shadowRoot: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
