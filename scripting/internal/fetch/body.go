package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/internal/promise"
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
	cbCtx.Logger().Debug("js/Body.json")
	return codec.EncodePromise(cbCtx, promise.ReadAll(instance), EncodeJSONBytes)
}

func EncodeJSONBytes[T any](scope js.Scope[T], b []byte) (js.Value[T], error) {
	return scope.JSONParse(string(b))
}
