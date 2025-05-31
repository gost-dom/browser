package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type v8EventListener struct {
	ctx *V8ScriptContext
	val jsFunction
}

func newV8EventListener(ctx *V8ScriptContext, val jsFunction) event.EventHandler {
	return v8EventListener{ctx, val}
}

func (l v8EventListener) HandleEvent(e *event.Event) error {
	f := l.val
	event, err := l.ctx.getJSInstance(e)
	if err == nil {
		iso := l.ctx.host.iso
		global := l.ctx.v8ctx.Global()
		_, err1 := f.Call(newV8Object(iso, global), event)
		err2 := l.ctx.clock.Tick()
		err = errors.Join(err1, err2)
	}
	return err
}

func (l v8EventListener) Equals(other event.EventHandler) bool {
	x, ok := other.(v8EventListener)
	return ok && x.val.StrictEquals(l.val)
}

func (w eventTargetV8Wrapper) CreateInstance(cbCtx jsCallbackContext) (jsValue, error) {
	t := event.NewEventTarget()
	cbCtx.This().SetNativeValue(t)
	return nil, nil
}

func (w eventTargetV8Wrapper) decodeEventListener(
	cbCtx jsCallbackContext,
	val jsValue,
) (event.EventHandler, error) {
	if fn, ok := val.AsFunction(); ok {
		return newV8EventListener(cbCtx.ScriptCtx(), fn), nil
	} else {
		return nil, cbCtx.ValueFactory().NewTypeError("Must be a function")
	}
}

func (w eventTargetV8Wrapper) defaultEventListenerOptions() []event.EventListenerOption {
	return nil
}

func (w eventTargetV8Wrapper) decodeEventListenerOptions(
	cbCtx jsCallbackContext,
	val jsValue,
) ([]event.EventListenerOption, error) {
	var options []func(*event.EventListener)
	if val.IsBoolean() && val.Boolean() {
		options = append(options, event.Capture)
	}
	if obj, ok := val.AsObject(); ok {
		if capture, err := obj.Get("capture"); err == nil &&
			capture != nil {
			if capture.Boolean() {
				options = append(options, event.Capture)
			}
		}
	}
	return options, nil
}

func (w eventTargetV8Wrapper) decodeEvent(
	cbCtx jsCallbackContext,
	val jsValue,
) (*event.Event, error) {
	obj, err := js.AssertObjectArg(cbCtx, val)
	if err == nil {
		return js.As[*event.Event](obj.NativeValue(), nil)
	}
	return nil, err
}
