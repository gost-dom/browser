// This file is generated. Do not edit.

package gojahost

import (
	"errors"
	g "github.com/dop251/goja"
	uievents "github.com/gost-dom/browser/internal/uievents"
)

func init() {
	installClass("MouseEvent", "UIEvent", newMouseEventWrapper)
}

func (w mouseEventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("getModifierState", wrapCallback(w.ctx, w.getModifierState))
	prototype.DefineAccessorProperty("screenX", wrapCallback(w.ctx, w.screenX), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("screenY", wrapCallback(w.ctx, w.screenY), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("clientX", wrapCallback(w.ctx, w.clientX), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("clientY", wrapCallback(w.ctx, w.clientY), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("layerX", wrapCallback(w.ctx, w.layerX), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("layerY", wrapCallback(w.ctx, w.layerY), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("relatedTarget", wrapCallback(w.ctx, w.relatedTarget), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w mouseEventWrapper) getModifierState(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.getModifierState")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventWrapper) screenX(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventWrapper) screenY(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventWrapper) clientX(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventWrapper) clientY(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventWrapper) layerX(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventWrapper) layerY(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventWrapper) relatedTarget(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.relatedTarget")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
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
	prototype.DefineAccessorProperty("view", wrapCallback(w.ctx, w.view), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("detail", wrapCallback(w.ctx, w.detail), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w uIEventWrapper) view(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.view")
	return cbCtx.ReturnWithError(errors.New("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w uIEventWrapper) detail(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.detail")
	return cbCtx.ReturnWithError(errors.New("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
