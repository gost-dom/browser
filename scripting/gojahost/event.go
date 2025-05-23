package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/entity"
)

type gojaEvent[T event.Event] struct {
	Value *g.Object
	Event T
}

func toBoolean(value g.Value) bool {
	return value != nil && value.ToBoolean()
}

type eventWrapper struct {
	baseInstanceWrapper[*event.Event]
}

func newEventWrapper(instance *GojaContext) wrapper {
	return &eventWrapper{newBaseInstanceWrapper[*event.Event](instance)}
}

func (w eventWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	arg1 := call.Argument(0).String()
	init := event.CustomEventInit{}
	newInstance := &event.Event{Type: arg1, Data: init}
	if arg2 := call.Argument(1); !g.IsUndefined(arg2) {
		if obj, ok := arg2.(*g.Object); ok {
			newInstance.Bubbles = toBoolean(obj.Get("bubbles"))
			newInstance.Cancelable = toBoolean(obj.Get("cancelable"))
		}
	}
	w.storeInternal(newInstance, call.This)
	return nil
}

func (w eventWrapper) type_(c g.FunctionCall) g.Value {
	return w.ctx.vm.ToValue(w.getInstance(c).Type)
}

func (w eventWrapper) cancelable(c g.FunctionCall) g.Value {
	return w.ctx.vm.ToValue(w.getInstance(c).Cancelable)
}

func (w eventWrapper) bubbles(c g.FunctionCall) g.Value {
	return w.ctx.vm.ToValue(w.getInstance(c).Bubbles)
}

func (w eventWrapper) eventPhase(c g.FunctionCall) g.Value {
	return w.ctx.vm.ToValue(w.getInstance(c).EventPhase)
}

type customEventWrapper struct {
	eventWrapper
}

func newCustomEventWrapper(instance *GojaContext) wrapper {
	return customEventWrapper{eventWrapper{newBaseInstanceWrapper[*event.Event](instance)}}
}

func (w eventWrapper) toEventTarget(t event.EventTarget) g.Value {
	if t == nil {
		return nil
	}
	if ider, ok := t.(entity.ObjectIder); ok {
		return w.toNode(ider)
	}
	panic("TODO: Handle instances of non-entity events")
}
