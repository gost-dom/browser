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
}

func (w mouseEventV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(w.scriptHost, info)
	type_, err1 := tryParseArg(args, 0, w.decodeDOMString)
	eventInitDict, err2 := tryParseArg(args, 1, w.decodeMouseEventInit)
	ctx := w.mustGetContext(info)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstanceEventInitDict(ctx, info.This(), type_, eventInitDict)
	}
	if args.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(ctx, info.This(), type_)
	}
	return nil, errors.New("MouseEvent.constructor: Missing arguments")
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
	args := newArgumentHelper(w.scriptHost, info)
	type_, err1 := tryParseArg(args, 0, w.decodeDOMString)
	eventInitDict, err2 := tryParseArg(args, 1, w.decodeUIEventInit)
	ctx := w.mustGetContext(info)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return w.CreateInstanceEventInitDict(ctx, info.This(), type_, eventInitDict)
	}
	if args.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
		}
		return w.CreateInstance(ctx, info.This(), type_)
	}
	return nil, errors.New("UIEvent.constructor: Missing arguments")
}

func (w uIEventV8Wrapper) view(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: UIEvent.view")
	return nil, errors.New("UIEvent.view: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w uIEventV8Wrapper) detail(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: UIEvent.detail")
	return nil, errors.New("UIEvent.detail: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
