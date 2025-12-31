// This file is generated. Do not edit.

package dom

import (
	dom "github.com/gost-dom/browser/dom"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeDOMTokenList[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("item", DOMTokenList_item)
	jsClass.CreateOperation("contains", DOMTokenList_contains)
	jsClass.CreateOperation("add", DOMTokenList_add)
	jsClass.CreateOperation("remove", DOMTokenList_remove)
	jsClass.CreateOperation("toggle", DOMTokenList_toggle)
	jsClass.CreateOperation("replace", DOMTokenList_replace)
	jsClass.CreateOperation("supports", DOMTokenList_supports)
	jsClass.CreateAttribute("length", DOMTokenList_length, nil)
	jsClass.CreateAttribute("value", DOMTokenList_value, DOMTokenList_setValue)
	jsClass.CreateOperation("toString", DOMTokenList_value)
	DOMTokenListCustomInitializer(jsClass)
}

func DOMTokenList_item[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func DOMTokenList_contains[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func DOMTokenList_add[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func DOMTokenList_remove[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func DOMTokenList_replace[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func DOMTokenList_supports[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "DOMTokenList.DOMTokenList_supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func DOMTokenList_length[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return codec.EncodeInt(cbCtx, result)
}

func DOMTokenList_value[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return codec.EncodeString(cbCtx, result)
}

func DOMTokenList_setValue[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err0 := js.As[dom.DOMTokenList](cbCtx.Instance())
	val, err1 := js.ParseSetterArg(cbCtx, codec.DecodeString)
	err = gosterror.First(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
