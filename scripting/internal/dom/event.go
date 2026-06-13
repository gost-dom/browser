package dom

import (
	"fmt"
	"strings"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/internal/entity"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// Document_createEvent implements the legacy Document.createEvent factory. It
// returns an uninitialized event whose class matches the requested legacy
// interface name (e.g. "MouseEvents" -> MouseEvent). The returned event is
// expected to be initialized via the legacy initEvent/initMouseEvent methods
// before use.
func Document_createEvent[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	args := cbCtx.Args()
	kind := ""
	if len(args) > 0 && args[0] != nil {
		kind = args[0].String()
	}
	className := "Event"
	switch strings.ToLower(kind) {
	case "mouseevent", "mouseevents":
		className = "MouseEvent"
	case "pointerevent", "pointerevents":
		className = "PointerEvent"
	case "keyboardevent", "keyboardevents", "keyevents":
		className = "KeyboardEvent"
	case "uievent", "uievents":
		className = "UIEvent"
	case "customevent":
		className = "CustomEvent"
	}
	return cbCtx.Constructor(className).NewInstance(&event.Event{})
}

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
