package xhr

import (
	"github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w FormData[T]) CustomInitializer(class js.Class[T]) {
	iterator := js.NewIterator2(
		codec.EncodeStringScoped,
		func(ctx js.Scope[T], v html.FormDataValue) (js.Value[T], error) {
			return codec.EncodeStringScoped(ctx, string(v))
		},
	)
	iterator.InstallPrototype(class)
}

func (w FormData[T]) CreateInstance(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	value := html.NewFormData()
	return codec.EncodeConstrucedValue(cbCtx, value)
}

func (w FormData[T]) CreateInstanceForm(
	cbCtx js.CallbackContext[T],
	form html.HTMLFormElement,
) (js.Value[T], error) {
	value := html.NewFormDataForm(form)
	return codec.EncodeConstrucedValue(cbCtx, value)
}

func (w FormData[T]) CreateInstanceFormSubmitter(
	cbCtx js.CallbackContext[T],
	form html.HTMLFormElement,
	submitter html.HTMLElement,
) (js.Value[T], error) {
	value := html.NewFormDataForm(form)
	if submitter != nil {
		value.AddElement(submitter)
	}
	return codec.EncodeConstrucedValue(cbCtx, value)
}

func (w FormData[T]) decodeFormDataValue(
	_ js.CallbackContext[T],
	val js.Value[T],
) (html.FormDataValue, error) {
	return html.FormDataValue(val.String()), nil
}

func (w FormData[T]) toFormDataEntryValue(
	cbCtx js.CallbackContext[T],
	val html.FormDataValue,
) (js.Value[T], error) {
	return codec.EncodeString(cbCtx, string(val))
}

func (w FormData[T]) toSequenceFormDataEntryValue(
	cbCtx js.CallbackContext[T],
	data []html.FormDataValue,
) (js.Value[T], error) {
	vals := make([]js.Value[T], len(data))
	for i, d := range data {
		vals[i] = cbCtx.NewString(string(d))
	}
	return cbCtx.NewArray(vals...), nil
}

func (w FormData[T]) decodeHTMLFormElement(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (html.HTMLFormElement, error) {
	var (
		res html.HTMLFormElement
		ok  bool
	)
	node, err := codec.DecodeNode(cbCtx, val)
	if err == nil {
		res, ok = node.(html.HTMLFormElement)
		if !ok {
			err = cbCtx.NewTypeError("Not a form")
		}
	}
	return res, err
}
