// This file is generated. Do not edit.

package dom

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
)

type Event[T any] struct{}

func NewEvent[T any](scriptHost js.ScriptEngine[T]) *Event[T] {
	return &Event[T]{}
}

func (wrapper Event[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Event[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreatePrototypeMethod("stopPropagation", w.stopPropagation)
	jsClass.CreatePrototypeMethod("preventDefault", w.preventDefault)
	jsClass.CreatePrototypeAttribute("type", w.type_, nil)
	jsClass.CreatePrototypeAttribute("target", w.target, nil)
	jsClass.CreatePrototypeAttribute("currentTarget", w.currentTarget, nil)
	jsClass.CreatePrototypeAttribute("eventPhase", w.eventPhase, nil)
	jsClass.CreatePrototypeAttribute("bubbles", w.bubbles, nil)
	jsClass.CreatePrototypeAttribute("cancelable", w.cancelable, nil)
	jsClass.CreatePrototypeAttribute("defaultPrevented", w.defaultPrevented, nil)
}

func (w Event[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	type_, errArg1 := js.ConsumeArgument(cbCtx, "type", nil, codec.DecodeString)
	eventInitDict, errArg2 := js.ConsumeArgument(cbCtx, "eventInitDict", codec.ZeroValue, codec.DecodeEventInit)
	err = errors.Join(errArg1, errArg2)
	if err != nil {
		return nil, err
	}
	return w.CreateInstance(cbCtx, type_, eventInitDict)
}

func (w Event[T]) stopPropagation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "stopPropagation"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return nil, nil
}

func (w Event[T]) preventDefault(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "preventDefault"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return nil, nil
}

func (w Event[T]) type_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "type_"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func (w Event[T]) target(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "target"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return w.toEventTarget(cbCtx, result)
}

func (w Event[T]) currentTarget(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "currentTarget"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget
	return w.toEventTarget(cbCtx, result)
}

func (w Event[T]) bubbles(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "bubbles"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Bubbles
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) cancelable(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "cancelable"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Cancelable
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) defaultPrevented(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Event"), slog.String("Method", "defaultPrevented"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DefaultPrevented
	return codec.EncodeBoolean(cbCtx, result)
}
