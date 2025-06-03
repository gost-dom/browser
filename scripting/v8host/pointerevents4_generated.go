// This file is generated. Do not edit.

package v8host

import (
	"errors"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("PointerEvent", "MouseEvent", createPointerEventPrototype)
}

func createPointerEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newPointerEventV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor)

	return constructor
}

func (w pointerEventV8Wrapper) installPrototype(ft *v8.FunctionTemplate) {
	jsClass := newV8Class(w.scriptHost, ft)
	jsClass.CreatePrototypeAttribute("width", w.width, nil)
	jsClass.CreatePrototypeAttribute("height", w.height, nil)
	jsClass.CreatePrototypeAttribute("pressure", w.pressure, nil)
	jsClass.CreatePrototypeAttribute("tangentialPressure", w.tangentialPressure, nil)
}

func (w pointerEventV8Wrapper) Constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.Constructor")
	type_, errArg1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := consumeOptionalArg(cbCtx, "eventInitDict", w.decodePointerEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w pointerEventV8Wrapper) width(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.width")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventV8Wrapper) height(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.height")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventV8Wrapper) pressure(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.pressure")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventV8Wrapper) tangentialPressure(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: PointerEvent.tangentialPressure")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
