// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type parentNodeV8Wrapper struct {
	handleReffedObject[dom.ParentNode, jsTypeParam]
}

func newParentNodeV8Wrapper(scriptHost *V8ScriptHost) *parentNodeV8Wrapper {
	return &parentNodeV8Wrapper{newHandleReffedObject[dom.ParentNode](scriptHost)}
}

func createParentNodePrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newParentNodeV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper parentNodeV8Wrapper) initialize(jsClass v8Class) {
	wrapper.installPrototype(jsClass)
}

func (w parentNodeV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeMethod("prepend", w.prepend)
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("replaceChildren", w.replaceChildren)
	jsClass.CreatePrototypeMethod("querySelector", w.querySelector)
	jsClass.CreatePrototypeMethod("querySelectorAll", w.querySelectorAll)
	jsClass.CreatePrototypeAttribute("firstElementChild", w.firstElementChild, nil)
	jsClass.CreatePrototypeAttribute("lastElementChild", w.lastElementChild, nil)
	jsClass.CreatePrototypeAttribute("childElementCount", w.childElementCount, nil)
}

func (w parentNodeV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w parentNodeV8Wrapper) prepend(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.prepend")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	nodes, errArg1 := consumeRestArguments(cbCtx, "nodes", nil, w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Prepend(nodes...)
	return nil, errCall
}

func (w parentNodeV8Wrapper) append(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.append")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	nodes, errArg1 := consumeRestArguments(cbCtx, "nodes", nil, w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Append(nodes...)
	return nil, errCall
}

func (w parentNodeV8Wrapper) replaceChildren(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.replaceChildren")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	nodes, errArg1 := consumeRestArguments(cbCtx, "nodes", nil, w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.ReplaceChildren(nodes...)
	return nil, errCall
}

func (w parentNodeV8Wrapper) querySelector(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.querySelector")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	selectors, errArg1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.QuerySelector(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper) querySelectorAll(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.querySelectorAll")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	selectors, errArg1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.QuerySelectorAll(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper) firstElementChild(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.firstElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.FirstElementChild()
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper) lastElementChild(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.lastElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.LastElementChild()
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper) childElementCount(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.childElementCount")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ChildElementCount()
	return w.toUnsignedLong(cbCtx, result)
}
