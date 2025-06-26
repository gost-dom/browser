package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func encodeResponse[T any](info js.Scope[T], res *fetch.Response) (js.Value[T], error) {
	return info.Constructor("Response").NewInstance(res)
}

func Fetch[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	url, err := js.ConsumeArgument(info, "url", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	f := fetch.New(info.Window())
	info.Logger().Debug("js/fetch: create promise")
	req := f.NewRequest(url)
	return codec.EncodePromise(info, f.FetchAsync(req), encodeResponse)
}
