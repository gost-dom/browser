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
	jsClass.CreateOperation("getRootNode", Node_getRootNode)
	jsClass.CreateOperation("cloneNode", Node_cloneNode)
	jsClass.CreateOperation("isEqualNode", Node_isEqualNode)
	jsClass.CreateOperation("isSameNode", Node_isSameNode)
	jsClass.CreateOperation("contains", Node_contains)
	jsClass.CreateOperation("insertBefore", Node_insertBefore)
	jsClass.CreateOperation("appendChild", Node_appendChild)
	jsClass.CreateOperation("replaceChild", Node_replaceChild)
	jsClass.CreateOperation("removeChild", Node_removeChild)
	jsClass.CreateAttribute("nodeType", Node_nodeType, nil)
	jsClass.CreateAttribute("nodeName", Node_nodeName, nil)
	jsClass.CreateAttribute("isConnected", Node_isConnected, nil)
	jsClass.CreateAttribute("ownerDocument", Node_ownerDocument, nil)
	jsClass.CreateAttribute("parentNode", Node_parentNode, nil)
	jsClass.CreateAttribute("parentElement", Node_parentElement, nil)
	jsClass.CreateAttribute("childNodes", Node_childNodes, nil)
	jsClass.CreateAttribute("firstChild", Node_firstChild, nil)
	jsClass.CreateAttribute("lastChild", Node_lastChild, nil)
	jsClass.CreateAttribute("previousSibling", Node_previousSibling, nil)
	jsClass.CreateAttribute("nextSibling", Node_nextSibling, nil)
	jsClass.CreateAttribute("nodeValue", Node_nodeValue, Node_setNodeValue)
	jsClass.CreateAttribute("textContent", Node_textContent, Node_setTextContent)
}

func NodeConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func Node_getRootNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.Node](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	options, errArg1 := js.ConsumeArgument(cbCtx, "options", codec.ZeroValue, decodeGetRootNodeOptions)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetRootNode(options)
	return codec.EncodeEntity(cbCtx, result)
}

func Node_cloneNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_isEqualNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_isSameNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_contains[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_insertBefore[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_appendChild[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_replaceChild[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_removeChild[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func Node_nodeType[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NodeType()
	return encodeNodeType(cbCtx, result)
}

func Node_nodeName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return codec.EncodeString(cbCtx, result)
}

func Node_isConnected[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return codec.EncodeBoolean(cbCtx, result)
}

func Node_ownerDocument[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_parentNode[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentNode()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_parentElement[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ParentElement()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_childNodes[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_firstChild[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_lastChild[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LastChild()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_previousSibling[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_nextSibling[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return codec.EncodeEntity(cbCtx, result)
}

func Node_nodeValue[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result, hasValue := instance.NodeValue()
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func Node_setNodeValue[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetNodeValue(val)
	return nil, nil
}

func Node_textContent[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.Node](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.TextContent()
	return codec.EncodeOptionalString(cbCtx, result)
}

func Node_setTextContent[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.Node](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTextContent(val)
	return nil, nil
}
