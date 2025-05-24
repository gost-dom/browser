// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("PointerEvent", "MouseEvent", createPointerEventPrototype)
}

func createPointerEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newPointerEventV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w pointerEventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("width",
		v8.NewFunctionTemplateWithError(iso, w.width),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("height",
		v8.NewFunctionTemplateWithError(iso, w.height),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("pressure",
		v8.NewFunctionTemplateWithError(iso, w.pressure),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("tangentialPressure",
		v8.NewFunctionTemplateWithError(iso, w.tangentialPressure),
		nil,
		v8.None)
}

func (w pointerEventV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	cbCtx := newArgumentHelper(w.scriptHost, info)
	type_, err1 := tryParseArg(args, 0, w.decodeString)
	eventInitDict, err2 := tryParseArg(args, 1, w.decodePointerEventInit)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstanceEventInitDict(cbCtx.Context(), info.This(), type_, eventInitDict)
	}
	if args.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(cbCtx.Context(), info.This(), type_)
	}
	return nil, errors.New("PointerEvent.constructor: Missing arguments")
}

func (w pointerEventV8Wrapper) width(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: PointerEvent.width")
	return nil, errors.New("PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventV8Wrapper) height(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: PointerEvent.height")
	return nil, errors.New("PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventV8Wrapper) pressure(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: PointerEvent.pressure")
	return nil, errors.New("PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w pointerEventV8Wrapper) tangentialPressure(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: PointerEvent.tangentialPressure")
	return nil, errors.New("PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
