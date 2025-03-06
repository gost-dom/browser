package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/internal/entity"
)

type genericElementWrapper struct {
	baseInstanceWrapper[entity.ObjectIder]
}

func newGenericElementWrapper(instance *GojaContext) wrapper {
	return genericElementWrapper{newBaseInstanceWrapper[entity.ObjectIder](instance)}
}
func (w genericElementWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
