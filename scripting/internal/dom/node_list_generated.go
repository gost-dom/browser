// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type NodeList[T any] struct{}

func NewNodeList[T any](scriptHost js.ScriptEngine[T]) *NodeList[T] {
	return &NodeList[T]{}
}

func (wrapper NodeList[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w NodeList[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("item", w.item)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
}

func (w NodeList[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: NodeList.Constructor", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w NodeList[T]) item(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: NodeList.item", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, errInst := js.As[dom.NodeList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	index, errArg1 := js.ConsumeArgument(cbCtx, "index", nil, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Item(index)
	return codec.EncodeEntity(cbCtx, result)
}

func (w NodeList[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	defer func() {
		cbCtx.Logger().Debug("JS Function call: NodeList.length", js.ThisLogAttr(cbCtx), js.LogAttr("res", res))
	}()
	instance, err := js.As[dom.NodeList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}
