// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("Event", "", createEventPrototype)
}

func createEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newEventV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w eventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("stopPropagation", v8.NewFunctionTemplateWithError(iso, w.stopPropagation))
	prototypeTmpl.Set("preventDefault", v8.NewFunctionTemplateWithError(iso, w.preventDefault))

	prototypeTmpl.SetAccessorProperty("type",
		v8.NewFunctionTemplateWithError(iso, w.type_),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, w.target),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("currentTarget",
		v8.NewFunctionTemplateWithError(iso, w.currentTarget),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("eventPhase",
		v8.NewFunctionTemplateWithError(iso, w.eventPhase),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("bubbles",
		v8.NewFunctionTemplateWithError(iso, w.bubbles),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("cancelable",
		v8.NewFunctionTemplateWithError(iso, w.cancelable),
		nil,
		v8.None)
}

func (w eventV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(w.scriptHost, info)
	type_, err1 := tryParseArg(args, 0, w.decodeDOMString)
	eventInitDict, err2 := tryParseArgWithDefault(args, 1, w.defaultEventInit, w.decodeEventInit)
	ctx := w.mustGetContext(info)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstance(ctx, info.This(), type_, eventInitDict)
	}
	return nil, errors.New("Event.constructor: Missing arguments")
}

func (w eventV8Wrapper) stopPropagation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Event.stopPropagation")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	instance.StopPropagation()
	return nil, nil
}

func (w eventV8Wrapper) preventDefault(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Event.preventDefault")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	instance.PreventDefault()
	return nil, nil
}

func (w eventV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Event.target")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return w.toNullableEventTarget(ctx, result)
}

func (w eventV8Wrapper) currentTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Event.currentTarget")
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.CurrentTarget()
	return w.toNullableEventTarget(ctx, result)
}
