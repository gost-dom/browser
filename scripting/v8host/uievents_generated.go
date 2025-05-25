// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	uievents "github.com/gost-dom/browser/internal/uievents"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("MouseEvent", "UIEvent", createMouseEventPrototype)
}

func createMouseEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newMouseEventV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w mouseEventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("getModifierState", v8.NewFunctionTemplateWithError(iso, w.getModifierState))

	prototypeTmpl.SetAccessorProperty("screenX",
		v8.NewFunctionTemplateWithError(iso, w.screenX),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("screenY",
		v8.NewFunctionTemplateWithError(iso, w.screenY),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientX",
		v8.NewFunctionTemplateWithError(iso, w.clientX),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientY",
		v8.NewFunctionTemplateWithError(iso, w.clientY),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("layerX",
		v8.NewFunctionTemplateWithError(iso, w.layerX),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("layerY",
		v8.NewFunctionTemplateWithError(iso, w.layerY),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("relatedTarget",
		v8.NewFunctionTemplateWithError(iso, w.relatedTarget),
		nil,
		v8.None)
}

func (w mouseEventV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	type_, err1 := parseArgument(cbCtx, 0, nil, w.decodeString)
	eventInitDict, err2 := parseArgument(cbCtx, 1, nil, w.decodeMouseEventInit)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstanceEventInitDict(cbCtx.Context(), info.This(), type_, eventInitDict)
	}
	if cbCtx.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(cbCtx.Context(), info.This(), type_)
	}
	return nil, errors.New("MouseEvent.constructor: Missing arguments")
}

func (w mouseEventV8Wrapper) getModifierState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.getModifierState")
	return nil, errors.New("MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventV8Wrapper) screenX(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.screenX")
	return nil, errors.New("MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventV8Wrapper) screenY(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.screenY")
	return nil, errors.New("MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventV8Wrapper) clientX(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.clientX")
	return nil, errors.New("MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventV8Wrapper) clientY(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.clientY")
	return nil, errors.New("MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventV8Wrapper) layerX(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.layerX")
	return nil, errors.New("MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventV8Wrapper) layerY(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.layerY")
	return nil, errors.New("MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w mouseEventV8Wrapper) relatedTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: MouseEvent.relatedTarget")
	return nil, errors.New("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func init() {
	registerJSClass("UIEvent", "Event", createUIEventPrototype)
}

type uIEventV8Wrapper struct {
	handleReffedObject[uievents.UIEvent]
}

func newUIEventV8Wrapper(scriptHost *V8ScriptHost) *uIEventV8Wrapper {
	return &uIEventV8Wrapper{newHandleReffedObject[uievents.UIEvent](scriptHost)}
}

func createUIEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newUIEventV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w uIEventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("view",
		v8.NewFunctionTemplateWithError(iso, w.view),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("detail",
		v8.NewFunctionTemplateWithError(iso, w.detail),
		nil,
		v8.None)
}

func (w uIEventV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: UIEvent.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	type_, err1 := parseArgument(cbCtx, 0, nil, w.decodeString)
	eventInitDict, err2 := parseArgument(cbCtx, 1, nil, w.decodeUIEventInit)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstanceEventInitDict(cbCtx.Context(), info.This(), type_, eventInitDict)
	}
	if cbCtx.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(cbCtx.Context(), info.This(), type_)
	}
	return nil, errors.New("UIEvent.constructor: Missing arguments")
}

func (w uIEventV8Wrapper) view(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: UIEvent.view")
	return nil, errors.New("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w uIEventV8Wrapper) detail(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: UIEvent.detail")
	return nil, errors.New("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
