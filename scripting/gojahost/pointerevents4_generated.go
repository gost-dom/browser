// This file is generated. Do not edit.

package gojahost

import g "github.com/dop251/goja"

func init() {
	installClass("PointerEvent", "MouseEvent", newPointerEventWrapper)
}

func (w pointerEventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.DefineAccessorProperty("width", w.ctx.vm.ToValue(w.width), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("height", w.ctx.vm.ToValue(w.height), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("pressure", w.ctx.vm.ToValue(w.pressure), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("tangentialPressure", w.ctx.vm.ToValue(w.tangentialPressure), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w pointerEventWrapper) width(c g.FunctionCall) g.Value {
	panic("PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventWrapper) height(c g.FunctionCall) g.Value {
	panic("PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventWrapper) pressure(c g.FunctionCall) g.Value {
	panic("PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventWrapper) tangentialPressure(c g.FunctionCall) g.Value {
	panic("PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
