// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	dom "github.com/gost-dom/browser/dom"
)

func init() {
	installClass("MouseEvent", "UIEvent", newMouseEventWrapper)
}

func (w mouseEventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("getModifierState", w.getModifierState)
	prototype.DefineAccessorProperty("screenX", w.ctx.vm.ToValue(w.screenX), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("screenY", w.ctx.vm.ToValue(w.screenY), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("clientX", w.ctx.vm.ToValue(w.clientX), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("clientY", w.ctx.vm.ToValue(w.clientY), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("layerX", w.ctx.vm.ToValue(w.layerX), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("layerY", w.ctx.vm.ToValue(w.layerY), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("relatedTarget", w.ctx.vm.ToValue(w.relatedTarget), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w mouseEventWrapper) getModifierState(c g.FunctionCall) g.Value {
	panic("MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) screenX(c g.FunctionCall) g.Value {
	panic("MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) screenY(c g.FunctionCall) g.Value {
	panic("MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) clientX(c g.FunctionCall) g.Value {
	panic("MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) clientY(c g.FunctionCall) g.Value {
	panic("MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) layerX(c g.FunctionCall) g.Value {
	panic("MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) layerY(c g.FunctionCall) g.Value {
	panic("MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) relatedTarget(c g.FunctionCall) g.Value {
	panic("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func init() {
	installClass("UIEvent", "Event", newUIEventWrapper)
}

type uIEventWrapper struct {
	baseInstanceWrapper[dom.UIEvent]
}

func newUIEventWrapper(instance *GojaContext) wrapper {
	return &uIEventWrapper{newBaseInstanceWrapper[dom.UIEvent](instance)}
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
