// This file is generated. Do not edit.

package v8host

import (
	"errors"
	uievents "github.com/gost-dom/browser/internal/uievents"
	js "github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("MouseEvent", "UIEvent", createMouseEventPrototype)
}

func createMouseEventPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := newMouseEventV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w mouseEventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("getModifierState", wrapV8Callback(w.scriptHost, w.getModifierState))

	prototypeTmpl.SetAccessorProperty("screenX",
		wrapV8Callback(w.scriptHost, w.screenX),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("screenY",
		wrapV8Callback(w.scriptHost, w.screenY),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientX",
		wrapV8Callback(w.scriptHost, w.clientX),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientY",
		wrapV8Callback(w.scriptHost, w.clientY),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("layerX",
		wrapV8Callback(w.scriptHost, w.layerX),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("layerY",
		wrapV8Callback(w.scriptHost, w.layerY),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("relatedTarget",
		wrapV8Callback(w.scriptHost, w.relatedTarget),
		nil,
		v8.None)
}

func (w mouseEventV8Wrapper) Constructor(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.Constructor")
	type_, err1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	eventInitDict, err2 := consumeArgument(cbCtx, "eventInitDict", nil, w.decodeMouseEventInit)
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
	return cbCtx.ReturnWithError(errors.New("MouseEvent.constructor: Missing arguments"))
}

func (w mouseEventV8Wrapper) getModifierState(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.getModifierState")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) screenX(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.screenX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) screenY(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.screenY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) clientX(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.clientX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) clientY(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.clientY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) layerX(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.layerX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) layerY(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.layerY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) relatedTarget(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: MouseEvent.relatedTarget")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
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
	wrapper := newUIEventV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w uIEventV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {

	prototypeTmpl.SetAccessorProperty("view",
		wrapV8Callback(w.scriptHost, w.view),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("detail",
		wrapV8Callback(w.scriptHost, w.detail),
		nil,
		v8.None)
}

func (w uIEventV8Wrapper) Constructor(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: UIEvent.Constructor")
	type_, err1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	eventInitDict, err2 := consumeArgument(cbCtx, "eventInitDict", nil, w.decodeUIEventInit)
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
	return cbCtx.ReturnWithError(errors.New("UIEvent.constructor: Missing arguments"))
}

func (w uIEventV8Wrapper) view(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: UIEvent.view")
	return cbCtx.ReturnWithError(errors.New("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w uIEventV8Wrapper) detail(cbCtx *argumentHelper) (js.Value, error) {
	cbCtx.logger().Debug("V8 Function call: UIEvent.detail")
	return cbCtx.ReturnWithError(errors.New("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
