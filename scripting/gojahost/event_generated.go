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
	args := newArgumentHelper(w.ctx, c)
	return args.ReturnWithTypeError("Goja constructor not yet implemented")
}

func (w eventWrapper) stopPropagation(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.stopPropagation")
	args := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	instance.StopPropagation()
	return args.ReturnWithValue(nil)
}

func (w eventWrapper) preventDefault(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.preventDefault")
	args := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	instance.PreventDefault()
	return args.ReturnWithValue(nil)
}

func (w eventWrapper) target(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.target")
	args := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.Target
	return args.ReturnWithValue(w.toEventTarget(result))
}

func (w eventWrapper) currentTarget(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.currentTarget")
	args := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.CurrentTarget
	return args.ReturnWithValue(w.toEventTarget(result))
}

func (w eventWrapper) defaultPrevented(c g.FunctionCall) g.Value {
	log.Debug(w.logger(c), "V8 Function call: Event.defaultPrevented")
	args := newArgumentHelper(w.ctx, c)
	instance := w.getInstance(c)
	result := instance.DefaultPrevented
	return args.ReturnWithValue(w.toBoolean(result))
}
