// This file is generated. Do not edit.

package gojahost

import (
	g "github.com/dop251/goja"
	log "github.com/gost-dom/browser/internal/log"
)

func init() {
	installClass("Event", "", newEventWrapper)
}

func (w eventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("stopPropagation", w.stopPropagation)
	prototype.Set("preventDefault", w.preventDefault)
	prototype.DefineAccessorProperty("type", w.ctx.vm.ToValue(w.type_), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("target", w.ctx.vm.ToValue(w.target), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("currentTarget", w.ctx.vm.ToValue(w.currentTarget), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("eventPhase", w.ctx.vm.ToValue(w.eventPhase), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("bubbles", w.ctx.vm.ToValue(w.bubbles), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("cancelable", w.ctx.vm.ToValue(w.cancelable), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("defaultPrevented", w.ctx.vm.ToValue(w.defaultPrevented), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w eventWrapper) Constructor(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.Constructor")
	panic("Goja constructor not yet implemented")
}

func (w eventWrapper) stopPropagation(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.stopPropagation")
	instance := w.getInstance(c)
	instance.StopPropagation()
	return nil
}

func (w eventWrapper) preventDefault(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.preventDefault")
	instance := w.getInstance(c)
	instance.PreventDefault()
	return nil
}

func (w eventWrapper) target(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.target")
	instance := w.getInstance(c)
	result := instance.Target
	return w.toEventTarget(result)
}

func (w eventWrapper) currentTarget(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.currentTarget")
	instance := w.getInstance(c)
	result := instance.CurrentTarget
	return w.toEventTarget(result)
}

func (w eventWrapper) defaultPrevented(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.defaultPrevented")
	instance := w.getInstance(c)
	result := instance.DefaultPrevented
	return w.toBoolean(result)
}
