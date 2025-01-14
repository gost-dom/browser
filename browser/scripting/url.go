package scripting

import (
	"runtime/cgo"

	"github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

type urlV8Wrapper struct {
	handleReffedObject[dom.URL]
}

func newUrlV8Wrapper(host *ScriptHost) urlV8Wrapper {
	return urlV8Wrapper{NewHandleReffedObject[dom.URL](host)}
}

type HandleDisposable cgo.Handle

func (h HandleDisposable) Dispose() { cgo.Handle(h).Delete() }

func (u urlV8Wrapper) CreateInstance(
	ctx *ScriptContext,
	this *v8.Object,
	url string,
) (*v8.Value, error) {
	value, err := dom.NewUrl(url)
	if err != nil {
		return nil, err
	}
	u.Store(value, ctx, this)
	return nil, nil
}

func (u urlV8Wrapper) CreateInstanceBase(
	ctx *ScriptContext,
	this *v8.Object,
	url string,
	base string,
) (*v8.Value, error) {
	value, err := dom.NewUrlBase(url, base)
	if err != nil {
		return nil, err
	}
	u.Store(value, ctx, this)
	return nil, nil
}
