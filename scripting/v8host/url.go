package v8host

import (
	"fmt"
	"runtime/cgo"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/constants"
	urlinterfaces "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	"github.com/gost-dom/browser/url"
)

type urlV8Wrapper struct {
	handleReffedObject[urlinterfaces.URL, jsTypeParam]
}

func newURLV8Wrapper(host *V8ScriptHost) urlV8Wrapper {
	return urlV8Wrapper{newHandleReffedObject[urlinterfaces.URL](host)}
}

type handleDisposable cgo.Handle

func (h handleDisposable) Dispose() { cgo.Handle(h).Delete() }

func (w urlV8Wrapper) CreateInstance(
	cbCtx jsCallbackContext,
	u string,
) (jsValue, error) {
	value, err := url.NewUrl(u)
	if err != nil {
		return nil, err
	}
	w.store(value, cbCtx)
	return nil, nil
}

func (w urlV8Wrapper) CreateInstanceBase(
	cbCtx jsCallbackContext,
	u string,
	base string,
) (jsValue, error) {
	value, err := url.NewUrlBase(u, base)
	if err != nil {
		return nil, err
	}
	w.store(value, cbCtx)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
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
	w.store(&res, cbCtx)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) toSequenceString_(
	cbCtx jsCallbackContext,
	values []string,
) (jsValue, error) {
	vs := make([]jsValue, len(values))
	fact := cbCtx.ValueFactory()
	for i, v := range values {
		vs[i] = fact.NewString(v)
	}
	return fact.NewArray(vs...), nil
}

func (w urlSearchParamsV8Wrapper) CustomInitializer(class jsClass) {
	it := newIterator2(w.scriptHost, w.toString_, w.toString_)
	it.installPrototype(class.(v8Class).ft)
}
