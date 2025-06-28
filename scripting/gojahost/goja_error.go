package gojahost

import "github.com/gost-dom/browser/scripting/internal/js"

type gojaError struct {
	js.Value[jsTypeParam]
	error
}

func newGojaError(ctx *GojaContext, err error) js.Error[jsTypeParam] {
	obj := newGojaObject(ctx, ctx.vm.NewGoError(err))
	return gojaError{obj, err}
}
