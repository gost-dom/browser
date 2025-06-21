package fetch

import (
	"io"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w Body[T]) json(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	p := cbCtx.NewPromise()
	res = p
	instance, err := js.As[fetch.Body](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodePromise(cbCtx, func() (js.Value[T], error) {
		b, err := io.ReadAll(instance)
		if err != nil {
			return nil, err
		}
		return cbCtx.JSONParse(string(b))
	})
}
