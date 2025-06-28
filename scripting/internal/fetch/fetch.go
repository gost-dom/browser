package fetch

import (
	"github.com/gost-dom/browser/internal/dom"
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
	opts, err := js.ConsumeArgument(info, "options", defaultRequestOptions, decodeRequestOptions)
	f := fetch.New(info.Window())
	info.Logger().Debug("js/fetch: create promise")
	req := f.NewRequest(url, opts...)
	return codec.EncodePromise(info, f.FetchAsync(req), encodeResponse)
}

func defaultRequestOptions() []fetch.RequestOption { return nil }

func decodeRequestOptions[T any](
	ctx js.CallbackContext[T],
	val js.Value[T],
) (opts []fetch.RequestOption, err error) {
	obj, ok := val.AsObject()
	if !ok {
		return nil, nil
	}
	var f js.Value[T]
	if f, err = obj.Get("signal"); err != nil {
		return
	}
	fobj, ok := f.AsObject()
	signal, err := js.As[*dom.AbortSignal](fobj.NativeValue(), nil)
	if err != nil {
		return nil, err
	}
	if signal != nil {
		ctx.Logger().Debug("ADD Signal", "signal", signal)
		opts = append(opts, fetch.WithSignal(signal))
	}
	return
}
