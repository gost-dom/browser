package sobekhost

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

type promise struct {
	js.Value[jsTypeParam]
	resolve func(any) error
	reject  func(any) error
}

func newPromise(ctx *scriptContext) promise {
	p, resolve, reject := ctx.vm.NewPromise()
	return promise{
		newValue(ctx, ctx.vm.ToValue(p)),
		resolve, reject,
	}
}

func (p promise) Resolve(value js.Value[jsTypeParam]) { p.resolve(value.Self().value) }
func (p promise) Reject(value js.Value[jsTypeParam])  { p.reject(value.Self().value) }
