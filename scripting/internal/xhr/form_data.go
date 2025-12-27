package xhr

import (
	"github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w FormData[T]) CustomInitializer(class js.Class[T]) {
	iterator := js.NewIterator2(
		codec.EncodeString,
		func(s js.Scope[T], v html.FormDataValue) (js.Value[T], error) {
			return codec.EncodeString(s, string(v))
		},
	)
	iterator.InstallPrototype(class)
}

func CreateFormData[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	value := html.NewFormData()
	return codec.EncodeConstructedValue(cbCtx, value)
}

func CreateFormDataForm[T any](
	cbCtx js.CallbackContext[T],
	form html.HTMLFormElement,
) (js.Value[T], error) {
	value := html.NewFormDataForm(form)
	return codec.EncodeConstructedValue(cbCtx, value)
}

func CreateFormDataFormSubmitter[T any](
	cbCtx js.CallbackContext[T],
	form html.HTMLFormElement,
	submitter html.HTMLElement,
) (js.Value[T], error) {
	value := html.NewFormDataForm(form)
	if submitter != nil {
		value.AddElement(submitter)
	}
	return codec.EncodeConstructedValue(cbCtx, value)
}

func decodeFormDataValue[T any](
	_ js.Scope[T],
	val js.Value[T],
) (html.FormDataValue, error) {
	return html.FormDataValue(val.String()), nil
}

func encodeFormDataEntryValue[T any](
	cbCtx js.CallbackContext[T],
	val html.FormDataValue,
) (js.Value[T], error) {
	return codec.EncodeString(cbCtx, string(val))
}

func encodeSequenceFormDataEntryValue[T any](
	cbCtx js.CallbackContext[T],
	data []html.FormDataValue,
) (js.Value[T], error) {
	vals := make([]js.Value[T], len(data))
	for i, d := range data {
		vals[i] = cbCtx.NewString(string(d))
	}
	return cbCtx.NewArray(vals...), nil
}

func decodeHTMLFormElement[T any](
	s js.Scope[T],
	val js.Value[T],
) (html.HTMLFormElement, error) {
	var (
		res html.HTMLFormElement
		ok  bool
	)
	node, err := codec.DecodeNode(s, val)
	if err == nil {
		res, ok = node.(html.HTMLFormElement)
		if !ok {
			err = s.NewTypeError("Not a form")
		}
	}
	return res, err
}
