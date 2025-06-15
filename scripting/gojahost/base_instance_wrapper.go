package gojahost

import (
	"github.com/dop251/goja"
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
)

type baseInstanceWrapper[T any] struct {
	ctx *GojaContext
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
