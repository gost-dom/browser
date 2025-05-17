package v8host

import (
	"fmt"
	"runtime/cgo"

	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/url"

	v8 "github.com/gost-dom/v8go"
)

type urlV8Wrapper struct {
	handleReffedObject[*url.URL]
}

func newURLV8Wrapper(host *V8ScriptHost) urlV8Wrapper {
	return urlV8Wrapper{newHandleReffedObject[*url.URL](host)}
}

type handleDisposable cgo.Handle

func (h handleDisposable) dispose() { cgo.Handle(h).Delete() }

func (w urlV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
	u string,
) (*v8.Value, error) {
	value, err := url.NewUrl(u)
	if err != nil {
		return nil, err
	}
	w.store(value, ctx, this)
	return nil, nil
}

func (w urlV8Wrapper) CreateInstanceBase(
	ctx *V8ScriptContext,
	this *v8.Object,
	u string,
	base string,
) (*v8.Value, error) {
	value, err := url.NewUrlBase(u, base)
	if err != nil {
		return nil, err
	}
	w.store(value, ctx, this)
	return nil, nil
}

func (w urlSearchParamsV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := info.Args()
	var res url.URLSearchParams
	if len(args) > 0 {
		arg := args[0]
		if !arg.IsString() {
			return nil, fmt.Errorf(
				"URLSearchParams: Constructor only supports no arguments, or a string. If the argument is _valid_: %s",
				constants.BUG_USSUE_URL,
			)
		}
		var err error
		res, err = url.ParseURLSearchParams(arg.String())
		if err != nil {
			return nil, err
		}
	}
	ctx := w.mustGetContext(info)
	w.store(&res, ctx, info.This())
	return nil, nil
}
