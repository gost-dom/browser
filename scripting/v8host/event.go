package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type eventV8Wrapper struct {
	handleReffedObject[*event.Event, jsTypeParam]
}

func newEventV8Wrapper(scriptHost *V8ScriptHost) *eventV8Wrapper {
	return &eventV8Wrapper{newHandleReffedObject[*event.Event](scriptHost)}
}

func (w eventV8Wrapper) defaultEventInit() eventInitWrapper {
	return eventInitWrapper{}
}

func (w eventV8Wrapper) CreateInstance(
	cbCtx jsCallbackContext,
	type_ string,
	o eventInitWrapper,
) (jsValue, error) {
	e := &event.Event{Type: type_, Bubbles: o.bubbles, Cancelable: o.cancelable, Data: o.init}
	return cbCtx.ReturnWithJSValueErr(w.store(e, cbCtx.ScriptCtx(), cbCtx.This()))
}

func (w eventV8Wrapper) toEventTarget(
	cbCtx jsCallbackContext,
	e event.EventTarget,
) (jsValue, error) {
	if e == nil {
		return cbCtx.ReturnWithValue(cbCtx.ValueFactory().Null())
	}
	if entity, ok := e.(entity.ObjectIder); ok {
		return w.toJSWrapper(cbCtx, entity)
	}
	return cbCtx.ReturnWithError(errors.New("TODO, Not yet supported"))
}

func (w eventV8Wrapper) eventPhase(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return w.toUnsignedShort(cbCtx, int(instance.EventPhase))
}
