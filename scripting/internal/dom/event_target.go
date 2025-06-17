package dom

import (
	"errors"

	"github.com/gost-dom/browser/dom/event"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type v8EventListener[T any] struct {
	// TODO: Replace with "scope" - as we keep on to this for longer than the
	// callback
	ctx js.CallbackContext[T]
	val js.Function[T]
}

func newV8EventListener[T any](ctx js.CallbackContext[T], val js.Function[T]) event.EventHandler {
	return v8EventListener[T]{ctx, val}
}

func (l v8EventListener[T]) HandleEvent(e *event.Event) error {
	f := l.val
	event, err := codec.EncodeEntity(l.ctx, e)
	if err == nil {
		global := l.ctx.GlobalThis()
		_, err1 := f.Call(global, event)
		err2 := l.ctx.Clock().Tick()
		err = errors.Join(err1, err2)
	}
	return err
}

func (l v8EventListener[T]) Equals(other event.EventHandler) bool {
	x, ok := other.(v8EventListener[T])
	return ok && x.val.StrictEquals(l.val)
}

func (w EventTarget[T]) CreateInstance(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	t := event.NewEventTarget()
	cbCtx.This().SetNativeValue(t)
	return nil, nil
}

func (w EventTarget[T]) decodeEventListener(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (event.EventHandler, error) {
	if fn, ok := val.AsFunction(); ok {
		return newV8EventListener(cbCtx, fn), nil
	} else {
		return nil, cbCtx.ValueFactory().NewTypeError("Must be a function")
	}
}

func (w EventTarget[T]) defaultEventListenerOptions() []event.EventListenerOption {
	return nil
}

func (w EventTarget[T]) decodeEventListenerOptions(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) ([]event.EventListenerOption, error) {
	var options []func(*event.EventListener)
	if val.IsBoolean() && val.Boolean() {
		options = append(options, event.Capture)
	}
	if obj, ok := val.AsObject(); ok {
		if capture, err := obj.Get("capture"); err == nil &&
			capture != nil {
			if capture.Boolean() {
				options = append(options, event.Capture)
			}
		}
	}
	return options, nil
}

func (w EventTarget[T]) decodeEvent(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (*event.Event, error) {
	obj, err := js.AssertObjectArg(cbCtx, val)
	if err == nil {
		return js.As[*event.Event](obj.NativeValue(), nil)
	}
	return nil, err
}
