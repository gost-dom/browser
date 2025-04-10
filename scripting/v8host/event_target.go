package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom/event"
	v8 "github.com/gost-dom/v8go"
)

type v8EventListener struct {
	ctx *V8ScriptContext
	val *v8.Value
}

func newV8EventListener(ctx *V8ScriptContext, val *v8.Value) event.EventHandler {
	return v8EventListener{ctx, val}
}

func (l v8EventListener) HandleEvent(e *event.Event) error {
	f, err := l.val.AsFunction()
	if err == nil {
		var event *v8.Value
		event, err = l.ctx.getInstanceForNode(e)
		if err == nil {
			_, err1 := f.Call(l.val, event)
			err2 := l.ctx.eventLoop.tick()
			err = errors.Join(err1, err2)
		}
	}
	return err
}

func (l v8EventListener) Equals(other event.EventHandler) bool {
	x, ok := other.(v8EventListener)
	return ok && x.val.StrictEquals(l.val)
}

type eventTargetV8Wrapper struct {
	handleReffedObject[event.EventTarget]
}

func newEventTargetV8Wrapper(host *V8ScriptHost) eventTargetV8Wrapper {
	return eventTargetV8Wrapper{newHandleReffedObject[event.EventTarget](host)}
}

func createEventTarget(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newEventTargetV8Wrapper(host)
	res := v8.NewFunctionTemplate(
		iso,
		func(info *v8.FunctionCallbackInfo) *v8.Value {
			ctx := host.mustGetContext(info.Context())
			wrapper.store(event.NewEventTarget(), ctx, info.This())
			return v8.Undefined(iso)
		},
	)
	proto := res.PrototypeTemplate()
	proto.Set(
		"addEventListener",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				target, err := wrapper.getInstance(info)
				if err != nil {
					return nil, err
				}
				args := newArgumentHelper(host, info)
				eventType, e1 := args.getStringArg(0)
				fn, e2 := args.getFunctionArg(1)
				var options []func(*event.EventListener)
				optionArg := args.getArg(2)
				if optionArg != nil {
					if optionArg.IsBoolean() && optionArg.Boolean() {
						options = append(options, event.Capture)
					}
					if optionArg.IsObject() {
						if capture, err := optionArg.Object().Get("capture"); err == nil &&
							capture != nil {
							if capture.Boolean() {
								options = append(options, event.Capture)
							}
						}
					}
				}
				err = errors.Join(e1, e2)
				if err == nil {
					listener := newV8EventListener(ctx, fn.Value)
					target.AddEventListener(eventType, listener, options...)
				}
				return v8.Undefined(iso), err
			}), v8.ReadOnly)
	proto.Set(
		"removeEventListener",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				target, err := wrapper.getInstance(info)
				if err != nil {
					return nil, err
				}
				args := newArgumentHelper(host, info)
				eventType, e1 := args.getStringArg(0)
				fn, e2 := args.getFunctionArg(1)
				err = errors.Join(e1, e2)
				if err == nil {
					listener := newV8EventListener(ctx, fn.Value)
					target.RemoveEventListener(eventType, listener)
				}
				return v8.Undefined(iso), err
			}), v8.ReadOnly)
	proto.Set(
		"dispatchEvent",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				target, err := wrapper.getInstance(info)
				if err != nil {
					return nil, err
				}
				e := info.Args()[0]
				handle := e.Object().GetInternalField(0).ExternalHandle()
				if evt, ok := handle.Value().(*event.Event); ok {
					return v8.NewValue(iso, target.DispatchEvent(evt))
				} else {
					return nil, v8.NewTypeError(iso, "Not an Event")
				}
			}), v8.ReadOnly)
	instanceTemplate := res.InstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	return res
}
