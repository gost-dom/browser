package dom

import (
	"errors"

	"github.com/gost-dom/browser/dom/event"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type v8EventListener[T any] struct {
	s   js.Scope[T]
	val js.Function[T]
}

func newV8EventListener[T any](s js.Scope[T], val js.Function[T]) event.EventHandler {
	return v8EventListener[T]{s, val}
}

func (l v8EventListener[T]) HandleEvent(e *event.Event) error {
	f := l.val
	event, err := codec.EncodeEntity(l.s, e)
	if err == nil {
		global := l.s.GlobalThis()
		_, err1 := f.Call(global, event)
		err2 := l.s.Clock().Tick()
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
	s js.Scope[T], v js.Value[T],
) (event.EventHandler, error) {
	if fn, ok := v.AsFunction(); ok {
		return newV8EventListener(s, fn), nil
	} else {
		return nil, s.NewTypeError("Must be a function")
	}
}

func (w EventTarget[T]) defaultEventListenerOptions() []event.EventListenerOption {
	return nil
}

func (w EventTarget[T]) decodeEventListenerOptions(
	_ js.Scope[T], v js.Value[T],
) ([]event.EventListenerOption, error) {
	var options []func(*event.EventListener)
	if v.IsBoolean() && v.Boolean() {
		options = append(options, event.Capture)
	}
	if obj, ok := v.AsObject(); ok {
		if capture, err := obj.Get("capture"); err == nil &&
			capture != nil {
			if capture.Boolean() {
				options = append(options, event.Capture)
			}
		}
	}
	return options, nil
}

func (w EventTarget[T]) decodeEvent(s js.Scope[T], v js.Value[T]) (*event.Event, error) {
	obj, err := js.AssertObjectArg(s, v)
	if err == nil {
		return js.As[*event.Event](obj.NativeValue(), nil)
	}
	return nil, err
}
