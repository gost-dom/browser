// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	uievents "github.com/gost-dom/browser/internal/uievents"
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
	prototype.DefineAccessorProperty("ctrlKey", w.ctx.vm.ToValue(w.ctrlKey), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("shiftKey", w.ctx.vm.ToValue(w.shiftKey), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("altKey", w.ctx.vm.ToValue(w.altKey), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("metaKey", w.ctx.vm.ToValue(w.metaKey), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("button", w.ctx.vm.ToValue(w.button), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("buttons", w.ctx.vm.ToValue(w.buttons), nil, g.FLAG_TRUE, g.FLAG_TRUE)
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

func (w mouseEventWrapper) ctrlKey(c g.FunctionCall) g.Value {
	panic("MouseEvent.ctrlKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) shiftKey(c g.FunctionCall) g.Value {
	panic("MouseEvent.shiftKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) altKey(c g.FunctionCall) g.Value {
	panic("MouseEvent.altKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) metaKey(c g.FunctionCall) g.Value {
	panic("MouseEvent.metaKey: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) button(c g.FunctionCall) g.Value {
	panic("MouseEvent.button: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) buttons(c g.FunctionCall) g.Value {
	panic("MouseEvent.buttons: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventWrapper) relatedTarget(c g.FunctionCall) g.Value {
	panic("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

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
