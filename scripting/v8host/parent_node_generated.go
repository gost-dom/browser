// This file is generated. Do not edit.

package v8host

import (
	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type parentNodeV8Wrapper[T any] struct {
	handleReffedObject[dom.ParentNode, T]
}

func newParentNodeV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *parentNodeV8Wrapper[T] {
	return &parentNodeV8Wrapper[T]{newHandleReffedObject[dom.ParentNode, T](scriptHost)}
}

func (wrapper parentNodeV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w parentNodeV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("prepend", w.prepend)
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("replaceChildren", w.replaceChildren)
	jsClass.CreatePrototypeMethod("querySelector", w.querySelector)
	jsClass.CreatePrototypeMethod("querySelectorAll", w.querySelectorAll)
	jsClass.CreatePrototypeAttribute("firstElementChild", w.firstElementChild, nil)
	jsClass.CreatePrototypeAttribute("lastElementChild", w.lastElementChild, nil)
	jsClass.CreatePrototypeAttribute("childElementCount", w.childElementCount, nil)
}

func (w parentNodeV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w parentNodeV8Wrapper[T]) prepend(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.prepend")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	nodes, errArg1 := consumeRestArguments(cbCtx, "nodes", w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Prepend(nodes...)
	return nil, errCall
}

func (w parentNodeV8Wrapper[T]) append(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.append")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	nodes, errArg1 := consumeRestArguments(cbCtx, "nodes", w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Append(nodes...)
	return nil, errCall
}

func (w parentNodeV8Wrapper[T]) replaceChildren(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.replaceChildren")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	nodes, errArg1 := consumeRestArguments(cbCtx, "nodes", w.decodeNodeOrText)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.ReplaceChildren(nodes...)
	return nil, errCall
}

func (w parentNodeV8Wrapper[T]) querySelector(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.querySelector")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	selectors, errArg1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.QuerySelector(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper[T]) querySelectorAll(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.querySelectorAll")
	instance, errInst := js.As[dom.ParentNode](cbCtx.Instance())
	if errInst != nil {
		return cbCtx.ReturnWithError(errInst)
	}
	selectors, errArg1 := consumeArgument(cbCtx, "selectors", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, errCall := instance.QuerySelectorAll(selectors)
	if errCall != nil {
		return nil, errCall
	}
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper[T]) firstElementChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.firstElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.FirstElementChild()
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper[T]) lastElementChild(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.lastElementChild")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.LastElementChild()
	return encodeEntity(cbCtx, result)
}

func (w parentNodeV8Wrapper[T]) childElementCount(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: ParentNode.childElementCount")
	instance, err := js.As[dom.ParentNode](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.ChildElementCount()
	return w.toUnsignedLong(cbCtx, result)
}
