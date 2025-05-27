// This file is generated. Do not edit.

package v8host

import (
	"errors"
	js "github.com/gost-dom/browser/scripting/internal/js"
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

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w pointerEventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {

	prototypeTmpl.SetAccessorProperty("width",
		wrapV8Callback(w.scriptHost, w.width),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("height",
		wrapV8Callback(w.scriptHost, w.height),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("pressure",
		wrapV8Callback(w.scriptHost, w.pressure),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("tangentialPressure",
		wrapV8Callback(w.scriptHost, w.tangentialPressure),
		nil,
		v8.None)
}

func (w pointerEventV8Wrapper) Constructor(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: PointerEvent.Constructor")
	type_, err1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	eventInitDict, err2 := consumeArgument(cbCtx, "eventInitDict", nil, w.decodePointerEventInit)
	if cbCtx.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return cbCtx.ReturnWithError(err)
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	if cbCtx.noOfReadArguments >= 1 {
		if err1 != nil {
			return cbCtx.ReturnWithError(err1)
		}
		return w.CreateInstance(cbCtx, type_)
	}
	return cbCtx.ReturnWithError(errors.New("PointerEvent.constructor: Missing arguments"))
}

func (w pointerEventV8Wrapper) width(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: PointerEvent.width")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.width: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventV8Wrapper) height(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: PointerEvent.height")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.height: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventV8Wrapper) pressure(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: PointerEvent.pressure")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.pressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w pointerEventV8Wrapper) tangentialPressure(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: PointerEvent.tangentialPressure")
	return cbCtx.ReturnWithError(errors.New("PointerEvent.tangentialPressure: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
