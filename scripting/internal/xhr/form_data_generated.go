// This file is generated. Do not edit.

package xhr

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type FormData[T any] struct{}

func NewFormData[T any](scriptHost js.ScriptEngine[T]) *FormData[T] {
	return &FormData[T]{}
}

func (wrapper FormData[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
	wrapper.CustomInitializer(jsClass)
}

func (w FormData[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("append", w.append)
	jsClass.CreatePrototypeMethod("delete", w.delete)
	jsClass.CreatePrototypeMethod("get", w.get)
	jsClass.CreatePrototypeMethod("getAll", w.getAll)
	jsClass.CreatePrototypeMethod("has", w.has)
	jsClass.CreatePrototypeMethod("set", w.set)
}

func (w FormData[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: FormData.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: FormData.Constructor", js.LogAttr("res", res))
	}()
	form, found, errArg := js.ConsumeOptionalArg(cbCtx, "form", w.decodeHTMLFormElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceForm(cbCtx, form)
	}
	submitter, found, errArg := js.ConsumeOptionalArg(cbCtx, "submitter", codec.DecodeHTMLElement)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceFormSubmitter(cbCtx, form, submitter)
	}
	return w.CreateInstance(cbCtx)
}

func (w FormData[T]) append(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: FormData.append - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: FormData.append", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func (w FormData[T]) delete(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: FormData.delete - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: FormData.delete", js.LogAttr("res", res))
	}()
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

func (w FormData[T]) get(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: FormData.get - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: FormData.get", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.Get(name)
	return w.toFormDataEntryValue(cbCtx, result)
}

func (w FormData[T]) getAll(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: FormData.getAll - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: FormData.getAll", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	result := instance.GetAll(name)
	return w.toSequenceFormDataEntryValue(cbCtx, result)
}

func (w FormData[T]) has(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: FormData.has - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: FormData.has", js.LogAttr("res", res))
	}()
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

func (w FormData[T]) set(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: FormData.set - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: FormData.set", js.LogAttr("res", res))
	}()
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}
