package dom

import (
	"fmt"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func CreateEvent[T any](
	cbCtx js.CallbackContext[T],
	type_ string,
	o codec.EventInit,
) (js.Value[T], error) {
	e := &event.Event{Type: type_, Bubbles: o.Bubbles, Cancelable: o.Cancelable, Data: o.Init}
	return codec.EncodeConstructedValue(cbCtx, e)
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
	return nil, cbCtx.NewTypeError(fmt.Sprintf(
		"encode EventTarget: Not an entity. %s",
		constants.MISSING_FEATURE_ISSUE_URL),
	)
}
func encodeEventPhase[T any](scope js.Scope[T], e event.EventPhase) (js.Value[T], error) {
	return codec.EncodeInt(scope, int(e))
}
