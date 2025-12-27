package url

import (
	"fmt"
	"runtime/cgo"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/constants"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/url"
)

type URL[T any] struct {
}

func NewURL[T any](host js.ScriptEngine[T]) URL[T] {
	return URL[T]{}
}

type handleDisposable cgo.Handle

func (h handleDisposable) Dispose() { cgo.Handle(h).Delete() }

func (w URL[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	u string,
) (js.Value[T], error) {
	value, err := url.NewUrl(u)
	if err != nil {
		return nil, err
	}
	return codec.EncodeConstrucedValue(cbCtx, value)
}

func (w URL[T]) CreateInstanceBase(
	cbCtx js.CallbackContext[T],
	u string,
	base string,
) (js.Value[T], error) {
	value, err := url.NewUrlBase(u, base)
	if err != nil {
		return nil, err
	}
	return codec.EncodeConstrucedValue(cbCtx, value)
}

func (w URLSearchParams[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	var err error
	arg, ok := cbCtx.ConsumeArg()
	var res url.URLSearchParams
	if ok {
		obj, isObj := arg.AsObject()
		switch {
		case arg.IsString():
			res, err = url.ParseURLSearchParams(arg.String())
			if err != nil {
				return nil, err
			}
		case isObj:
			if fd, ok := obj.NativeValue().(*html.FormData); ok {
				res = url.URLSearchParams{}
				for _, pair := range fd.Entries {
					res.Append(pair.Name, string(pair.Value))
				}
				break
			}
			if keys, err := obj.Keys(); err == nil {
				for _, key := range keys {
					if val, err := obj.Get(key); err == nil {
						res.Append(key, val.String())
					}
				}
				break
			}
			fallthrough
		default:
			return nil, fmt.Errorf(
				"URLSearchParams: unsupported argument. If the argument is _valid_: %s",
				constants.BUG_ISSUE_URL,
			)
		}
	}
	return codec.EncodeConstrucedValue(cbCtx, &res)
}

func encodeSequenceString_[T any](
	cbCtx js.CallbackContext[T],
	values []string,
) (js.Value[T], error) {
	vs := make([]js.Value[T], len(values))
	for i, v := range values {
		vs[i] = cbCtx.NewString(v)
	}
	return cbCtx.NewArray(vs...), nil
}

func (w URLSearchParams[T]) CustomInitializer(class js.Class[T]) {
	it := js.NewIterator2(codec.EncodeString[T], codec.EncodeString[T])
	it.InstallPrototype(class)
}
