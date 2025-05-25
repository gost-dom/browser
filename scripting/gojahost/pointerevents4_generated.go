// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	log "github.com/gost-dom/browser/internal/log"
)

func init() {
	installClass("PointerEvent", "MouseEvent", newPointerEventWrapper)
}

func (w pointerEventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.DefineAccessorProperty("width", wrapCallback(w.ctx, w.width), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("height", wrapCallback(w.ctx, w.height), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("pressure", wrapCallback(w.ctx, w.pressure), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("tangentialPressure", wrapCallback(w.ctx, w.tangentialPressure), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w pointerEventWrapper) Constructor(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: PointerEvent.Constructor")
	cbCtx := newArgumentHelper(w.ctx, c)
	return cbCtx.ReturnWithTypeError("Goja constructor not yet implemented")
}

func (w pointerEventWrapper) width(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: PointerEvent.width")
	panic("PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventWrapper) height(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: PointerEvent.height")
	panic("PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventWrapper) pressure(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: PointerEvent.pressure")
	panic("PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventWrapper) tangentialPressure(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: PointerEvent.tangentialPressure")
	panic("PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
