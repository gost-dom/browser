package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w Request[T]) decodeRequestInfo(
	_ js.Scope[T],
	val js.Value[T],
) (string, error) {
	return val.String(), nil
}

func (w Request[T]) decodeRequestInit(
	s js.Scope[T],
	v js.Value[T],
) ([]fetch.RequestOption, error) {
	return decodeRequestInit(s, v)
}

func (w Request[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	url string,
	options ...fetch.RequestOption,
) (js.Value[T], error) {
	win, err := codec.GetWindow(cbCtx)
	if err != nil {
		return nil, err
	}
	f := fetch.New(win)
	req := f.NewRequest(url, options...)
	return codec.EncodeConstructedValue(cbCtx, &req)
}

func (w Request[T]) toHeaders(cbCtx js.CallbackContext[T], h *fetch.Headers) (js.Value[T], error) {
	return codec.EncodeEntityScopedWithPrototype(cbCtx, h, "Headers")
}
