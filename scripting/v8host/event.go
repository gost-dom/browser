package v8host

import (
	"github.com/gost-dom/browser/dom/events"
	"github.com/gost-dom/browser/internal/entity"
	v8 "github.com/gost-dom/v8go"
)

func (w eventV8Wrapper) defaultEventInit() events.EventOption {
	return events.EventOptions(nil)
}

func (w eventV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
	type_ string,
	o events.EventOption,
) (*v8.Value, error) {
	e := events.NewEvent(type_, o)
	return w.store(e, ctx, this)
}

func (w eventV8Wrapper) toNullableEventTarget(
	ctx *V8ScriptContext,
	e events.EventTarget,
) (*v8.Value, error) {
	if e == nil {
		return v8.Null(w.scriptHost.iso), nil
	}
	if entity, ok := e.(entity.Entity); ok {
		return ctx.getInstanceForNode(entity)
	}
	return nil, v8.NewError(w.iso(), "TODO, Not yet supported")
}

func (w eventV8Wrapper) eventPhase(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	return v8.NewValue(w.iso(), uint32(instance.EventPhase()))
}
