// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
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
	iso := scriptHost.iso
	wrapper := newNodeV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

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

func (w nodeV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeV8Wrapper) getRootNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.getRootNode")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](cbCtx.Instance())
	options, err1 := consumeArgument(cbCtx, "options", w.defaultGetRootNodeOptions, w.decodeGetRootNodeOptions)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetRootNode(options)
		return cbCtx.Context().getInstanceForNode(result)
	}
	return nil, errors.New("Node.getRootNode: Missing arguments")
}

func (w nodeV8Wrapper) cloneNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.cloneNode")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](cbCtx.Instance())
	subtree, err1 := consumeArgument(cbCtx, "subtree", w.defaultboolean, w.decodeBoolean)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.CloneNode(subtree)
		return cbCtx.Context().getInstanceForNode(result)
	}
	return nil, errors.New("Node.cloneNode: Missing arguments")
}

func (w nodeV8Wrapper) isSameNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.isSameNode")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](cbCtx.Instance())
	otherNode, err1 := consumeArgument(cbCtx, "otherNode", zeroValue, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.IsSameNode(otherNode)
		return w.toBoolean(cbCtx.Context(), result)
	}
	return nil, errors.New("Node.isSameNode: Missing arguments")
}

func (w nodeV8Wrapper) contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.contains")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](cbCtx.Instance())
	other, err1 := consumeArgument(cbCtx, "other", zeroValue, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(other)
		return w.toBoolean(cbCtx.Context(), result)
	}
	return nil, errors.New("Node.contains: Missing arguments")
}

func (w nodeV8Wrapper) insertBefore(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.insertBefore")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](cbCtx.Instance())
	node, err1 := consumeArgument(cbCtx, "node", nil, w.decodeNode)
	child, err2 := consumeArgument(cbCtx, "child", zeroValue, w.decodeNode)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.InsertBefore(node, child)
		if callErr != nil {
			return nil, callErr
		} else {
			return cbCtx.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("Node.insertBefore: Missing arguments")
}

func (w nodeV8Wrapper) appendChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.appendChild")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](cbCtx.Instance())
	node, err1 := consumeArgument(cbCtx, "node", nil, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.AppendChild(node)
		if callErr != nil {
			return nil, callErr
		} else {
			return cbCtx.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("Node.appendChild: Missing arguments")
}

func (w nodeV8Wrapper) removeChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.removeChild")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](cbCtx.Instance())
	child, err1 := consumeArgument(cbCtx, "child", nil, w.decodeNode)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.RemoveChild(child)
		if callErr != nil {
			return nil, callErr
		} else {
			return cbCtx.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("Node.removeChild: Missing arguments")
}

func (w nodeV8Wrapper) nodeName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.nodeName")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return w.toString_(cbCtx.Context(), result)
}

func (w nodeV8Wrapper) isConnected(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.isConnected")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return w.toBoolean(cbCtx.Context(), result)
}

func (w nodeV8Wrapper) ownerDocument(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.ownerDocument")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) parentElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.parentElement")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentElement()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) childNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.childNodes")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return w.toNodeList(cbCtx.Context(), result)
}

func (w nodeV8Wrapper) firstChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.firstChild")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) previousSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.previousSibling")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) nextSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.nextSibling")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return cbCtx.Context().getInstanceForNode(result)
}
