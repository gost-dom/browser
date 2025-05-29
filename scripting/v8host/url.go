package v8host

import (
	"fmt"
	"iter"
	"runtime/cgo"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/constants"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	"github.com/gost-dom/browser/url"

	v8 "github.com/gost-dom/v8go"
)

type urlV8Wrapper struct {
	handleReffedObject[urlinterfaces.URL]
}

func newURLV8Wrapper(host *V8ScriptHost) urlV8Wrapper {
	return urlV8Wrapper{newHandleReffedObject[urlinterfaces.URL](host)}
}

type handleDisposable cgo.Handle

func (h handleDisposable) dispose() { cgo.Handle(h).Delete() }

func (w urlV8Wrapper) CreateInstance(
	cbCtx *v8CallbackContext,
	u string,
) (jsValue, error) {
	value, err := url.NewUrl(u)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w urlV8Wrapper) CreateInstanceBase(
	cbCtx *v8CallbackContext,
	u string,
	base string,
) (jsValue, error) {
	value, err := url.NewUrlBase(u, base)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w urlSearchParamsV8Wrapper) Constructor(cbCtx *v8CallbackContext) (jsValue, error) {
	var err error
	ctx := cbCtx.ScriptCtx()
	args := cbCtx.consumeRest()
	var res url.URLSearchParams
	if len(args) > 0 {
		arg := args[0]

		switch {
		case arg.IsString():
			res, err = url.ParseURLSearchParams(arg.String())
			if err != nil {
				return cbCtx.ReturnWithError(err)
			}
		case arg.IsObject():
			if gv, err2 := v8ValueToGoValue(arg); err2 == nil {
				if fd, ok := gv.(*html.FormData); ok {
					res = url.URLSearchParams{}
					for _, pair := range fd.Entries {
						res.Append(pair.Name, string(pair.Value))

					}
					break
				}
			}
			o, _ := arg.AsObject()
			if k, err := ctx.v8ctx.RunScript("Object.keys", ""); err == nil {
				if f, err := k.AsFunction(); err == nil {
					if keys, err := f.Call(ctx.v8ctx.Global(), arg); err == nil {
						if goKeys, err := v8ValueToGoValue(keys); err == nil {
							if arr, ok := goKeys.([]any); ok {
								res = url.URLSearchParams{}
								for _, key := range arr {
									if strKey, isString := key.(string); isString {
										if val, err := o.Get(strKey); err == nil {
											res.Append(strKey, val.String())
										}
									}
								}
								break
							}
						}
					}
				}
			}

			fallthrough
		default:
			return cbCtx.ReturnWithError(
				fmt.Errorf(
					"URLSearchParams: unsupported argument. If the argument is _valid_: %s",
					constants.BUG_ISSUE_URL,
				))
		}
	}
	w.store(&res, ctx, cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w urlSearchParamsV8Wrapper) toSequenceString_(
	cbCtx *v8CallbackContext,
	values []string,
) (jsValue, error) {
	vs := make([]*v8.Value, len(values))
	for i, v := range values {
		vs[i], _ = v8.NewValue(cbCtx.iso(), v)
	}
	arr, err := toArray(cbCtx.ScriptCtx().v8ctx, vs...)
	return newV8Value(nil, arr), err
}

func (w urlSearchParamsV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	it := newIterator2(w.scriptHost, w.toString_, w.toString_)
	it.installPrototype(constructor)
}

type Keys[K, V any] struct {
	iter iterable2[K, V]
}

func (k Keys[K, V]) All() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range k.iter.All() {
			if !yield(k) {
				return
			}
		}
	}
}

type iterValues[K, V any] struct {
	iter iterable2[K, V]
}

func (k iterValues[K, V]) All() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range k.iter.All() {
			if !yield(v) {
				return
			}
		}
	}
}
