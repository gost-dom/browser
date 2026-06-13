package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func ResponseConstructor[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return nil, fmt.Errorf("gost-dom/fetch: Response constructor not implemented")
}

// Response_ok implements the Response.ok getter (status in the range 200-299).
func Response_ok[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[*fetch.Response](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeBoolean(cbCtx, instance.Ok())
}

// Response_statusText implements the Response.statusText getter.
func Response_statusText[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[*fetch.Response](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeString(cbCtx, instance.StatusText())
}
