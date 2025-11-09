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
		gojaScope: newGojaScope(ctx),
		this:      this,
		instance:  instance,
	}
}

func (s gojaCallbackScope) This() js.Object[jsTypeParam] {
	return newGojaObject(s.GojaContext, s.this)
}

func (s gojaCallbackScope) Instance() (any, error) {
	if s.instance == nil {
		panic(s.vm.NewTypeError("No embedded value"))
	}
	return s.instance, nil
}

func (s gojaCallbackScope) Logger() *slog.Logger {
	if l := s.logger(); l != nil {
		return l
	}
	return log.Default()
}

type callbackContext struct {
	gojaCallbackScope
	args     []goja.Value
	argIndex int
}

func newArgumentHelper(ctx *GojaContext, c goja.FunctionCall) *callbackContext {
	// I would consider this a bug in goja. When calling a function in global
	// scope, `this` is "undefined". It should have been `globalThis`.
	callThis := c.This
	if !callThis.ToBoolean() {
		callThis = ctx.vm.GlobalObject()
	}
	this := callThis.ToObject(ctx.vm)
	var instance any
	if wrapped := this.GetSymbol(ctx.wrappedGoObj); wrapped != nil && wrapped.ToBoolean() {
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
			panic(ctx.vm.ToValue(err))
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
	return newGojaValue(c.GojaContext, c.args[index]), true
}

func (c *callbackContext) ReturnWithTypeError(msg string) (js.Value[jsTypeParam], error) {
	return nil, c.NewTypeError(msg)
}
