package sobekhost

import "github.com/gost-dom/browser/scripting/internal/js"

type scriptError struct {
	js.Value[jsTypeParam]
	error
}

func newScriptError(ctx *scriptContext, err error) js.Error[jsTypeParam] {
	obj := newObject(ctx, ctx.vm.NewGoError(err))
	return scriptError{obj, err}
}
