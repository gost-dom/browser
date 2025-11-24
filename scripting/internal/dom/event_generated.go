// This file is generated. Do not edit.

package dom

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
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
	cbCtx.Logger().Debug("JS Function call: Event.Constructor - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.Constructor", js.LogAttr("res", res))
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
	cbCtx.Logger().Debug("JS Function call: Event.stopPropagation - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.stopPropagation", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return nil, nil
}

func (w Event[T]) preventDefault(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Event.preventDefault - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.preventDefault", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return nil, nil
}

func (w Event[T]) type_(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Event.type_ - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.type_", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return codec.EncodeString(cbCtx, result)
}

func (w Event[T]) target(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Event.target - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.target", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return w.toEventTarget(cbCtx, result)
}

func (w Event[T]) currentTarget(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Event.currentTarget - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.currentTarget", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget
	return w.toEventTarget(cbCtx, result)
}

func (w Event[T]) bubbles(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Event.bubbles - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.bubbles", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Bubbles
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) cancelable(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Event.cancelable - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.cancelable", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Cancelable
	return codec.EncodeBoolean(cbCtx, result)
}

func (w Event[T]) defaultPrevented(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	cbCtx.Logger().Debug("JS Function call: Event.defaultPrevented - completed", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		cbCtx.Logger().Debug("JS Function call: Event.defaultPrevented", js.LogAttr("res", res))
	}()
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DefaultPrevented
	return codec.EncodeBoolean(cbCtx, result)
}
