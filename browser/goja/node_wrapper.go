package goja

import (
	. "github.com/dop251/goja"
)

func (w nodeWrapper) constructor(call ConstructorCall, r *Runtime) *Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func (w nodeWrapper) NodeType(c FunctionCall) Value {
	instance := w.getInstance(c)
	return w.toUnsignedShort(int(instance.NodeType()))
}
