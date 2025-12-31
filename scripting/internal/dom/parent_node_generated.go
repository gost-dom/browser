// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeParentNode[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("prepend", ParentNode_prepend)
	jsClass.CreateOperation("append", ParentNode_append)
	jsClass.CreateOperation("replaceChildren", ParentNode_replaceChildren)
	jsClass.CreateOperation("querySelector", ParentNode_querySelector)
	jsClass.CreateOperation("querySelectorAll", ParentNode_querySelectorAll)
	jsClass.CreateAttribute("children", ParentNode_children, nil)
	jsClass.CreateAttribute("firstElementChild", ParentNode_firstElementChild, nil)
	jsClass.CreateAttribute("lastElementChild", ParentNode_lastElementChild, nil)
	jsClass.CreateAttribute("childElementCount", ParentNode_childElementCount, nil)
}

func ParentNode_prepend[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	nodes, errArg1 := js.ConsumeRestArguments(cbCtx, "nodes", decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Prepend(nodes...)
	return nil, errCall
}

func ParentNode_append[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	nodes, errArg1 := js.ConsumeRestArguments(cbCtx, "nodes", decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Append(nodes...)
	return nil, errCall
}

func ParentNode_replaceChildren[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	nodes, errArg1 := js.ConsumeRestArguments(cbCtx, "nodes", decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.ReplaceChildren(nodes...)
	return nil, errCall
}

func ParentNode_querySelector[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	selectors, errArg1 := js.ConsumeArgument(cbCtx, "selectors", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.QuerySelector(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func ParentNode_querySelectorAll[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	selectors, errArg1 := js.ConsumeArgument(cbCtx, "selectors", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.QuerySelectorAll(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeEntity(cbCtx, result)
}

func ParentNode_children[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Children()
	return encodeHTMLCollection(cbCtx, result)
}

func ParentNode_firstElementChild[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstElementChild()
	return codec.EncodeEntity(cbCtx, result)
}

func ParentNode_lastElementChild[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LastElementChild()
	return codec.EncodeEntity(cbCtx, result)
}

func ParentNode_childElementCount[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildElementCount()
	return codec.EncodeInt(cbCtx, result)
}
