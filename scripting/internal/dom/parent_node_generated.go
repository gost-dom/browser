// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type ParentNode[T any] struct{}

func NewParentNode[T any](scriptHost js.ScriptEngine[T]) *ParentNode[T] {
	return &ParentNode[T]{}
}

func (wrapper ParentNode[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w ParentNode[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("prepend", w.prepend)
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("replaceChildren", w.replaceChildren)
	jsClass.CreatePrototypeMethod("querySelector", w.querySelector)
	jsClass.CreatePrototypeMethod("querySelectorAll", w.querySelectorAll)
	jsClass.CreatePrototypeAttribute("firstElementChild", w.firstElementChild, nil)
	jsClass.CreatePrototypeAttribute("lastElementChild", w.lastElementChild, nil)
	jsClass.CreatePrototypeAttribute("childElementCount", w.childElementCount, nil)
}

func (w ParentNode[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w ParentNode[T]) prepend(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.prepend")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	nodes, errArg1 := js.ConsumeRestArguments(cbCtx, "nodes", w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Prepend(nodes...)
	return nil, errCall
}

func (w ParentNode[T]) append(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.append")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	nodes, errArg1 := js.ConsumeRestArguments(cbCtx, "nodes", w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Append(nodes...)
	return nil, errCall
}

func (w ParentNode[T]) replaceChildren(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.replaceChildren")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	nodes, errArg1 := js.ConsumeRestArguments(cbCtx, "nodes", w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.ReplaceChildren(nodes...)
	return nil, errCall
}

func (w ParentNode[T]) querySelector(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.querySelector")
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

func (w ParentNode[T]) querySelectorAll(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.querySelectorAll")
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

func (w ParentNode[T]) firstElementChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.firstElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.FirstElementChild()
	return codec.EncodeEntity(cbCtx, result)
}

func (w ParentNode[T]) lastElementChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.lastElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.LastElementChild()
	return codec.EncodeEntity(cbCtx, result)
}

func (w ParentNode[T]) childElementCount(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: ParentNode.childElementCount")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.ChildElementCount()
	return codec.EncodeInt(cbCtx, result)
}
