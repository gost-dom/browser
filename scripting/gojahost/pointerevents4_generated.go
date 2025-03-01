// This file is generated. Do not edit.

package gojahost

import g "github.com/dop251/goja"

func init() {
	installClass("PointerEvent", "MouseEvent", newPointerEventWrapper)
}

func (w pointerEventWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("getCoalescedEvents", w.getCoalescedEvents)
	prototype.Set("getPredictedEvents", w.getPredictedEvents)
	prototype.DefineAccessorProperty("pointerId", w.ctx.vm.ToValue(w.pointerId), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("width", w.ctx.vm.ToValue(w.width), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("height", w.ctx.vm.ToValue(w.height), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("pressure", w.ctx.vm.ToValue(w.pressure), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("tangentialPressure", w.ctx.vm.ToValue(w.tangentialPressure), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("tiltX", w.ctx.vm.ToValue(w.tiltX), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("tiltY", w.ctx.vm.ToValue(w.tiltY), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("twist", w.ctx.vm.ToValue(w.twist), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("altitudeAngle", w.ctx.vm.ToValue(w.altitudeAngle), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("azimuthAngle", w.ctx.vm.ToValue(w.azimuthAngle), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("pointerType", w.ctx.vm.ToValue(w.pointerType), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("isPrimary", w.ctx.vm.ToValue(w.isPrimary), nil, g.FLAG_TRUE, g.FLAG_TRUE)
	prototype.DefineAccessorProperty("persistentDeviceId", w.ctx.vm.ToValue(w.persistentDeviceId), nil, g.FLAG_TRUE, g.FLAG_TRUE)
}

func (w pointerEventWrapper) getCoalescedEvents(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.GetCoalescedEvents()
	return w.to(result)
}

func (w pointerEventWrapper) getPredictedEvents(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.GetPredictedEvents()
	return w.to(result)
}

func (w pointerEventWrapper) pointerId(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.PointerId()
	return w.toLong(result)
}

func (w pointerEventWrapper) width(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Width()
	return w.toDouble(result)
}

func (w pointerEventWrapper) height(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Height()
	return w.toDouble(result)
}

func (w pointerEventWrapper) pressure(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Pressure()
	return w.toFloat(result)
}

func (w pointerEventWrapper) tangentialPressure(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.TangentialPressure()
	return w.toFloat(result)
}

func (w pointerEventWrapper) tiltX(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.TiltX()
	return w.toLong(result)
}

func (w pointerEventWrapper) tiltY(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.TiltY()
	return w.toLong(result)
}

func (w pointerEventWrapper) twist(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.Twist()
	return w.toLong(result)
}

func (w pointerEventWrapper) altitudeAngle(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.AltitudeAngle()
	return w.toDouble(result)
}

func (w pointerEventWrapper) azimuthAngle(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.AzimuthAngle()
	return w.toDouble(result)
}

func (w pointerEventWrapper) pointerType(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.PointerType()
	return w.toDOMString(result)
}

func (w pointerEventWrapper) isPrimary(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.IsPrimary()
	return w.toBoolean(result)
}

func (w pointerEventWrapper) persistentDeviceId(c g.FunctionCall) g.Value {
	instance := w.getInstance(c)
	result := instance.PersistentDeviceId()
	return w.toLong(result)
}
