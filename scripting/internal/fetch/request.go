package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/internal/constants"
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
	_ js.Scope[T],
	v js.Value[T],
) ([]fetch.RequestOption, error) {
	if v != nil {
		return nil, fmt.Errorf(
			"gost-dom/fetch: Request: requestInit not yet supported. %s",
			constants.MISSING_FEATURE_ISSUE_URL,
		)
	}
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
