package gojahost

import (
	"github.com/dop251/goja"
)

func (w locationWrapper) constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
