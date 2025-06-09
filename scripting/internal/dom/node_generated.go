// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type NodeV8Wrapper[T any] struct{}

func NewNodeV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *NodeV8Wrapper[T] {
	return &NodeV8Wrapper[T]{}
}

func (wrapper NodeV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w NodeV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w NodeV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w NodeV8Wrapper[T]) getRootNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.getRootNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	options, errArg1 := js.ConsumeArgument(cbCtx, "options", w.defaultGetRootNodeOptions, w.decodeGetRootNodeOptions)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetRootNode(options)
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) cloneNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.cloneNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	subtree, errArg1 := js.ConsumeArgument(cbCtx, "subtree", w.defaultboolean, codec.DecodeBoolean)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CloneNode(subtree)
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) isSameNode(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.isSameNode")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	otherNode, errArg1 := js.ConsumeArgument(cbCtx, "otherNode", codec.ZeroValue, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.IsSameNode(otherNode)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w NodeV8Wrapper[T]) contains(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.contains")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	other, errArg1 := js.ConsumeArgument(cbCtx, "other", codec.ZeroValue, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Contains(other)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w NodeV8Wrapper[T]) insertBefore(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.insertBefore")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	child, errArg2 := js.ConsumeArgument(cbCtx, "child", codec.ZeroValue, codec.DecodeNode)
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

func (w NodeV8Wrapper[T]) appendChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.appendChild")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.AppendChild(node)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) removeChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.removeChild")
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	child, errArg1 := js.ConsumeArgument(cbCtx, "child", nil, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.RemoveChild(child)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) nodeName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.nodeName")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return codec.EncodeString(cbCtx, result)
}

func (w NodeV8Wrapper[T]) isConnected(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.isConnected")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w NodeV8Wrapper[T]) ownerDocument(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.ownerDocument")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) parentElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.parentElement")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentElement()
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) childNodes(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.childNodes")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) firstChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.firstChild")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) previousSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.previousSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeV8Wrapper[T]) nextSibling(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Node.nextSibling")
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return codec.EncodeEntity(cbCtx, result)
}
