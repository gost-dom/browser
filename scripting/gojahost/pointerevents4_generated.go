// This file is generated. Do not edit.

package gojahost

import (
	"errors"
	g "github.com/dop251/goja"
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

func (w pointerEventWrapper) width(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.width")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventWrapper) height(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.height")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventWrapper) pressure(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.pressure")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventWrapper) tangentialPressure(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.tangentialPressure")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
