package v8host

import (
	"fmt"
	"runtime/cgo"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/constants"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/url"
)

type urlV8Wrapper[T any] struct {
	handleReffedObject[urlinterfaces.URL, T]
}

func newURLV8Wrapper[T any](host js.ScriptEngine[T]) urlV8Wrapper[T] {
	return urlV8Wrapper[T]{newHandleReffedObject[urlinterfaces.URL](host)}
}

type handleDisposable cgo.Handle

func (h handleDisposable) Dispose() { cgo.Handle(h).Delete() }

func (w urlV8Wrapper[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	u string,
) (js.Value[T], error) {
	value, err := url.NewUrl(u)
	if err != nil {
		return nil, err
	}
	storeNewValue(value, cbCtx)
	return nil, nil
}

func (w urlV8Wrapper[T]) CreateInstanceBase(
	cbCtx js.CallbackContext[T],
	u string,
	base string,
) (js.Value[T], error) {
	value, err := url.NewUrlBase(u, base)
	if err != nil {
		return nil, err
	}
	storeNewValue(value, cbCtx)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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
	storeNewValue(&res, cbCtx)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper[T]) toSequenceString_(
	cbCtx js.CallbackContext[T],
	values []string,
) (js.Value[T], error) {
	vs := make([]js.Value[T], len(values))
	fact := cbCtx.ValueFactory()
	for i, v := range values {
		vs[i] = fact.NewString(v)
	}
	return fact.NewArray(vs...), nil
}

func (w urlSearchParamsV8Wrapper[T]) CustomInitializer(class js.Class[T]) {
	it := newIterator2(w.toString_, w.toString_)
	it.installPrototype(class)
}
