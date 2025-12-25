// This file is generated. Do not edit.

package url

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type URLSearchParams[T any] struct{}

func NewURLSearchParams[T any](scriptHost js.ScriptEngine[T]) *URLSearchParams[T] {
	return &URLSearchParams[T]{}
}

func (wrapper URLSearchParams[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w URLSearchParams[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("append", w.append)
	jsClass.CreateOperation("delete", w.delete)
	jsClass.CreateOperation("get", w.get)
	jsClass.CreateOperation("getAll", w.getAll)
	jsClass.CreateOperation("has", w.has)
	jsClass.CreateOperation("set", w.set)
	jsClass.CreateOperation("sort", w.sort)
	jsClass.CreateOperation("toString", w.toString)
	jsClass.CreatePrototypeAttribute("size", w.size, nil)
}

func (w URLSearchParams[T]) append(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func (w URLSearchParams[T]) delete(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	value, found, errArg := js.ConsumeOptionalArg(cbCtx, "value", codec.DecodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		instance.DeleteValue(name, value)
		return nil, nil
	}
	instance.Delete(name)
	return nil, nil
}

func (w URLSearchParams[T]) get(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result, hasValue := instance.Get(name)
	return codec.EncodeNillableString(cbCtx, result, hasValue)
}

func (w URLSearchParams[T]) getAll(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return w.toSequenceString_(cbCtx, result)
}

func (w URLSearchParams[T]) has(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	value, found, errArg := js.ConsumeOptionalArg(cbCtx, "value", codec.DecodeString)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		result := instance.HasValue(name, value)
		return codec.EncodeBoolean(cbCtx, result)
	}
	result := instance.Has(name)
	return codec.EncodeBoolean(cbCtx, result)
}

func (w URLSearchParams[T]) set(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, codec.DecodeString)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}

func (w URLSearchParams[T]) sort(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Sort()
	return nil, nil
}

func (w URLSearchParams[T]) toString(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.String()
	return codec.EncodeString(cbCtx, result)
}

func (w URLSearchParams[T]) size(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Size()
	return codec.EncodeInt(cbCtx, result)
}
