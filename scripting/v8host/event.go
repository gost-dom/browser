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
	cbCtx *v8CallbackContext,
	type_ string,
	o eventInitWrapper,
) (jsValue, error) {
	e := &event.Event{Type: type_, Bubbles: o.bubbles, Cancelable: o.cancelable, Data: o.init}
	return cbCtx.ReturnWithJSValueErr(w.store(e, cbCtx.ScriptCtx(), cbCtx.This()))
}

func (w eventV8Wrapper) toEventTarget(
	cbCtx *v8CallbackContext,
	e event.EventTarget,
) (jsValue, error) {
	if e == nil {
		return cbCtx.ReturnWithValue(v8.Null(w.scriptHost.iso))
	}
	if entity, ok := e.(entity.ObjectIder); ok {
		return cbCtx.ReturnWithJSValueErr(cbCtx.ScriptCtx().getJSInstance(entity))
	}
	return cbCtx.ReturnWithError(v8.NewError(w.iso(), "TODO, Not yet supported"))
}

func (w eventV8Wrapper) eventPhase(cbCtx *v8CallbackContext) (jsValue, error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	return cbCtx.ReturnWithValueErr(v8.NewValue(w.iso(), uint32(instance.EventPhase)))
}
