package gojahost

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaPromise struct {
	js.Value[jsTypeParam]
	resolve func(any) error
	reject  func(any) error
}

func newGojaPromise(ctx *GojaContext) gojaPromise {
	p, resolve, reject := ctx.vm.NewPromise()
	return gojaPromise{
		newGojaValue(ctx, ctx.vm.ToValue(p)),
		resolve, reject,
	}
}

func (p gojaPromise) Resolve(value js.Value[jsTypeParam]) {
	p.resolve(value.Self().value)
}

func (p gojaPromise) Reject(err error) {
	p.reject(p.Self().ctx.vm.NewGoError(err))
}
