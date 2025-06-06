// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type nodeV8Wrapper[T any] struct{}

func newNodeV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *nodeV8Wrapper[T] {
	return &nodeV8Wrapper[T]{}
}

func (wrapper nodeV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w nodeV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("getRootNode", w.getRootNode)
	jsClass.CreatePrototypeMethod("cloneNode", w.cloneNode)
	jsClass.CreatePrototypeMethod("isSameNode", w.isSameNode)
	jsClass.CreatePrototypeMethod("contains", w.contains)
	jsClass.CreatePrototypeMethod("insertBefore", w.insertBefore)
	jsClass.CreatePrototypeMethod("appendChild", w.appendChild)
	jsClass.CreatePrototypeMethod("removeChild", w.removeChild)
	jsClass.CreatePrototypeAttribute("nodeType", w.nodeType, nil)
	jsClass.CreatePrototypeAttribute("nodeName", w.nodeName, nil)
	jsClass.CreatePrototypeAttribute("isConnected", w.isConnected, nil)
	jsClass.CreatePrototypeAttribute("ownerDocument", w.ownerDocument, nil)
	jsClass.CreatePrototypeAttribute("parentElement", w.parentElement, nil)
	jsClass.CreatePrototypeAttribute("childNodes", w.childNodes, nil)
	jsClass.CreatePrototypeAttribute("firstChild", w.firstChild, nil)
	jsClass.CreatePrototypeAttribute("previousSibling", w.previousSibling, nil)
	jsClass.CreatePrototypeAttribute("nextSibling", w.nextSibling, nil)
	jsClass.CreatePrototypeAttribute("textContent", w.textContent, w.setTextContent)
}

func (w nodeV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w nodeV8Wrapper[T]) getRootNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.getRootNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	options, errArg1 := consumeArgument(cbCtx, "options", w.defaultGetRootNodeOptions, w.decodeGetRootNodeOptions)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetRootNode(options)
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) cloneNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.cloneNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	subtree, errArg1 := consumeArgument(cbCtx, "subtree", w.defaultboolean, codec.DecodeBoolean)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CloneNode(subtree)
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) isSameNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.isSameNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	otherNode, errArg1 := consumeArgument(cbCtx, "otherNode", zeroValue, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.IsSameNode(otherNode)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w nodeV8Wrapper[T]) contains(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.contains")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	other, errArg1 := consumeArgument(cbCtx, "other", zeroValue, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Contains(other)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w nodeV8Wrapper[T]) insertBefore(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.insertBefore")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	node, errArg1 := consumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	child, errArg2 := consumeArgument(cbCtx, "child", zeroValue, codec.DecodeNode)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result, errCall := instance.InsertBefore(node, child)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) appendChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.appendChild")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	node, errArg1 := consumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.AppendChild(node)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) removeChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.removeChild")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	child, errArg1 := consumeArgument(cbCtx, "child", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.RemoveChild(child)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) nodeName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.nodeName")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NodeName()
	return codec.EncodeString(cbCtx, result)
}

func (w nodeV8Wrapper[T]) isConnected(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.isConnected")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.IsConnected()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w nodeV8Wrapper[T]) ownerDocument(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.ownerDocument")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.OwnerDocument()
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) parentElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.parentElement")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ParentElement()
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) childNodes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.childNodes")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ChildNodes()
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) firstChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.firstChild")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.FirstChild()
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) previousSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.previousSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.PreviousSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func (w nodeV8Wrapper[T]) nextSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.nextSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.NextSibling()
	return codec.EncodeEntity(cbCtx, result)
}
