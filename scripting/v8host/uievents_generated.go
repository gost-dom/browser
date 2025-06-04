// This file is generated. Do not edit.

package v8host

import (
	"errors"
	uievents "github.com/gost-dom/browser/internal/uievents"
)

func init() {
	registerClass("MouseEvent", "UIEvent", newMouseEventV8Wrapper)
}

func createMouseEventPrototype(scriptHost *V8ScriptHost) jsClass {
	wrapper := newMouseEventV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper mouseEventV8Wrapper) initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
}

func (w mouseEventV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeMethod("getModifierState", w.getModifierState)
	jsClass.CreatePrototypeAttribute("screenX", w.screenX, nil)
	jsClass.CreatePrototypeAttribute("screenY", w.screenY, nil)
	jsClass.CreatePrototypeAttribute("clientX", w.clientX, nil)
	jsClass.CreatePrototypeAttribute("clientY", w.clientY, nil)
	jsClass.CreatePrototypeAttribute("layerX", w.layerX, nil)
	jsClass.CreatePrototypeAttribute("layerY", w.layerY, nil)
	jsClass.CreatePrototypeAttribute("relatedTarget", w.relatedTarget, nil)
}

func (w mouseEventV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.constructor")
	type_, errArg1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := consumeOptionalArg(cbCtx, "eventInitDict", w.decodeMouseEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w mouseEventV8Wrapper) getModifierState(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.getModifierState")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.getModifierState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) screenX(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) screenY(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.screenY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.screenY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) clientX(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) clientY(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.clientY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.clientY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) layerX(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerX")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerX: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) layerY(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.layerY")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.layerY: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w mouseEventV8Wrapper) relatedTarget(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: MouseEvent.relatedTarget")
	return cbCtx.ReturnWithError(errors.New("MouseEvent.relatedTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func init() {
	registerClass("UIEvent", "Event", newUIEventV8Wrapper)
}

type uIEventV8Wrapper struct {
	handleReffedObject[uievents.UIEvent, jsTypeParam]
}

func newUIEventV8Wrapper(scriptHost *V8ScriptHost) *uIEventV8Wrapper {
	return &uIEventV8Wrapper{newHandleReffedObject[uievents.UIEvent](scriptHost)}
}

func createUIEventPrototype(scriptHost *V8ScriptHost) jsClass {
	wrapper := newUIEventV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper uIEventV8Wrapper) initialize(jsClass jsClass) {
	wrapper.installPrototype(jsClass)
}

func (w uIEventV8Wrapper) installPrototype(jsClass jsClass) {
	jsClass.CreatePrototypeAttribute("view", w.view, nil)
	jsClass.CreatePrototypeAttribute("detail", w.detail, nil)
}

func (w uIEventV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.constructor")
	type_, errArg1 := consumeArgument(cbCtx, "type", nil, w.decodeString)
	if errArg1 != nil {
		return nil, errArg1
	}
	eventInitDict, found, errArg := consumeOptionalArg(cbCtx, "eventInitDict", w.decodeUIEventInit)
	if found {
		if errArg != nil {
			return nil, errArg
		}
		return w.CreateInstanceEventInitDict(cbCtx, type_, eventInitDict)
	}
	return w.CreateInstance(cbCtx, type_)
}

func (w uIEventV8Wrapper) view(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.view")
	return cbCtx.ReturnWithError(errors.New("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w uIEventV8Wrapper) detail(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: UIEvent.detail")
	return cbCtx.ReturnWithError(errors.New("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
