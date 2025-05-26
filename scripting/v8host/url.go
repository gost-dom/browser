package v8host

import (
	"errors"
	"fmt"
	"iter"
	"runtime/cgo"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/constants"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	log "github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
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
	cbCtx *argumentHelper,
	u string,
) js.CallbackRVal {
	value, err := url.NewUrl(u)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w urlV8Wrapper) CreateInstanceBase(
	cbCtx *argumentHelper,
	u string,
	base string,
) js.CallbackRVal {
	value, err := url.NewUrlBase(u, base)
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	w.store(value, cbCtx.ScriptCtx(), cbCtx.This())
	return cbCtx.ReturnWithValue(nil)
}

func (w urlSearchParamsV8Wrapper) Constructor(cbCtx *argumentHelper) js.CallbackRVal {
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
	cbCtx *argumentHelper,
	// ctx *V8ScriptContext,
	values []string,
) js.CallbackRVal {
	vs := make([]*v8.Value, len(values))
	for i, v := range values {
		vs[i], _ = v8.NewValue(cbCtx.iso(), v)
	}
	return cbCtx.ReturnWithValueErr(toArray(cbCtx.ScriptCtx().v8ctx, vs...))
}

func (w urlSearchParamsV8Wrapper) CustomInitialiser(constructor *v8.FunctionTemplate) {
	iso := w.scriptHost.iso
	it := newIterator2(
		w.scriptHost,
		func(k string, v string, ctx *V8ScriptContext) (*v8.Value, *v8.Value, error) {
			log.Info(w.scriptHost.logger, "Iterate", "key", k, "value", v)
			r1, e1 := v8.NewValue(iso, k)
			r2, e2 := v8.NewValue(iso, v)
			return r1, r2, errors.Join(e1, e2)
		},
	)
	fmt.Println("Install iterator")
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
