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

func Response_ok[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Response](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeBoolean(cbCtx, instance.Ok())
}

func Response_statusText[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[*fetch.Response](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeString(cbCtx, instance.StatusText())
}
