package sobekhost

import (
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/grafana/sobek"
)

type baseInstanceWrapper[T any] struct {
	ctx *GojaContext
}

func (w baseInstanceWrapper[T]) vm() *sobek.Runtime {
	return w.ctx.vm
}

func newBaseInstanceWrapper[T any](instance *GojaContext) baseInstanceWrapper[T] {
	return baseInstanceWrapper[T]{instance}
}

func (c *GojaContext) storeInternal(value any, obj *sobek.Object) {
	log.Debug(c.logger(), "storeInternal")
	obj.DefineDataPropertySymbol(
		c.wrappedGoObj,
		c.vm.ToValue(value),
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
		sobek.FLAG_FALSE,
	)
	if e, ok := value.(entity.ObjectIder); ok {
		c.cachedNodes[e.ObjectId()] = obj
	}
}
