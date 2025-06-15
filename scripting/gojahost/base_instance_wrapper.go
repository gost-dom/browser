package gojahost

import (
	"github.com/dop251/goja"
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type baseInstanceWrapper[T any] struct {
	ctx *GojaContext
}

func (w baseInstanceWrapper[T]) toGojaVal(v g.Value) (js.Value[jsTypeParam], error) {
	return newGojaValue(w.ctx, v), nil
}

func (w baseInstanceWrapper[T]) vm() *goja.Runtime {
	return w.ctx.vm
}

func newBaseInstanceWrapper[T any](instance *GojaContext) baseInstanceWrapper[T] {
	return baseInstanceWrapper[T]{instance}
}

func (c *GojaContext) storeInternal(value any, obj *g.Object) {
	log.Debug(c.logger(), "storeInternal")
	obj.DefineDataPropertySymbol(
		c.wrappedGoObj,
		c.vm.ToValue(value),
		g.FLAG_FALSE,
		g.FLAG_FALSE,
		g.FLAG_FALSE,
	)
	if e, ok := value.(entity.ObjectIder); ok {
		c.cachedNodes[e.ObjectId()] = obj
	}
}

func (w baseInstanceWrapper[T]) storeInternal(value any, obj *g.Object) {
	w.ctx.storeInternal(value, obj)
}

func getInstanceValue[T any](c *GojaContext, v g.Value) (T, bool) {
	res, ok := v.(*g.Object).GetSymbol(c.wrappedGoObj).Export().(T)
	return res, ok
}

func (w baseInstanceWrapper[T]) getInstance(c g.FunctionCall) T {
	if c.This == nil {
		panic("No this pointer")
	}
	if res, ok := getInstanceValue[T](w.ctx, c.This); ok {
		return res
	} else {
		panic(w.ctx.vm.NewTypeError("Not an entity"))
	}
}
