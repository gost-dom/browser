package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/dom/events"
)

type eventWrapper struct {
	baseInstanceWrapper[*events.Event]
}

func newEventWrapperAsWrapper(instance *GojaContext) wrapper { return newEventWrapper(instance) }
func newEventWrapper(instance *GojaContext) eventWrapper {
	return eventWrapper{newBaseInstanceWrapper[*events.Event](instance)}
}

type gojaEvent[T events.Event] struct {
	Value *g.Object
	Event T
}

func toBoolean(value g.Value) bool {
	return value != nil && value.ToBoolean()
}

func (w eventWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	arg1 := call.Argument(0).String()
	options := make([]events.EventOption, 0, 2)
	if arg2 := call.Argument(1); !g.IsUndefined(arg2) {
		if obj, ok := arg2.(*g.Object); ok {
			options = append(options, events.EventCancelable(toBoolean(obj.Get("cancelable"))))
			options = append(options, events.EventBubbles(toBoolean(obj.Get("bubbles"))))
		}
	}
	newInstance := events.NewCustomEvent(arg1, options...)
	w.storeInternal(newInstance, call.This)
	return nil
}

func (w eventWrapper) PreventDefault(c g.FunctionCall) g.Value {
	w.getInstance(c).PreventDefault()
	return nil
}

func (w eventWrapper) GetType(c g.FunctionCall) g.Value {
	return w.ctx.vm.ToValue(w.getInstance(c).Type())
}

func (w eventWrapper) initializePrototype(prototype *g.Object,
	vm *g.Runtime) {
	prototype.Set("preventDefault", w.PreventDefault)
	prototype.DefineAccessorProperty(
		"type",
		w.ctx.vm.ToValue(w.GetType),
		nil,
		g.FLAG_TRUE,
		g.FLAG_TRUE,
	)
}

type customEventWrapper struct {
	eventWrapper
}

func newCustomEventWrapper(instance *GojaContext) wrapper {
	return customEventWrapper{newEventWrapper(instance)}
}
