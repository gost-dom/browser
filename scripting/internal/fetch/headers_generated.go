// This file is generated. Do not edit.

package fetch

import (
	fetch "github.com/gost-dom/browser/internal/fetch"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Headers[T any] struct{}

func NewHeaders[T any](scriptHost js.ScriptEngine[T]) *Headers[T] {
	return &Headers[T]{}
}

func (wrapper Headers[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w Headers[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("append", Headers_append)
	jsClass.CreateOperation("delete", Headers_delete)
	jsClass.CreateOperation("get", Headers_get)
	jsClass.CreateOperation("getSetCookie", Headers_getSetCookie)
	jsClass.CreateOperation("has", Headers_has)
	jsClass.CreateOperation("set", Headers_set)
}

func HeadersConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	init, errArg1 := js.ConsumeArgument(cbCtx, "init", nil, decodeHeadersInit)
	if errArg1 != nil {
		return nil, errArg1
	}
	return CreateHeaders(cbCtx, init...)
}

func Headers_append[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*fetch.Headers](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeByteString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeByteString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func Headers_delete[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*fetch.Headers](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeByteString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Delete(name)
	return nil, nil
}

func Headers_get[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*fetch.Headers](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeByteString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Get(name)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func Headers_getSetCookie[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Headers.Headers_getSetCookie: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Headers_has[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*fetch.Headers](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeByteString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Has(name)
	return codec.EncodeBoolean(cbCtx, result)
}

func Headers_set[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*fetch.Headers](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeByteString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeByteString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}
