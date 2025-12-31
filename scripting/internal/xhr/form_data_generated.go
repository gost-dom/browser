// This file is generated. Do not edit.

package xhr

import (
	html "github.com/gost-dom/browser/html"
	gosterror "github.com/gost-dom/browser/internal/gosterror"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type FormData[T any] struct{}

func NewFormData[T any](scriptHost js.ScriptEngine[T]) FormData[T] {
	return FormData[T]{}
}

func (wrapper FormData[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w FormData[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("append", FormData_append)
	jsClass.CreateOperation("delete", FormData_delete)
	jsClass.CreateOperation("get", FormData_get)
	jsClass.CreateOperation("getAll", FormData_getAll)
	jsClass.CreateOperation("has", FormData_has)
	jsClass.CreateOperation("set", FormData_set)
}

func FormDataConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	form, found, errArg := js.ConsumeOptionalArg(cbCtx, "form", decodeHTMLFormElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return CreateFormDataForm(cbCtx, form)
	}
	submitter, found, errArg := js.ConsumeOptionalArg(cbCtx, "submitter", codec.DecodeHTMLElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return CreateFormDataFormSubmitter(cbCtx, form, submitter)
	}
	return CreateFormData(cbCtx)
}

func FormData_append[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, decodeFormDataValue)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func FormData_delete[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	instance.Delete(name)
	return nil, nil
}

func FormData_get[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Get(name)
	return encodeFormDataEntryValue(cbCtx, result)
}

func FormData_getAll[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return encodeSequenceFormDataEntryValue(cbCtx, result)
}

func FormData_has[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Has(name)
	return codec.EncodeBoolean(cbCtx, result)
}

func FormData_set[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, decodeFormDataValue)
	err = gosterror.First(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}
