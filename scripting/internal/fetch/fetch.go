package fetch

import (
	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func Fetch[T any](info js.CallbackContext[T]) (js.Value[T], error) {
	url, err := js.ConsumeArgument(info, "url", nil, codec.DecodeString)
	if err != nil {
		return nil, err
	}
	f := fetch.New(info.Window())
	return codec.EncodePromise(info, func() (js.Value[T], error) {
		r, err := f.Fetch(f.NewRequest(url))
		if err != nil {
			return nil, err
		}

		return info.Constructor("Response").NewInstance(r)
	})
}
