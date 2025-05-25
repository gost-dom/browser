package gojahost

import "github.com/dop251/goja"

type callbackContext struct {
	ctx *GojaContext
}

func newArgumentHelper(ctx *GojaContext, c goja.FunctionCall) *callbackContext {
	return &callbackContext{ctx}
}

func (ctx *callbackContext) ReturnWithValue(val goja.Value) goja.Value { return val }
func (ctx *callbackContext) ReturnWithTypeError(msg string) goja.Value {
	panic(ctx.ctx.vm.NewTypeError(msg))
}

type callbackFunction = func(goja.FunctionCall) goja.Value

func wrapCallback(ctx *GojaContext, cb callbackFunction) goja.Value {
	return ctx.vm.ToValue(cb)
}
