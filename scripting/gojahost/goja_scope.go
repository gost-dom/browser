package gojahost

import (
	"github.com/dop251/goja"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type gojaScope struct {
	ctx    *GojaContext
	global js.Object[jsTypeParam]
}

func newGojaScope(ctx *GojaContext) js.Scope[jsTypeParam] {
	return gojaScope{ctx, newGojaObject(ctx, ctx.vm.GlobalObject())}
}

func (s gojaScope) Window() html.Window                { return s.ctx.window }
func (s gojaScope) GlobalThis() js.Object[jsTypeParam] { return s.global }

func (s gojaScope) Clock() *clock.Clock { return s.ctx.clock }
func (s gojaScope) Constructor(name string) js.Constructor[jsTypeParam] {
	if f, ok := s.ctx.globals[name]; ok {
		return gojaConstructor{s.ctx, f.Prototype}
	}
	return nil
}

func (s gojaScope) GetValue(e entity.ObjectIder) (js.Value[jsTypeParam], bool) {
	v, ok := s.ctx.cachedNodes[e.ObjectId()]
	return newGojaValue(s.ctx, v), ok
}
func (s gojaScope) SetValue(e entity.ObjectIder, v js.Value[jsTypeParam]) {
	s.ctx.cachedNodes[e.ObjectId()] = v.Self().value
}

func (s gojaScope) ValueFactory() js.ValueFactory[jsTypeParam] {
	return newGojaValueFactory(s.ctx)
}

type gojaConstructor struct {
	ctx       *GojaContext
	prototype *goja.Object
}

func (c gojaConstructor) NewInstance(native any) (js.Object[jsTypeParam], error) {
	obj := c.ctx.vm.CreateObject(c.prototype)
	c.ctx.storeInternal(native, obj)
	return newGojaObject(c.ctx, obj), nil
}
