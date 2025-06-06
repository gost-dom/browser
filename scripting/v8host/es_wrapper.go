package v8host

import (
	"errors"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type converters[T any] struct{}

type eventInitWrapper struct {
	bubbles    bool
	cancelable bool
	init       any
}

func (w converters[T]) decodeEventInit(
	_ js.CallbackContext[T],
	val js.Value[T],
) (eventInitWrapper, error) {
	options, ok := val.AsObject()
	if !ok {
		return eventInitWrapper{}, errors.New("Not an event init object")
	}

	bubbles, err1 := options.Get("bubbles")
	cancelable, err2 := options.Get("cancelable")
	err := errors.Join(err1, err2)
	if err != nil {
		return eventInitWrapper{}, err
	}
	init := eventInitWrapper{
		bubbles:    bubbles.Boolean(),
		cancelable: cancelable.Boolean(),
	}
	return init, nil
}

func (c converters[T]) defaultHTMLElement() html.HTMLElement { return nil }

func (w converters[T]) toNull(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return cbCtx.ValueFactory().Null(), nil
}

func (w converters[T]) toNullableString_(
	cbCtx js.CallbackContext[T],
	str *string,
) (js.Value[T], error) {
	if str == nil {
		return w.toNull(cbCtx)
	}
	return w.toString_(cbCtx, *str)
}

func (w converters[T]) toNillableString_(
	cbCtx js.CallbackContext[T],
	str string,
	hasVal bool,
) (js.Value[T], error) {
	if !hasVal {
		return w.toNull(cbCtx)
	}
	return w.toString_(cbCtx, str)
}

func (w converters[T]) toUnsignedLong(cbCtx js.CallbackContext[T], val int) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewUint32(uint32(val)), nil
}

func (w converters[T]) toUnsignedShort(cbCtx js.CallbackContext[T], val int) (js.Value[T], error) {
	// TODO: This should be uint16 - but v8go doesn't support uint16
	return cbCtx.ValueFactory().NewUint32(uint32(val)), nil
}

func (w converters[T]) toLong(cbCtx js.CallbackContext[T], val int) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewInt64(int64(val)), nil
}

func (w converters[T]) toAny(cbCtx js.CallbackContext[T], val string) (js.Value[T], error) {
	return w.toString_(cbCtx, val)
}

func (w converters[T]) toString_(cbCtx js.CallbackContext[T], val string) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewString(val), nil
}

func (w converters[T]) toBoolean(cbCtx js.CallbackContext[T], val bool) (js.Value[T], error) {
	return cbCtx.ValueFactory().NewBoolean(val), nil
}

func (w converters[T]) toJSWrapper(
	cbCtx js.CallbackContext[T],
	val entity.ObjectIder,
) (js.Value[T], error) {
	return encodeEntity(cbCtx, val)
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

func (o handleReffedObject[T, U]) store(
	value any,
	cbCtx js.CallbackContext[U],
) (js.Value[U], error) {
	this := cbCtx.This()
	if e, ok := value.(entity.ObjectIder); ok {
		cbCtx.Scope().SetValue(e, this)
	}

	this.SetNativeValue(value)
	return this, nil
}
