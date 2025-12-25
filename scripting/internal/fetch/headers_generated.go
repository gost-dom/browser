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
	jsClass.CreateOperation("append", w.append)
	jsClass.CreateOperation("delete", w.delete)
	jsClass.CreateOperation("get", w.get)
	jsClass.CreateOperation("getSetCookie", w.getSetCookie)
	jsClass.CreateOperation("has", w.has)
	jsClass.CreateOperation("set", w.set)
}

func (w Headers[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	init, errArg1 := js.ConsumeArgument(cbCtx, "init", nil, w.decodeHeadersInit)
	if errArg1 != nil {
		return nil, errArg1
	}
	return w.CreateInstance(cbCtx, init...)
}

func (w Headers[T]) append(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func (w Headers[T]) delete(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func (w Headers[T]) get(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func (w Headers[T]) getSetCookie(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Headers.getSetCookie: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Headers[T]) has(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func (w Headers[T]) set(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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
