// This file is generated. Do not edit.

package url

import (
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeURLSearchParams[T any](jsClass js.Class[T]) {
	jsClass.CreateOperation("append", URLSearchParams_append)
	jsClass.CreateOperation("delete", URLSearchParams_delete)
	jsClass.CreateOperation("get", URLSearchParams_get)
	jsClass.CreateOperation("getAll", URLSearchParams_getAll)
	jsClass.CreateOperation("has", URLSearchParams_has)
	jsClass.CreateOperation("set", URLSearchParams_set)
	jsClass.CreateOperation("sort", URLSearchParams_sort)
	jsClass.CreateOperation("toString", URLSearchParams_toString)
	jsClass.CreateAttribute("size", URLSearchParams_size, nil)
	URLSearchParamsCustomInitializer(jsClass)
}

func URLSearchParams_append[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func URLSearchParams_delete[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func URLSearchParams_get[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func URLSearchParams_getAll[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return encodeSequenceString_(cbCtx, result)
}

func URLSearchParams_has[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func URLSearchParams_set[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
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

func URLSearchParams_sort[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Sort()
	return nil, nil
}

func URLSearchParams_toString[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.String()
	return codec.EncodeString(cbCtx, result)
}

func URLSearchParams_size[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[urlinterfaces.URLSearchParams](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Size()
	return codec.EncodeInt(cbCtx, result)
}
