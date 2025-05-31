package v8host

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/v8go"
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

func (w formDataV8Wrapper) CreateInstance(cbCtx jsCallbackContext) (jsValue, error) {
	value := html.NewFormData()
	w.store(value, cbCtx)
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) CreateInstanceForm(
	cbCtx jsCallbackContext,
	form html.HTMLFormElement,
) (jsValue, error) {
	value := html.NewFormDataForm(form)
	w.store(value, cbCtx)
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) CreateInstanceFormSubmitter(
	cbCtx jsCallbackContext,
	form html.HTMLFormElement,
	submitter html.HTMLElement,
) (jsValue, error) {
	value := html.NewFormDataForm(form)
	if submitter != nil {
		value.AddElement(submitter)
	}
	w.store(value, cbCtx)
	return cbCtx.ReturnWithValue(nil)
}

func (w formDataV8Wrapper) decodeFormDataValue(
	cbCtx jsCallbackContext,
	val jsValue,
) (html.FormDataValue, error) {
	return html.FormDataValue(val.String()), nil
}

func (w formDataV8Wrapper) toFormDataEntryValue(
	cbCtx jsCallbackContext,
	val html.FormDataValue,
) (jsValue, error) {
	return w.toString_(cbCtx, string(val))
}

func (w formDataV8Wrapper) toSequenceFormDataEntryValue(
	cbCtx jsCallbackContext,
	data []html.FormDataValue,
) (jsValue, error) {
	vals := make([]jsValue, len(data))
	for i, d := range data {
		vals[i] = cbCtx.ValueFactory().NewString(string(d))
	}
	return cbCtx.ValueFactory().NewArray(vals...), nil
}
