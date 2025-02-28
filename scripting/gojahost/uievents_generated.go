// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	uievents "github.com/gost-dom/browser/internal/uievents"
)

func init() {
	installClass("UIEvent", "Event", newUIEventWrapper)
}

type uIEventWrapper struct {
	baseInstanceWrapper[uievents.UIEvent]
}

func newUIEventWrapper(instance *GojaContext) wrapper {
	return &uIEventWrapper{newBaseInstanceWrapper[uievents.UIEvent](instance)}
}

func (w uIEventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.DefineAccessorProperty("view", w.ctx.vm.ToValue(w.view), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("detail", w.ctx.vm.ToValue(w.detail), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w uIEventWrapper) view(c g.FunctionCall) g.Value {
	panic("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w uIEventWrapper) detail(c g.FunctionCall) g.Value {
	panic("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
