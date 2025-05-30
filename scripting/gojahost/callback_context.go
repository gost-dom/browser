package gojahost

import (
	"errors"
	"log/slog"

	"github.com/dop251/goja"
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/internal/log"
)

type callbackContext struct {
	ctx      *GojaContext
	call     goja.FunctionCall
	argIndex int
}

func newArgumentHelper(ctx *GojaContext, c goja.FunctionCall) *callbackContext {
	return &callbackContext{ctx, c, 0}
}

func (ctx *callbackContext) Logger() *slog.Logger {
	if l := ctx.ctx.window.Logger(); l != nil {
		return l
	}
	return log.Default()
}

func (ctx *callbackContext) Argument(index int) g.Value {
	return ctx.call.Argument(index)
}

func (ctx *callbackContext) Instance() (any, error) {
	instance := ctx.call.This.(*g.Object).GetSymbol(ctx.ctx.wrappedGoObj)
	if instance == nil {
		// TODO: Should be a TypeError
		return nil, errors.New("No embedded value")
	}
	return instance.Export(), nil
}

func (ctx *callbackContext) ReturnWithValue(val goja.Value) goja.Value { return val }
func (ctx *callbackContext) ReturnWithTypeError(msg string) goja.Value {
	panic(ctx.ctx.vm.NewTypeError(msg))
}
func (ctx *callbackContext) ReturnWithError(err error) goja.Value {
	panic(err)
}

type callbackFunction = func(*callbackContext) goja.Value

func wrapCallback(ctx *GojaContext, cb callbackFunction) goja.Value {
	return ctx.vm.ToValue(func(c goja.FunctionCall) goja.Value {
		return cb(newArgumentHelper(ctx, c))
	})
}

func (c *callbackContext) consumeValue() g.Value {
	index := c.argIndex
	c.argIndex++
	return c.call.Argument(index)
}
