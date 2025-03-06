package v8host

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/entity"
	v8 "github.com/gost-dom/v8go"
)

// the eventWrapper type is a temporary solution because the code generator
// creates function calls, not field lookups.
//
// Once that if fixed, this type can be removed.
type eventWrapper struct{ *event.Event }

func (w eventWrapper) Type() string {
	return w.Event.Type
}
func (w eventWrapper) Cancelable() bool {
	return w.Event.Cancelable
}
func (w eventWrapper) Bubbles() bool {
	return w.Event.Bubbles
}

func (w eventV8Wrapper) defaultEventInit() eventInitWrapper {
	return eventInitWrapper{}
}

func (w eventV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
	type_ string,
	o eventInitWrapper,
) (*v8.Value, error) {
	e := eventWrapper{
		&event.Event{Type: type_, Bubbles: o.bubbles, Cancelable: o.cancelable, Init: o.init},
	}
	return w.store(e, ctx, this)
}

func (w eventV8Wrapper) toNullableEventTarget(
	ctx *V8ScriptContext,
	e event.EventTarget,
) (*v8.Value, error) {
	if e == nil {
		return v8.Null(w.scriptHost.iso), nil
	}
	if entity, ok := e.(entity.ObjectIder); ok {
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
