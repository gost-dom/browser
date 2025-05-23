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
	iso := w.scriptHost.iso
	prototypeTmpl.Set("getRootNode", v8.NewFunctionTemplateWithError(iso, w.getRootNode))
	prototypeTmpl.Set("cloneNode", v8.NewFunctionTemplateWithError(iso, w.cloneNode))
	prototypeTmpl.Set("isSameNode", v8.NewFunctionTemplateWithError(iso, w.isSameNode))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, w.contains))
	prototypeTmpl.Set("insertBefore", v8.NewFunctionTemplateWithError(iso, w.insertBefore))
	prototypeTmpl.Set("appendChild", v8.NewFunctionTemplateWithError(iso, w.appendChild))
	prototypeTmpl.Set("removeChild", v8.NewFunctionTemplateWithError(iso, w.removeChild))

	prototypeTmpl.SetAccessorProperty("nodeType",
		v8.NewFunctionTemplateWithError(iso, w.nodeType),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nodeName",
		v8.NewFunctionTemplateWithError(iso, w.nodeName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("isConnected",
		v8.NewFunctionTemplateWithError(iso, w.isConnected),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("ownerDocument",
		v8.NewFunctionTemplateWithError(iso, w.ownerDocument),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("parentElement",
		v8.NewFunctionTemplateWithError(iso, w.parentElement),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("childNodes",
		v8.NewFunctionTemplateWithError(iso, w.childNodes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("firstChild",
		v8.NewFunctionTemplateWithError(iso, w.firstChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("previousSibling",
		v8.NewFunctionTemplateWithError(iso, w.previousSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nextSibling",
		v8.NewFunctionTemplateWithError(iso, w.nextSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("textContent",
		v8.NewFunctionTemplateWithError(iso, w.textContent),
		v8.NewFunctionTemplateWithError(iso, w.setTextContent),
		v8.None)
}

func (w nodeV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w nodeV8Wrapper) getRootNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.getRootNode")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](args.Instance())
	options, err1 := tryParseArgWithDefault(args, 0, w.defaultGetRootNodeOptions, w.decodeGetRootNodeOptions)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetRootNode(options)
		return args.Context().getInstanceForNode(result)
	}
	return nil, errors.New("Node.getRootNode: Missing arguments")
}

func (w nodeV8Wrapper) cloneNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.cloneNode")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](args.Instance())
	subtree, err1 := tryParseArgWithDefault(args, 0, w.defaultboolean, w.decodeBoolean)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.CloneNode(subtree)
		return args.Context().getInstanceForNode(result)
	}
	return nil, errors.New("Node.cloneNode: Missing arguments")
}

func (w nodeV8Wrapper) isSameNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.isSameNode")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](args.Instance())
	otherNode, err1 := tryParseArgNullableType(args, 0, w.decodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.IsSameNode(otherNode)
		return w.toBoolean(args.Context(), result)
	}
	return nil, errors.New("Node.isSameNode: Missing arguments")
}

func (w nodeV8Wrapper) contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.contains")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](args.Instance())
	other, err1 := tryParseArgNullableType(args, 0, w.decodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(other)
		return w.toBoolean(args.Context(), result)
	}
	return nil, errors.New("Node.contains: Missing arguments")
}

func (w nodeV8Wrapper) insertBefore(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.insertBefore")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](args.Instance())
	node, err1 := tryParseArg(args, 0, w.decodeNode)
	child, err2 := tryParseArgNullableType(args, 1, w.decodeNode)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.InsertBefore(node, child)
		if callErr != nil {
			return nil, callErr
		} else {
			return args.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("Node.insertBefore: Missing arguments")
}

func (w nodeV8Wrapper) appendChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.appendChild")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](args.Instance())
	node, err1 := tryParseArg(args, 0, w.decodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.AppendChild(node)
		if callErr != nil {
			return nil, callErr
		} else {
			return args.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("Node.appendChild: Missing arguments")
}

func (w nodeV8Wrapper) removeChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.removeChild")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.Node](args.Instance())
	child, err1 := tryParseArg(args, 0, w.decodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.RemoveChild(child)
		if callErr != nil {
			return nil, callErr
		} else {
			return args.Context().getInstanceForNode(result)
		}
	}
	return nil, errors.New("Node.removeChild: Missing arguments")
}

func (w nodeV8Wrapper) nodeName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.nodeName")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return w.toString_(args.Context(), result)
}

func (w nodeV8Wrapper) isConnected(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.isConnected")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return w.toBoolean(args.Context(), result)
}

func (w nodeV8Wrapper) ownerDocument(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.ownerDocument")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return args.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) parentElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.parentElement")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentElement()
	return args.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) childNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.childNodes")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return w.toNodeList(args.Context(), result)
}

func (w nodeV8Wrapper) firstChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.firstChild")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return args.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) previousSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.previousSibling")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return args.Context().getInstanceForNode(result)
}

func (w nodeV8Wrapper) nextSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Node.nextSibling")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.Node](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return args.Context().getInstanceForNode(result)
}
