// This file is generated. Do not edit.

package dom

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Node[T any] struct{}

func NewNode[T any](scriptHost js.ScriptEngine[T]) *Node[T] {
	return &Node[T]{}
}

func (wrapper Node[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Node[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("getRootNode", w.getRootNode)
	jsClass.CreatePrototypeMethod("cloneNode", w.cloneNode)
	jsClass.CreatePrototypeMethod("isEqualNode", w.isEqualNode)
	jsClass.CreatePrototypeMethod("isSameNode", w.isSameNode)
	jsClass.CreatePrototypeMethod("contains", w.contains)
	jsClass.CreatePrototypeMethod("insertBefore", w.insertBefore)
	jsClass.CreatePrototypeMethod("appendChild", w.appendChild)
	jsClass.CreatePrototypeMethod("replaceChild", w.replaceChild)
	jsClass.CreatePrototypeMethod("removeChild", w.removeChild)
	jsClass.CreatePrototypeAttribute("nodeType", w.nodeType, nil)
	jsClass.CreatePrototypeAttribute("nodeName", w.nodeName, nil)
	jsClass.CreatePrototypeAttribute("isConnected", w.isConnected, nil)
	jsClass.CreatePrototypeAttribute("ownerDocument", w.ownerDocument, nil)
	jsClass.CreatePrototypeAttribute("parentNode", w.parentNode, nil)
	jsClass.CreatePrototypeAttribute("parentElement", w.parentElement, nil)
	jsClass.CreatePrototypeAttribute("childNodes", w.childNodes, nil)
	jsClass.CreatePrototypeAttribute("firstChild", w.firstChild, nil)
	jsClass.CreatePrototypeAttribute("previousSibling", w.previousSibling, nil)
	jsClass.CreatePrototypeAttribute("nextSibling", w.nextSibling, nil)
	jsClass.CreatePrototypeAttribute("textContent", w.textContent, w.setTextContent)
}

func (w Node[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.Constructor", js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Node[T]) getRootNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.getRootNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.getRootNode", js.LogAttr("res", res))
	}()
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

func (w Node[T]) cloneNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.cloneNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.cloneNode", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	subtree, errArg1 := js.ConsumeArgument(cbCtx, "subtree", codec.ZeroValue, codec.DecodeBoolean)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.CloneNode(subtree)
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) isEqualNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.isEqualNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.isEqualNode", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	otherNode, errArg1 := js.ConsumeArgument(cbCtx, "otherNode", codec.ZeroValue, codec.DecodeNode)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.IsEqualNode(otherNode)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Node[T]) isSameNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.isSameNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.isSameNode", js.LogAttr("res", res))
	}()
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

func (w Node[T]) contains(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.contains - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.contains", js.LogAttr("res", res))
	}()
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

func (w Node[T]) insertBefore(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.insertBefore - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.insertBefore", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	child, errArg2 := js.ConsumeArgument(cbCtx, "child", codec.ZeroValue, codec.DecodeNode)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result, errCall := instance.InsertBefore(node, child)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) appendChild(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.appendChild - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.appendChild", js.LogAttr("res", res))
	}()
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

func (w Node[T]) replaceChild(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.replaceChild", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	child, errArg2 := js.ConsumeArgument(cbCtx, "child", nil, codec.DecodeNode)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result, errCall := instance.ReplaceChild(node, child)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) removeChild(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.removeChild - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.removeChild", js.LogAttr("res", res))
	}()
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

func (w Node[T]) nodeName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.nodeName - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.nodeName", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return codec.EncodeString(cbCtx, result)
}

func (w Node[T]) isConnected(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.isConnected - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.isConnected", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Node[T]) ownerDocument(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.ownerDocument - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.ownerDocument", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) parentNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.parentNode - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.parentNode", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentNode()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) parentElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.parentElement - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.parentElement", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentElement()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) childNodes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.childNodes - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.childNodes", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) firstChild(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.firstChild - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.firstChild", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) previousSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.previousSibling - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.previousSibling", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) nextSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Node.nextSibling - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Node.nextSibling", js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return codec.EncodeEntity(cbCtx, result)
}
