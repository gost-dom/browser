package gojahost

import (
	"log/slog"

	"github.com/dop251/goja"
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaCallbackScope struct {
	gojaScope
	this     *goja.Object
	instance any
}

func newCallbackScope(ctx *GojaContext, this *goja.Object, instance any) gojaCallbackScope {
	return gojaCallbackScope{
		gojaScope: gojaScope{ctx, newGojaObject(ctx, ctx.vm.GlobalObject())},
		this:      this,
		instance:  instance,
	}
}

func (c gojaCallbackScope) This() js.Object[jsTypeParam] {
	return newGojaObject(c.ctx, c.this)
}

func (ctx gojaCallbackScope) Instance() (any, error) {
	if ctx.instance == nil {
		panic(ctx.ctx.vm.NewTypeError("No embedded value"))
	}
	return ctx.instance, nil
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
	this := c.This.ToObject(ctx.vm)
	var instance any
	if wrapped := this.GetSymbol(ctx.wrappedGoObj); wrapped != nil {
		instance = wrapped.Export()
	}
	return &callbackContext{
		newCallbackScope(ctx, this, instance),
		c.Arguments, 0}
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
		return toGojaValue(res)
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
