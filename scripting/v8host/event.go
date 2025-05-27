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
) js.CallbackRVal {
	e := &event.Event{Type: type_, Bubbles: o.bubbles, Cancelable: o.cancelable, Data: o.init}
	return cbCtx.ReturnWithValueErr(w.store(e, cbCtx.ScriptCtx(), cbCtx.This()))
}

func (w eventV8Wrapper) toEventTarget(
	cbCtx *argumentHelper,
	e event.EventTarget,
) js.CallbackRVal {
	if e == nil {
		return cbCtx.ReturnWithValue(v8.Null(w.scriptHost.iso))
	}
	if entity, ok := e.(entity.ObjectIder); ok {
		return cbCtx.ReturnWithJSValueErr(cbCtx.ScriptCtx().getJSInstance(entity))
	}
	return cbCtx.ReturnWithError(v8.NewError(w.iso(), "TODO, Not yet supported"))
}

func (w eventV8Wrapper) eventPhase(cbCtx *argumentHelper) js.CallbackRVal {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	return cbCtx.ReturnWithValueErr(v8.NewValue(w.iso(), uint32(instance.EventPhase)))
}
