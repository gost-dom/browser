package sobekengine

import (
	"log/slog"

	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/grafana/sobek"
)

type callbackScope struct {
	scope
	this     *sobek.Object
	instance any
}

func newCallbackScope(ctx *scriptContext, this *sobek.Object, instance any) callbackScope {
	return callbackScope{
		scope:    newScope(ctx),
		this:     this,
		instance: instance,
	}
}

func (s callbackScope) This() js.Object[jsTypeParam] {
	return newObject(s.scriptContext, s.this)
}

func (s callbackScope) Instance() (any, error) {
	if s.instance == nil {
		panic(s.vm.NewTypeError("No embedded value"))
	}
	return s.instance, nil
}

func (s callbackScope) Logger() *slog.Logger {
	if l := s.logger(); l != nil {
		return l
	}
	return log.Default()
}

type callbackContext struct {
	callbackScope
	args     []sobek.Value
	argIndex int
}

func newArgumentHelper(ctx *scriptContext, c sobek.FunctionCall) *callbackContext {
	// BUG: Consider if this is still an issue
	// I would consider this a bug in sobek. When calling a function in global
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

func (ctx *callbackContext) Argument(index int) sobek.Value {
	return ctx.args[index]
}

func wrapJSCallback(ctx *scriptContext, cb js.FunctionCallback[jsTypeParam]) sobek.Value {
	return ctx.vm.ToValue(func(c sobek.FunctionCall) sobek.Value {
		res, err := cb(newArgumentHelper(ctx, c))
		panicIfError(ctx, err)
		return unwrapValue(res)
	})
}

func panicIfError(ctx *scriptContext, err error) {
	if err != nil {
		if sobekErr, ok := err.(sobekError); ok {
			panic(sobekErr.Object)
		} else {
			panic(ctx.vm.ToValue(err))
		}
	}

}

func (c *callbackContext) Args() []js.Value[jsTypeParam] {
	res := make([]js.Value[jsTypeParam], len(c.args))
	for i, a := range c.args {
		res[i] = newValue(c.scriptContext, a)
	}
	return res
}

func (c *callbackContext) ConsumeArg() (js.Value[jsTypeParam], bool) {
	index := c.argIndex
	c.argIndex++
	if index >= len(c.args) {
		return nil, false
	}
	return newValue(c.scriptContext, c.args[index]), true
}

func (c *callbackContext) ReturnWithTypeError(msg string) (js.Value[jsTypeParam], error) {
	return nil, c.NewTypeError(msg)
}
