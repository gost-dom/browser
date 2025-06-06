// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	event "github.com/gost-dom/browser/dom/event"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	installClass("Event", "", newEventWrapper)
}

type eventWrapper struct {
	baseInstanceWrapper[*event.Event]
}

func newEventWrapper(instance *GojaContext) wrapper {
	return &eventWrapper{newBaseInstanceWrapper[*event.Event](instance)}
}

func (w eventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("stopPropagation", wrapCallback(w.ctx, w.stopPropagation))
	prototype.Set("preventDefault", wrapCallback(w.ctx, w.preventDefault))
	prototype.DefineAccessorProperty("type", wrapCallback(w.ctx, w.type_), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("target", wrapCallback(w.ctx, w.target), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("currentTarget", wrapCallback(w.ctx, w.currentTarget), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("eventPhase", wrapCallback(w.ctx, w.eventPhase), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("bubbles", wrapCallback(w.ctx, w.bubbles), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("cancelable", wrapCallback(w.ctx, w.cancelable), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("defaultPrevented", wrapCallback(w.ctx, w.defaultPrevented), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w eventWrapper) stopPropagation(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.stopPropagation")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	instance.StopPropagation()
	return cbCtx.ReturnWithValue(nil)
}

func (w eventWrapper) preventDefault(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.preventDefault")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	instance.PreventDefault()
	return cbCtx.ReturnWithValue(nil)
}

func (w eventWrapper) type_(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.type_")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Type
	return cbCtx.ReturnWithValueErr(w.toString_(cbCtx, result))
}

func (w eventWrapper) target(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.target")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Target
	return cbCtx.ReturnWithValueErr(w.toEventTarget(cbCtx, result))
}

func (w eventWrapper) currentTarget(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.currentTarget")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.CurrentTarget
	return cbCtx.ReturnWithValueErr(w.toEventTarget(cbCtx, result))
}

func (w eventWrapper) bubbles(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.bubbles")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Bubbles
	return cbCtx.ReturnWithValueErr(w.toBoolean(cbCtx, result))
}

func (w eventWrapper) cancelable(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.cancelable")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.Cancelable
	return cbCtx.ReturnWithValueErr(w.toBoolean(cbCtx, result))
}

func (w eventWrapper) defaultPrevented(cbCtx *callbackContext) g.Value {
	cbCtx.Logger().Debug("V8 Function call: Event.defaultPrevented")
	instance, instErr := js.As[*event.Event](cbCtx.Instance())
	if instErr != nil {
		return cbCtx.ReturnWithError(instErr)
	}
	result := instance.DefaultPrevented
	return cbCtx.ReturnWithValueErr(w.toBoolean(cbCtx, result))
}
