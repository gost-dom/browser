package v8host

import (
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type eventV8Wrapper[T any] struct {
	handleReffedObject[*event.Event, T]
}

func newEventV8Wrapper(scriptHost jsScriptEngine) *eventV8Wrapper[jsTypeParam] {
	return &eventV8Wrapper[jsTypeParam]{newHandleReffedObject[*event.Event](scriptHost)}
}

func (w eventV8Wrapper[T]) defaultEventInit() eventInitWrapper {
	return eventInitWrapper{}
}

func (w eventV8Wrapper[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	type_ string,
	o eventInitWrapper,
) (js.Value[T], error) {
	e := &event.Event{Type: type_, Bubbles: o.bubbles, Cancelable: o.cancelable, Data: o.init}
	return w.store(e, cbCtx)
}

func (w eventV8Wrapper[T]) toEventTarget(
	cbCtx js.CallbackContext[T],
	e event.EventTarget,
) (js.Value[T], error) {
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

func (w eventV8Wrapper[T]) eventPhase(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return w.toUnsignedShort(cbCtx, int(instance.EventPhase))
}
