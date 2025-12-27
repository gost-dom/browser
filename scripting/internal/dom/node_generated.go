// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
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
	jsClass.CreateOperation("getRootNode", w.getRootNode)
	jsClass.CreateOperation("cloneNode", w.cloneNode)
	jsClass.CreateOperation("isEqualNode", w.isEqualNode)
	jsClass.CreateOperation("isSameNode", w.isSameNode)
	jsClass.CreateOperation("contains", w.contains)
	jsClass.CreateOperation("insertBefore", w.insertBefore)
	jsClass.CreateOperation("appendChild", w.appendChild)
	jsClass.CreateOperation("replaceChild", w.replaceChild)
	jsClass.CreateOperation("removeChild", w.removeChild)
	jsClass.CreateAttribute("nodeType", w.nodeType, nil)
	jsClass.CreateAttribute("nodeName", w.nodeName, nil)
	jsClass.CreateAttribute("isConnected", w.isConnected, nil)
	jsClass.CreateAttribute("ownerDocument", w.ownerDocument, nil)
	jsClass.CreateAttribute("parentNode", w.parentNode, nil)
	jsClass.CreateAttribute("parentElement", w.parentElement, nil)
	jsClass.CreateAttribute("childNodes", w.childNodes, nil)
	jsClass.CreateAttribute("firstChild", w.firstChild, nil)
	jsClass.CreateAttribute("lastChild", w.lastChild, nil)
	jsClass.CreateAttribute("previousSibling", w.previousSibling, nil)
	jsClass.CreateAttribute("nextSibling", w.nextSibling, nil)
	jsClass.CreateAttribute("nodeValue", w.nodeValue, w.setNodeValue)
	jsClass.CreateAttribute("textContent", w.textContent, w.setTextContent)
}

func NodeConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Node[T]) getRootNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	options, errArg1 := js.ConsumeArgument(cbCtx, "options", w.defaultGetRootNodeOptions, decodeGetRootNodeOptions)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetRootNode(options)
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) cloneNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	child, errArg2 := js.ConsumeArgument(cbCtx, "child", codec.ZeroValue, codec.DecodeNode)
	err = gosterror.First(errArg1, errArg2)
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
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	node, errArg1 := js.ConsumeArgument(cbCtx, "node", nil, codec.DecodeNode)
	child, errArg2 := js.ConsumeArgument(cbCtx, "child", nil, codec.DecodeNode)
	err = gosterror.First(errArg1, errArg2)
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
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return codec.EncodeString(cbCtx, result)
}

func (w Node[T]) isConnected(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Node[T]) ownerDocument(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) parentNode(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentNode()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) parentElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentElement()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) childNodes(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) firstChild(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) lastChild(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LastChild()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) previousSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) nextSibling(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Node[T]) nodeValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, hasValue := instance.NodeValue()
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w Node[T]) setNodeValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetNodeValue(val)
	return nil, nil
}
