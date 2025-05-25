// This file is generated. Do not edit.

package v8host

import (
	"errors"
	event "github.com/gost-dom/browser/dom/event"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("Event", "", createEventPrototype)
}

func createEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newEventV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w eventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("stopPropagation", wrapV8Callback(w.scriptHost, w.stopPropagation))
	prototypeTmpl.Set("preventDefault", wrapV8Callback(w.scriptHost, w.preventDefault))

	prototypeTmpl.SetAccessorProperty("type",
		wrapV8Callback(w.scriptHost, w.type_),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		wrapV8Callback(w.scriptHost, w.target),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("currentTarget",
		wrapV8Callback(w.scriptHost, w.currentTarget),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("eventPhase",
		wrapV8Callback(w.scriptHost, w.eventPhase),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("bubbles",
		wrapV8Callback(w.scriptHost, w.bubbles),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("cancelable",
		wrapV8Callback(w.scriptHost, w.cancelable),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("defaultPrevented",
		wrapV8Callback(w.scriptHost, w.defaultPrevented),
		nil,
		v8.None)
}

func (w eventV8Wrapper) Constructor(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.Constructor")
	type_, err1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	eventInitDict, err2 := consumeArgument(cbCtx, "eventInitDict", w.defaultEventInit, w.decodeEventInit)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		return w.CreateInstance(cbCtx, type_, eventInitDict)
	}
	return cbCtx.ReturnWithError(errors.New("Event.constructor: Missing arguments"))
}

func (w eventV8Wrapper) stopPropagation(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.stopPropagation")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return cbCtx.ReturnWithValue(nil)
}

func (w eventV8Wrapper) preventDefault(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.preventDefault")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return cbCtx.ReturnWithValue(nil)
}

func (w eventV8Wrapper) type_(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.type_")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Type
	return w.toString_(cbCtx, result)
}

func (w eventV8Wrapper) target(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.target")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target
	return w.toEventTarget(cbCtx, result)
}

func (w eventV8Wrapper) currentTarget(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.currentTarget")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget
	return w.toEventTarget(cbCtx, result)
}

func (w eventV8Wrapper) bubbles(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.bubbles")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Bubbles
	return w.toBoolean(cbCtx, result)
}

func (w eventV8Wrapper) cancelable(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.cancelable")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Cancelable
	return w.toBoolean(cbCtx, result)
}

func (w eventV8Wrapper) defaultPrevented(cbCtx *argumentHelper) (*v8.Value, error) {
	cbCtx.logger().Debug("V8 Function call: Event.defaultPrevented")
	instance, err := js.As[*event.Event](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.DefaultPrevented
	return w.toBoolean(cbCtx, result)
}
