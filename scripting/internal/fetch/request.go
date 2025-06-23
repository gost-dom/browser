package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w Request[T]) decodeRequestInfo(
	_ js.CallbackContext[T],
	val js.Value[T],
) (string, error) {
	return val.String(), nil
}

func (w Request[T]) decodeRequestInit(
	_ js.CallbackContext[T],
	_ js.Value[T],
) ([]fetch.RequestOption, error) {
	return nil, nil
}

func (w Request[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	url string,
	options ...fetch.RequestOption,
) (js.Value[T], error) {
	f := fetch.New(cbCtx.Window())
	req := f.NewRequest(url)
	return codec.EncodeConstrucedValue(cbCtx, &req)
}
