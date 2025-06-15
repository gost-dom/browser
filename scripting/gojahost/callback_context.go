package gojahost

import (
	"errors"
	"log/slog"

	"github.com/dop251/goja"
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaCallbackScope struct {
	ctx  *GojaContext
	this *goja.Object
}

func (c gojaCallbackScope) This() js.Object[jsTypeParam] {
	return newGojaObject(c.ctx, c.this)
}

func (ctx gojaCallbackScope) Instance() (any, error) {
	ctx.Logger().Debug("gojaCallbackScope.Instance()", "this", ctx.this)
	instance := ctx.this.GetSymbol(ctx.ctx.wrappedGoObj)
	if instance == nil {
		// TODO: Should be a TypeError
		return nil, errors.New("No embedded value")
	}
	return instance.Export(), nil
}

func (ctx gojaCallbackScope) Logger() *slog.Logger {
	if l := ctx.ctx.logger(); l != nil {
		return l
	}
	return log.Default()
}

type callbackContext struct {
	gojaCallbackScope
	args     []goja.Value
	argIndex int
}

func (c gojaCallbackScope) Scope() js.Scope[jsTypeParam] {
	return newGojaScope(c.ctx)
}

func (c gojaCallbackScope) ValueFactory() js.ValueFactory[jsTypeParam] {
	return newGojaValueFactory(c.ctx)
}

func newArgumentHelper(ctx *GojaContext, c goja.FunctionCall) *callbackContext {
	return &callbackContext{
		gojaCallbackScope{
			ctx,
			c.This.ToObject(ctx.vm),
		}, c.Arguments, 0}
}

func (ctx *callbackContext) Argument(index int) g.Value {
	return ctx.args[index]
}

func wrapJSCallback(ctx *GojaContext, cb js.FunctionCallback[jsTypeParam]) goja.Value {
	return ctx.vm.ToValue(func(c goja.FunctionCall) goja.Value {
		res, err := cb(newArgumentHelper(ctx, c))
		if err != nil {
			panic(err)
		}
		if res == nil {
			return goja.Undefined()
		}
		return res.Self().value
	})
}

func (c *callbackContext) ConsumeArg() (js.Value[jsTypeParam], bool) {
	index := c.argIndex
	c.argIndex++
	if index >= len(c.args) {
		return nil, false
	}
	return newGojaValue(c.ctx, c.args[index]), true
}

func (c *callbackContext) ReturnWithTypeError(msg string) (js.Value[jsTypeParam], error) {
	return nil, c.ValueFactory().NewTypeError(msg)
}
