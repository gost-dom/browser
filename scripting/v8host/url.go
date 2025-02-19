package v8host

import (
	"runtime/cgo"

	"github.com/gost-dom/browser/url"

	v8 "github.com/tommie/v8go"
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
