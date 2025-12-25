// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type DOMTokenList[T any] struct{}

func NewDOMTokenList[T any](scriptHost js.ScriptEngine[T]) *DOMTokenList[T] {
	return &DOMTokenList[T]{}
}

func (wrapper DOMTokenList[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w DOMTokenList[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("item", w.item)
	jsClass.CreateOperation("contains", w.contains)
	jsClass.CreateOperation("add", w.add)
	jsClass.CreateOperation("remove", w.remove)
	jsClass.CreateOperation("toggle", w.toggle)
	jsClass.CreateOperation("replace", w.replace)
	jsClass.CreateOperation("supports", w.supports)
	jsClass.CreateAttribute("length", w.length, nil)
	jsClass.CreateAttribute("value", w.value, w.setValue)
	jsClass.CreateOperation("toString", w.value)
}

func (w DOMTokenList[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w DOMTokenList[T]) item(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	index, errArg1 := js.ConsumeArgument(cbCtx, "index", nil, codec.DecodeInt)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Item(index)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w DOMTokenList[T]) contains(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	token, errArg1 := js.ConsumeArgument(cbCtx, "token", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Contains(token)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w DOMTokenList[T]) add(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	tokens, errArg1 := js.ConsumeRestArguments(cbCtx, "tokens", codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Add(tokens...)
	return nil, errCall
}

func (w DOMTokenList[T]) remove(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	tokens, errArg1 := js.ConsumeRestArguments(cbCtx, "tokens", codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	errCall := instance.Remove(tokens...)
	return nil, errCall
}

func (w DOMTokenList[T]) replace(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[dom.DOMTokenList](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	token, errArg1 := js.ConsumeArgument(cbCtx, "token", nil, codec.DecodeString)
	newToken, errArg2 := js.ConsumeArgument(cbCtx, "newToken", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	result, errCall := instance.Replace(token, newToken)
	if errCall != nil {
		return nil, errCall
	}
	return codec.EncodeBoolean(cbCtx, result)
}

func (w DOMTokenList[T]) supports(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "DOMTokenList.supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w DOMTokenList[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}

func (w DOMTokenList[T]) value(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func (w DOMTokenList[T]) setValue(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
