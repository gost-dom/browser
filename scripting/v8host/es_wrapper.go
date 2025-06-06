package v8host

import (
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type converters[T any] struct{}

type eventInitWrapper struct {
	bubbles    bool
	cancelable bool
	init       any
}

func (w converters[T]) toString_(
	cbCtx js.CallbackContext[T],
	val string,
) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewString(val), nil
}

// handleReffedObject serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type handleReffedObject[T, U any] struct {
	scriptHost js.ScriptEngine[U]
	converters[U]
}

func newHandleReffedObject[T, U any](
	host js.ScriptEngine[U],
) handleReffedObject[T, U] {
	return handleReffedObject[T, U]{
		scriptHost: host,
	}
}

func storeNewValue[T any](
	value any,
	cbCtx js.CallbackContext[T],
) (js.Value[T], error) {
	this := cbCtx.This()
	if e, ok := value.(entity.ObjectIder); ok {
		cbCtx.Scope().SetValue(e, this)
	}

	this.SetNativeValue(value)
	return this, nil
}
