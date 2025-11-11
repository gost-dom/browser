package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w Headers[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	options ...fetch.HeaderOption,
) (js.Value[T], error) {
	return codec.EncodeConstrucedValue(cbCtx, fetch.Headers{})
}

func (w Headers[T]) decodeHeadersInit(_ js.Scope[T], v js.Value[T]) ([]fetch.HeaderOption, error) {
	return nil, nil
}
