package dom

import (
	"github.com/gost-dom/browser/dom/event"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type eventListener[T any] struct {
	s   js.Scope[T]
	val js.Function[T]
}

func newEventListener[T any](s js.Scope[T], val js.Function[T]) event.EventHandler {
	return eventListener[T]{s, val}
}

func (l eventListener[T]) HandleEvent(e *event.Event) error {
	f := l.val
	event, err := codec.EncodeEntity(l.s, e)
	if err == nil {
		global := l.s.GlobalThis()
		var res js.Value[T]
		err = l.s.Clock().Do(func() error {
			res, err = f.Call(global, event)
			return err
		})
		cancel := js.IsBoolean(res) && !res.Boolean() || err != nil
		if cancel {
			e.PreventDefault()
		}
	}
	return err
}

func (l eventListener[T]) Equals(other event.EventHandler) bool {
	x, ok := other.(eventListener[T])
	return ok && x.val.StrictEquals(l.val)
}

func CreateEventTarget[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	t := event.NewEventTarget()
	cbCtx.This().SetNativeValue(t)
	return nil, nil
}

func decodeEventListener[T any](
	s js.Scope[T], v js.Value[T],
) (event.EventHandler, error) {
	if fn, ok := v.AsFunction(); ok {
		return newEventListener(s, fn), nil
	} else {
		return nil, s.NewTypeError("Must be a function")
	}
}

func decodeEventListenerOptions[T any](
	scope js.Scope[T], val js.Value[T],
) (res []event.EventListenerOption, err error) {
	isObj := val.IsObject()
	if isObj {
		if res, err = codec.DecodeOptions(scope, val, codec.Options[T, event.EventListenerOption]{
			"capture": codec.OptDecoder[T](codec.DecodeBoolean, event.WithCapture),
			"once":    codec.OptDecoder[T](codec.DecodeBoolean, event.WithOnce),
		}); err == nil {
			return
		}
	}
	if val.IsBoolean() && val.Boolean() {
		return []event.EventListenerOption{event.Capture}, nil
	}
	return
}

func decodeEvent[T any](s js.Scope[T], v js.Value[T]) (*event.Event, error) {
	obj, err := js.AssertObjectArg(s, v)
	if err == nil {
		return js.As[*event.Event](obj.NativeValue(), nil)
	}
	return nil, err
}
