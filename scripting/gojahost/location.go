package gojahost

import (
	"github.com/dop251/goja"
	"github.com/gost-dom/browser/html"
)

type locationWrapper struct {
	baseInstanceWrapper[html.Location]
}

func newLocationWrapper(instance *GojaContext) wrapper {
	return &locationWrapper{newBaseInstanceWrapper[html.Location](instance)}
}

func (w locationWrapper) constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

// func (w locationWrapper) initializePrototype(prototype *goja.Object, vm *goja.Runtime) {
// }
