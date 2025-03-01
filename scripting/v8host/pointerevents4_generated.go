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
	prototypeTmpl.Set("getCoalescedEvents", v8.NewFunctionTemplateWithError(iso, w.getCoalescedEvents))
	prototypeTmpl.Set("getPredictedEvents", v8.NewFunctionTemplateWithError(iso, w.getPredictedEvents))

	prototypeTmpl.SetAccessorProperty("pointerId",
		v8.NewFunctionTemplateWithError(iso, w.pointerId),
		nil,
		v8.None)
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
	prototypeTmpl.SetAccessorProperty("tiltX",
		v8.NewFunctionTemplateWithError(iso, w.tiltX),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("tiltY",
		v8.NewFunctionTemplateWithError(iso, w.tiltY),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("twist",
		v8.NewFunctionTemplateWithError(iso, w.twist),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("altitudeAngle",
		v8.NewFunctionTemplateWithError(iso, w.altitudeAngle),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("azimuthAngle",
		v8.NewFunctionTemplateWithError(iso, w.azimuthAngle),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("pointerType",
		v8.NewFunctionTemplateWithError(iso, w.pointerType),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("isPrimary",
		v8.NewFunctionTemplateWithError(iso, w.isPrimary),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("persistentDeviceId",
		v8.NewFunctionTemplateWithError(iso, w.persistentDeviceId),
		nil,
		v8.None)
}

func (w pointerEventV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(w.scriptHost, info)
	type_, err1 := tryParseArg(args, 0, w.decodeDOMString)
	eventInitDict, err2 := tryParseArg(args, 1, w.decodePointerEventInit)
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
	return nil, errors.New("PointerEvent.constructor: Missing arguments")
}

func (w pointerEventV8Wrapper) getCoalescedEvents(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.getCoalescedEvents")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetCoalescedEvents()
	return w.to(ctx, result)
}

func (w pointerEventV8Wrapper) getPredictedEvents(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.getPredictedEvents")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetPredictedEvents()
	return w.to(ctx, result)
}

func (w pointerEventV8Wrapper) pointerId(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.pointerId")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.PointerId()
	return w.toLong(ctx, result)
}

func (w pointerEventV8Wrapper) width(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.width")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Width()
	return w.toDouble(ctx, result)
}

func (w pointerEventV8Wrapper) height(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.height")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Height()
	return w.toDouble(ctx, result)
}

func (w pointerEventV8Wrapper) pressure(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.pressure")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Pressure()
	return w.toFloat(ctx, result)
}

func (w pointerEventV8Wrapper) tangentialPressure(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.tangentialPressure")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.TangentialPressure()
	return w.toFloat(ctx, result)
}

func (w pointerEventV8Wrapper) tiltX(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.tiltX")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.TiltX()
	return w.toLong(ctx, result)
}

func (w pointerEventV8Wrapper) tiltY(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.tiltY")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.TiltY()
	return w.toLong(ctx, result)
}

func (w pointerEventV8Wrapper) twist(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.twist")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Twist()
	return w.toLong(ctx, result)
}

func (w pointerEventV8Wrapper) altitudeAngle(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.altitudeAngle")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.AltitudeAngle()
	return w.toDouble(ctx, result)
}

func (w pointerEventV8Wrapper) azimuthAngle(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.azimuthAngle")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.AzimuthAngle()
	return w.toDouble(ctx, result)
}

func (w pointerEventV8Wrapper) pointerType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.pointerType")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.PointerType()
	return w.toDOMString(ctx, result)
}

func (w pointerEventV8Wrapper) isPrimary(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.isPrimary")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.IsPrimary()
	return w.toBoolean(ctx, result)
}

func (w pointerEventV8Wrapper) persistentDeviceId(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: PointerEvent.persistentDeviceId")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.PersistentDeviceId()
	return w.toLong(ctx, result)
}
