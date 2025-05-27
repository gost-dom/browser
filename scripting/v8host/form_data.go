package v8host

import (
	"errors"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/v8go"

	v8 "github.com/gost-dom/v8go"
)

func (w formDataV8Wrapper) CustomInitialiser(constructor *v8go.FunctionTemplate) {
	iso := w.scriptHost.iso
	iterator := newIterator2(
		w.scriptHost,
		func(k string, v html.FormDataValue, ctx *V8ScriptContext) (v1 *v8.Value, v2 *v8.Value, err error) {
			var err1, err2 error
			v1, err1 = v8go.NewValue(iso, k)
			v2, err2 = v8go.NewValue(iso, string(v))
			err = errors.Join(err1, err2)
			return
		},
	)
	iterator.installPrototype(constructor)
}

func (w formDataV8Wrapper) CreateInstance(cbCtx *argumentHelper) (jsValue, error) {
	value := html.NewFormData()
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) CreateInstanceForm(
	cbCtx *argumentHelper,
	form html.HTMLFormElement,
) (jsValue, error) {
	value := html.NewFormDataForm(form)
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) CreateInstanceFormSubmitter(
	cbCtx *argumentHelper,
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
	cbCtx *argumentHelper,
	val html.FormDataValue,
) (jsValue, error) {
	return w.toString_(cbCtx, string(val))
}

func (w formDataV8Wrapper) toSequenceFormDataEntryValue(
	cbCtx *argumentHelper,
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
	return cbCtx.ReturnWithValueErr(toArray(cbCtx.ScriptCtx().v8ctx, vals...))
}
