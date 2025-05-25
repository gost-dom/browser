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

func (w nodeV8Wrapper) Constructor(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeV8Wrapper) getRootNode(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.getRootNode")
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	options, err1 := consumeArgument(cbCtx, "options", w.defaultGetRootNodeOptions, w.decodeGetRootNodeOptions)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.GetRootNode(options)
		return cbCtx.getInstanceForNode(result)
	}
	return cbCtx.ReturnWithError(errors.New("Node.getRootNode: Missing arguments"))
}

func (w nodeV8Wrapper) cloneNode(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.cloneNode")
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	subtree, err1 := consumeArgument(cbCtx, "subtree", w.defaultboolean, w.decodeBoolean)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.CloneNode(subtree)
		return cbCtx.getInstanceForNode(result)
	}
	return cbCtx.ReturnWithError(errors.New("Node.cloneNode: Missing arguments"))
}

func (w nodeV8Wrapper) isSameNode(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.isSameNode")
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	otherNode, err1 := consumeArgument(cbCtx, "otherNode", zeroValue, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.IsSameNode(otherNode)
		return w.toBoolean(cbCtx, result)
	}
	return cbCtx.ReturnWithError(errors.New("Node.isSameNode: Missing arguments"))
}

func (w nodeV8Wrapper) contains(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.contains")
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	other, err1 := consumeArgument(cbCtx, "other", zeroValue, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result := instance.Contains(other)
		return w.toBoolean(cbCtx, result)
	}
	return cbCtx.ReturnWithError(errors.New("Node.contains: Missing arguments"))
}

func (w nodeV8Wrapper) insertBefore(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.insertBefore")
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	node, err1 := consumeArgument(cbCtx, "node", nil, w.decodeNode)
	child, err2 := consumeArgument(cbCtx, "child", zeroValue, w.decodeNode)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result, callErr := instance.InsertBefore(node, child)
		if callErr != nil {
			return cbCtx.ReturnWithError(callErr)
		} else {
			return cbCtx.getInstanceForNode(result)
		}
	}
	return cbCtx.ReturnWithError(errors.New("Node.insertBefore: Missing arguments"))
}

func (w nodeV8Wrapper) appendChild(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.appendChild")
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	node, err1 := consumeArgument(cbCtx, "node", nil, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result, callErr := instance.AppendChild(node)
		if callErr != nil {
			return cbCtx.ReturnWithError(callErr)
		} else {
			return cbCtx.getInstanceForNode(result)
		}
	}
	return cbCtx.ReturnWithError(errors.New("Node.appendChild: Missing arguments"))
}

func (w nodeV8Wrapper) removeChild(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.removeChild")
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	child, err1 := consumeArgument(cbCtx, "child", nil, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		result, callErr := instance.RemoveChild(child)
		if callErr != nil {
			return cbCtx.ReturnWithError(callErr)
		} else {
			return cbCtx.getInstanceForNode(result)
		}
	}
	return cbCtx.ReturnWithError(errors.New("Node.removeChild: Missing arguments"))
}

func (w nodeV8Wrapper) nodeName(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.nodeName")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NodeName()
	return w.toString_(cbCtx, result)
}

func (w nodeV8Wrapper) isConnected(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.isConnected")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.IsConnected()
	return w.toBoolean(cbCtx, result)
}

func (w nodeV8Wrapper) ownerDocument(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.ownerDocument")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OwnerDocument()
	return cbCtx.getInstanceForNode(result)
}

func (w nodeV8Wrapper) parentElement(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.parentElement")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ParentElement()
	return cbCtx.getInstanceForNode(result)
}

func (w nodeV8Wrapper) childNodes(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.childNodes")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ChildNodes()
	return w.toNodeList(cbCtx, result)
}

func (w nodeV8Wrapper) firstChild(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.firstChild")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.FirstChild()
	return cbCtx.getInstanceForNode(result)
}

func (w nodeV8Wrapper) previousSibling(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.previousSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousSibling()
	return cbCtx.getInstanceForNode(result)
}

func (w nodeV8Wrapper) nextSibling(cbCtx *argumentHelper) js.CallbackRVal {
	cbCtx.logger().Debug("V8 Function call: Node.nextSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextSibling()
	return cbCtx.getInstanceForNode(result)
}
