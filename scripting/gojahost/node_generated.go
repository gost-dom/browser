// This file is generated. Do not edit.

package gojahost

import (
	"errors"
	g "github.com/dop251/goja"
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	installClass("Node", "EventTarget", newNodeWrapper)
}

type nodeWrapper struct {
	baseInstanceWrapper[dom.Node]
}

func newNodeWrapper(instance *GojaContext) wrapper {
	return &nodeWrapper{newBaseInstanceWrapper[dom.Node](instance)}
}

func (w nodeWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("getRootNode", wrapCallback(w.ctx, w.getRootNode))
	prototype.Set("cloneNode", wrapCallback(w.ctx, w.cloneNode))
	prototype.Set("isSameNode", wrapCallback(w.ctx, w.isSameNode))
	prototype.Set("contains", wrapCallback(w.ctx, w.contains))
	prototype.Set("insertBefore", wrapCallback(w.ctx, w.insertBefore))
	prototype.Set("appendChild", wrapCallback(w.ctx, w.appendChild))
	prototype.Set("removeChild", wrapCallback(w.ctx, w.removeChild))
	prototype.DefineAccessorProperty("nodeType", wrapCallback(w.ctx, w.nodeType), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("nodeName", wrapCallback(w.ctx, w.nodeName), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("isConnected", wrapCallback(w.ctx, w.isConnected), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("ownerDocument", wrapCallback(w.ctx, w.ownerDocument), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("parentElement", wrapCallback(w.ctx, w.parentElement), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("childNodes", wrapCallback(w.ctx, w.childNodes), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("firstChild", wrapCallback(w.ctx, w.firstChild), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("previousSibling", wrapCallback(w.ctx, w.previousSibling), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("nextSibling", wrapCallback(w.ctx, w.nextSibling), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("textContent", wrapCallback(w.ctx, w.textContent), wrapCallback(w.ctx, w.setTextContent), g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w nodeWrapper) Constructor(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeWrapper) getRootNode(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.getRootNode")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	options := w.decodeGetRootNodeOptions(cbCtx.Argument(0))
	result := instance.GetRootNode(options)
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) cloneNode(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.cloneNode")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	subtree := w.decodeboolean(cbCtx.Argument(0))
	result := instance.CloneNode(subtree)
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) isSameNode(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.isSameNode")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	otherNode := w.decodeNode(cbCtx.Argument(0))
	result := instance.IsSameNode(otherNode)
	return cbCtx.ReturnWithValue(w.toBoolean(result))
}

func (w nodeWrapper) contains(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.contains")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	other := w.decodeNode(cbCtx.Argument(0))
	result := instance.Contains(other)
	return cbCtx.ReturnWithValue(w.toBoolean(result))
}

func (w nodeWrapper) insertBefore(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.insertBefore")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	node := w.decodeNode(cbCtx.Argument(0))
	child := w.decodeNode(cbCtx.Argument(1))
	result, err := instance.InsertBefore(node, child)
	if err != nil {
		panic(err)
	}
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) appendChild(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.appendChild")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	node := w.decodeNode(cbCtx.Argument(0))
	result, err := instance.AppendChild(node)
	if err != nil {
		panic(err)
	}
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) removeChild(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.removeChild")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	child := w.decodeNode(cbCtx.Argument(0))
	result, err := instance.RemoveChild(child)
	if err != nil {
		panic(err)
	}
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) nodeName(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.nodeName")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.NodeName()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w nodeWrapper) isConnected(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.isConnected")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.IsConnected()
	return cbCtx.ReturnWithValue(w.toBoolean(result))
}

func (w nodeWrapper) ownerDocument(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.ownerDocument")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.OwnerDocument()
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) parentElement(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.parentElement")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.ParentElement()
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) childNodes(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.childNodes")
	return cbCtx.ReturnWithError(errors.New("Node.childNodes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w nodeWrapper) firstChild(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.firstChild")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.FirstChild()
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) previousSibling(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.previousSibling")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.PreviousSibling()
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}

func (w nodeWrapper) nextSibling(cbCtx *callbackContext) g.Value {
	cbCtx.logger().Debug("V8 Function call: Node.nextSibling")
	instance, instErr := js.As[dom.Node](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.NextSibling()
	return cbCtx.ReturnWithValue(w.toJSWrapper(result))
}
