package v8host

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type eventV8Wrapper struct {
	handleReffedObject[*event.Event]
}

func newEventV8Wrapper(scriptHost *V8ScriptHost) *eventV8Wrapper {
	return &eventV8Wrapper{newHandleReffedObject[*event.Event](scriptHost)}
}

func (w eventV8Wrapper) defaultEventInit() eventInitWrapper {
	return eventInitWrapper{}
}

func (w eventV8Wrapper) CreateInstance(
	cbCtx *argumentHelper,
	type_ string,
	o eventInitWrapper,
) (*v8.Value, error) {
	e := &event.Event{Type: type_, Bubbles: o.bubbles, Cancelable: o.cancelable, Data: o.init}
	return w.store(e, cbCtx.ScriptCtx(), cbCtx.This())
}

func (w eventV8Wrapper) toEventTarget(
	ctx *V8ScriptContext,
	e event.EventTarget,
) (*v8.Value, error) {
	if e == nil {
		return v8.Null(w.scriptHost.iso), nil
	}
	if entity, ok := e.(entity.ObjectIder); ok {
		return ctx.getInstanceForNode(entity)
	}
	return nil, v8.NewError(w.iso(), "TODO, Not yet supported")
}

func (w eventV8Wrapper) eventPhase(cbCtx *argumentHelper) (*v8.Value, error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return v8.NewValue(w.iso(), uint32(instance.EventPhase))
}

func (w eventV8Wrapper) type_(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.type")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return v8.NewValue(w.iso(), instance.Type)
}

func (w eventV8Wrapper) cancelable(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.cancelable")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return v8.NewValue(w.iso(), instance.Cancelable)
}

func (w eventV8Wrapper) bubbles(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.bubbles")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return v8.NewValue(w.iso(), instance.Bubbles)
}
