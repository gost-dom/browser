package v8host

import (
	"github.com/gost-dom/browser/html"
)

func (w formDataV8Wrapper) CustomInitializer(class jsClass) {
	iterator := newIterator2(
		w.toString_,
		func(ctx jsCallbackContext, v html.FormDataValue) (jsValue, error) {
			return w.toString_(ctx, string(v))
		},
	)
	iterator.installPrototype(class)
}

func (w formDataV8Wrapper) CreateInstance(cbCtx jsCallbackContext) (jsValue, error) {
	value := html.NewFormData()
	return w.store(value, cbCtx)
}

func (w formDataV8Wrapper) CreateInstanceForm(
	cbCtx jsCallbackContext,
	form html.HTMLFormElement,
) (jsValue, error) {
	value := html.NewFormDataForm(form)
	return w.store(value, cbCtx)
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
	return w.store(value, cbCtx)
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
