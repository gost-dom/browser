// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("Node", "EventTarget", createNodePrototype)
}

type nodeV8Wrapper struct {
	handleReffedObject[dom.Node]
}

func newNodeV8Wrapper(scriptHost *V8ScriptHost) *nodeV8Wrapper {
	return &nodeV8Wrapper{newHandleReffedObject[dom.Node](scriptHost)}
}

func createNodePrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newNodeV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w nodeV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("getRootNode", wrapV8Callback(w.scriptHost, w.getRootNode))
	prototypeTmpl.Set("cloneNode", wrapV8Callback(w.scriptHost, w.cloneNode))
	prototypeTmpl.Set("isSameNode", wrapV8Callback(w.scriptHost, w.isSameNode))
	prototypeTmpl.Set("contains", wrapV8Callback(w.scriptHost, w.contains))
	prototypeTmpl.Set("insertBefore", wrapV8Callback(w.scriptHost, w.insertBefore))
	prototypeTmpl.Set("appendChild", wrapV8Callback(w.scriptHost, w.appendChild))
	prototypeTmpl.Set("removeChild", wrapV8Callback(w.scriptHost, w.removeChild))

	prototypeTmpl.SetAccessorProperty("nodeType",
		wrapV8Callback(w.scriptHost, w.nodeType),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nodeName",
		wrapV8Callback(w.scriptHost, w.nodeName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("isConnected",
		wrapV8Callback(w.scriptHost, w.isConnected),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("ownerDocument",
		wrapV8Callback(w.scriptHost, w.ownerDocument),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("parentElement",
		wrapV8Callback(w.scriptHost, w.parentElement),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("childNodes",
		wrapV8Callback(w.scriptHost, w.childNodes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("firstChild",
		wrapV8Callback(w.scriptHost, w.firstChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("previousSibling",
		wrapV8Callback(w.scriptHost, w.previousSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nextSibling",
		wrapV8Callback(w.scriptHost, w.nextSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("textContent",
		wrapV8Callback(w.scriptHost, w.textContent),
		wrapV8Callback(w.scriptHost, w.setTextContent),
		v8.None)
}

func (w nodeV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeV8Wrapper) getRootNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.getRootNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	options, errArg1 := consumeArgument(cbCtx, "options", w.defaultGetRootNodeOptions, w.decodeGetRootNodeOptions)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetRootNode(options)
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) cloneNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.cloneNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	subtree, errArg1 := consumeArgument(cbCtx, "subtree", w.defaultboolean, w.decodeBoolean)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CloneNode(subtree)
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) isSameNode(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.isSameNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	otherNode, errArg1 := consumeArgument(cbCtx, "otherNode", zeroValue, w.decodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.IsSameNode(otherNode)
	return w.toBoolean(cbCtx, result)
}

func (w nodeV8Wrapper) contains(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.contains")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	other, errArg1 := consumeArgument(cbCtx, "other", zeroValue, w.decodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Contains(other)
	return w.toBoolean(cbCtx, result)
}

func (w nodeV8Wrapper) insertBefore(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.insertBefore")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	node, errArg1 := consumeArgument(cbCtx, "node", nil, w.decodeNode)
	child, errArg2 := consumeArgument(cbCtx, "child", zeroValue, w.decodeNode)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result, errCall := instance.InsertBefore(node, child)
	if errCall != nil {
		return nil, errCall
	}
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) appendChild(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.appendChild")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	node, errArg1 := consumeArgument(cbCtx, "node", nil, w.decodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.AppendChild(node)
	if errCall != nil {
		return nil, errCall
	}
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) removeChild(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.removeChild")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	child, errArg1 := consumeArgument(cbCtx, "child", nil, w.decodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.RemoveChild(child)
	if errCall != nil {
		return nil, errCall
	}
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) nodeName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.nodeName")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NodeName()
	return w.toString_(cbCtx, result)
}

func (w nodeV8Wrapper) isConnected(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.isConnected")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.IsConnected()
	return w.toBoolean(cbCtx, result)
}

func (w nodeV8Wrapper) ownerDocument(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.ownerDocument")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OwnerDocument()
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) parentElement(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.parentElement")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ParentElement()
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) childNodes(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.childNodes")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ChildNodes()
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) firstChild(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.firstChild")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.FirstChild()
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) previousSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.previousSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousSibling()
	return w.toJSWrapper(cbCtx, result)
}

func (w nodeV8Wrapper) nextSibling(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.logger().Debug("V8 Function call: Node.nextSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextSibling()
	return w.toJSWrapper(cbCtx, result)
}
