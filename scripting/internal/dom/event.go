package dom

import (
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (w Event[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	type_ string,
	o codec.EventInit,
) (js.Value[T], error) {
	e := &event.Event{Type: type_, Bubbles: o.Bubbles, Cancelable: o.Cancelable, Data: o.Init}
	return codec.EncodeConstrucedValue(cbCtx, e)
}

func encodeEventTarget[T any](
	cbCtx js.CallbackContext[T],
	e event.EventTarget,
) (js.Value[T], error) {
	if e == nil {
		return cbCtx.Null(), nil
	}
	if entity, ok := e.(entity.Components); ok {
		return codec.EncodeEntity(cbCtx, entity)
	}
	return cbCtx.ReturnWithTypeError(fmt.Sprintf(
		"encode EventTarget: Not an antity. %s",
		constants.MISSING_FEATURE_ISSUE_URL),
	)
}

func (w Event[T]) eventPhase(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	return codec.EncodeInt(cbCtx, int(instance.EventPhase))
}
