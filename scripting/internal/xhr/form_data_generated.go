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

func (w FormData[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: FormData.Constructor")
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

func (w FormData[T]) append(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: FormData.append")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Append(name, value)
	return nil, nil
}

func (w FormData[T]) delete(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: FormData.delete")
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

func (w FormData[T]) get(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: FormData.get")
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

func (w FormData[T]) getAll(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: FormData.getAll")
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

func (w FormData[T]) has(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: FormData.has")
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

func (w FormData[T]) set(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: FormData.set")
	instance, errInst := js.As[*html.FormData](cbCtx.Instance())
	if errInst != nil {
		return nil, errInst
	}
	name, errArg1 := js.ConsumeArgument(cbCtx, "name", nil, codec.DecodeString)
	value, errArg2 := js.ConsumeArgument(cbCtx, "value", nil, w.decodeFormDataValue)
	err := errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	instance.Set(name, value)
	return nil, nil
}
