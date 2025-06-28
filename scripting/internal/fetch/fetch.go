package fetch

import (
	"errors"

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
	signal, ok, err1 := decodeOption(ctx, val, "signal", fetch.WithSignal)
	if ok && err1 == nil {
		opts = append(opts, signal)
	}
	err = errors.Join(err1)
	return
}

type Decoder[T, U any] = func(scope js.Scope[T], v js.Value[T]) (U, error)

func decodeOption[T, U, V any](
	_ js.Scope[T],
	v js.Value[T],
	key string,
	encode func(U) V,
) (res V, ok bool, err error) {
	var obj js.Object[T]
	if obj, ok = v.AsObject(); !ok {
		return
	}
	var opt js.Value[T]
	if opt, err = obj.Get(key); err != nil {
		return
	}
	optObj, ok := opt.AsObject()
	if ok {
		native, succ := optObj.NativeValue().(U)
		if succ {
			res = encode(native)
		} else {
			err = errors.New("Option is not correct type")
		}
	}
	return
}
