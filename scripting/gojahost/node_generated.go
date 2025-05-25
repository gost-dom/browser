// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	dom "github.com/gost-dom/browser/dom"
	log "github.com/gost-dom/browser/internal/log"
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

func (w nodeWrapper) Constructor(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.Constructor")
	cbCtx := newArgumentHelper(w.ctx, c)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeWrapper) getRootNode(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.getRootNode")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	options := w.decodeGetRootNodeOptions(c.Arguments[0])
	result := instance.GetRootNode(options)
	return cbCtx.ReturnWithValue(w.toNode(result))
}

func (w nodeWrapper) cloneNode(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.cloneNode")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	subtree := w.decodeboolean(c.Arguments[0])
	result := instance.CloneNode(subtree)
	return cbCtx.ReturnWithValue(w.toNode(result))
}

func (w nodeWrapper) isSameNode(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.isSameNode")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	otherNode := w.decodeNode(c.Arguments[0])
	result := instance.IsSameNode(otherNode)
	return cbCtx.ReturnWithValue(w.toBoolean(result))
}

func (w nodeWrapper) contains(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.contains")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	other := w.decodeNode(c.Arguments[0])
	result := instance.Contains(other)
	return cbCtx.ReturnWithValue(w.toBoolean(result))
}

func (w nodeWrapper) insertBefore(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.insertBefore")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	child := w.decodeNode(c.Arguments[1])
	result, err := instance.InsertBefore(node, child)
	if err != nil {
		panic(err)
	}
	return cbCtx.ReturnWithValue(w.toNode(result))
}

func (w nodeWrapper) appendChild(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.appendChild")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	node := w.decodeNode(c.Arguments[0])
	result, err := instance.AppendChild(node)
	if err != nil {
		panic(err)
	}
	return cbCtx.ReturnWithValue(w.toNode(result))
}

func (w nodeWrapper) removeChild(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.removeChild")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	child := w.decodeNode(c.Arguments[0])
	result, err := instance.RemoveChild(child)
	if err != nil {
		panic(err)
	}
	return cbCtx.ReturnWithValue(w.toNode(result))
}

func (w nodeWrapper) nodeName(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.nodeName")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.NodeName()
	return cbCtx.ReturnWithValue(w.toString_(result))
}

func (w nodeWrapper) isConnected(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.isConnected")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.IsConnected()
	return cbCtx.ReturnWithValue(w.toBoolean(result))
}

func (w nodeWrapper) ownerDocument(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.ownerDocument")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.OwnerDocument()
	return cbCtx.ReturnWithValue(w.toDocument(result))
}

func (w nodeWrapper) parentElement(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.parentElement")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.ParentElement()
	return cbCtx.ReturnWithValue(w.toElement(result))
}

func (w nodeWrapper) childNodes(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.childNodes")
	panic("Node.childNodes: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w nodeWrapper) firstChild(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.firstChild")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.FirstChild()
	return cbCtx.ReturnWithValue(w.toNode(result))
}

func (w nodeWrapper) previousSibling(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.previousSibling")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.PreviousSibling()
	return cbCtx.ReturnWithValue(w.toNode(result))
}

func (w nodeWrapper) nextSibling(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Node.nextSibling")
	cbCtx := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.NextSibling()
	return cbCtx.ReturnWithValue(w.toNode(result))
}
