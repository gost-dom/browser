package v8host

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/v8go"

	v8 "github.com/gost-dom/v8go"
)

func (w formDataV8Wrapper) CustomInitialiser(constructor *v8go.FunctionTemplate) {
	iterator := newIterator2(
		w.scriptHost,
		w.toString_,
		func(ctx jsCallbackContext, v html.FormDataValue) (jsValue, error) {
			return w.toString_(ctx, string(v))
		},
	)
	iterator.installPrototype(constructor)
}

func (w formDataV8Wrapper) CreateInstance(cbCtx *v8CallbackContext) (jsValue, error) {
	value := html.NewFormData()
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) CreateInstanceForm(
	cbCtx *v8CallbackContext,
	form html.HTMLFormElement,
) (jsValue, error) {
	value := html.NewFormDataForm(form)
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) CreateInstanceFormSubmitter(
	cbCtx *v8CallbackContext,
	form html.HTMLFormElement,
	submitter html.HTMLElement,
) (jsValue, error) {
	value := html.NewFormDataForm(form)
	if submitter != nil {
		value.AddElement(submitter)
	}
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) decodeFormDataValue(
	cbCtx jsCallbackContext,
	val jsValue,
) (html.FormDataValue, error) {
	return html.FormDataValue(val.String()), nil
}

func (w formDataV8Wrapper) toFormDataEntryValue(
	cbCtx *v8CallbackContext,
	val html.FormDataValue,
) (jsValue, error) {
	return w.toString_(cbCtx, string(val))
}

func (w formDataV8Wrapper) toSequenceFormDataEntryValue(
	cbCtx *v8CallbackContext,
	data []html.FormDataValue,
) (jsValue, error) {
	vals := make([]*v8.Value, len(data))
	for i, d := range data {
		var err error
		vals[i], err = v8go.NewValue(cbCtx.iso(), string(d))
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
	}
	arr, err := toArray(cbCtx.ScriptCtx().v8ctx, vals...)
	return newV8Value(cbCtx.iso(), arr), err
}
