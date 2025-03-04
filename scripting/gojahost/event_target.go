package gojahost

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/events"

	"github.com/dop251/goja"
)

type eventTargetWrapper struct {
	baseInstanceWrapper[events.EventTarget]
}

func newEventTargetWrapper(instance *GojaContext) wrapper {
	return eventTargetWrapper{newBaseInstanceWrapper[events.EventTarget](instance)}
}

type gojaEventListener struct {
	instance *GojaContext
	v        goja.Value
	f        goja.Callable
}

func newGojaEventListener(r *GojaContext, v goja.Value) events.EventHandler {
	f, ok := goja.AssertFunction(v)
	if !ok {
		panic("TODO")
	}
	return &gojaEventListener{r, v, f}
}

func (h *gojaEventListener) HandleEvent(e *events.Event) error {
	customEvent := h.instance.globals["Event"]
	switch e.EventInit.(type) {
	case dom.PointerEventInitDict:
		customEvent = h.instance.globals["PointerEvent"]
	case dom.MouseEventInitDict:
		customEvent = h.instance.globals["MouseEvent"]
	case dom.UIEventInitDict:
		customEvent = h.instance.globals["UIEvent"]
	case events.CustomEventInitDict:
		customEvent = h.instance.globals["CustomEvent"]
	default:
		customEvent = h.instance.globals["Event"]
	}
	obj := h.instance.vm.CreateObject(customEvent.Prototype)
	customEvent.Wrapper.storeInternal(e, obj)
	_, err := h.f(obj, obj)
	return err
}

func (h *gojaEventListener) Equals(e events.EventHandler) bool {
	if ge, ok := e.(*gojaEventListener); ok && ge.v.StrictEquals(h.v) {
		return true
	} else {
		return false
	}
}

func (w eventTargetWrapper) constructor(
	call goja.ConstructorCall,
	r *goja.Runtime,
) *goja.Object {
	newInstance := events.NewEventTarget()
	w.storeInternal(newInstance, call.This)
	return nil
}

func (w eventTargetWrapper) getEventTarget(c goja.FunctionCall) events.EventTarget {
	if c.This == nil {
		panic("No this pointer")
	}
	if c.This == w.ctx.vm.GlobalObject() {
		return w.ctx.window
	}
	instance, ok := c.This.Export().(events.EventTarget)
	if !ok {
		panic(w.ctx.vm.NewTypeError("Not an event target"))
	}
	return instance
}

func (w eventTargetWrapper) addEventListener(c goja.FunctionCall) goja.Value {
	instance := w.getInstance(c)
	name := c.Argument(0).String()
	instance.AddEventListener(name, newGojaEventListener(w.ctx, c.Argument(1)))
	return nil
}

func (w eventTargetWrapper) dispatchEvent(c goja.FunctionCall) goja.Value {
	instance := w.getInstance(c)
	internal := c.Argument(0).(*goja.Object).GetSymbol(w.ctx.wrappedGoObj).Export()
	if event, ok := internal.(*events.Event); ok {
		return w.ctx.vm.ToValue(instance.DispatchEvent(event))
	} else {
		panic(w.ctx.vm.NewTypeError("Not an event"))
	}
}

func (w eventTargetWrapper) initializePrototype(prototype *goja.Object,
	vm *goja.Runtime) {
	prototype.Set("addEventListener", w.addEventListener)
	prototype.Set("dispatchEvent", w.dispatchEvent)
}
