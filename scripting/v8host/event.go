package v8host

import (
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type eventV8Wrapper struct {
	handleReffedObject[*event.Event, jsTypeParam]
}

func newEventV8Wrapper(scriptHost jsScriptEngine) *eventV8Wrapper {
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
	return w.store(e, cbCtx)
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
	return cbCtx.ReturnWithTypeError(fmt.Sprintf(
		"encode EventTarget: Not an antity. %s",
		constants.MISSING_FEATURE_ISSUE_URL),
	)
}

func (w eventV8Wrapper) eventPhase(cbCtx jsCallbackContext) (jsValue, error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return w.toUnsignedShort(cbCtx, int(instance.EventPhase))
}
