package v8host

import (
	"errors"
	"fmt"
	"runtime/cgo"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/v8go"
)

type customEventV8Wrapper struct {
	handleReffedObject[*event.Event]
}

func newCustomEventV8Wrapper(scriptHost *V8ScriptHost) *customEventV8Wrapper {
	return &customEventV8Wrapper{newHandleReffedObject[*event.Event](scriptHost)}
}

func createCustomEvent(host *V8ScriptHost) *v8go.FunctionTemplate {
	iso := host.iso
	wrapper := newCustomEventV8Wrapper(host)

	res := v8go.NewFunctionTemplateWithError(iso, wrapper.constructor)
	res.InstanceTemplate().SetInternalFieldCount(1)
	res.PrototypeTemplate().
		SetAccessorProperty("detail", v8go.NewFunctionTemplateWithError(iso, wrapper.detail), nil, v8go.None)
	return res
}
func (w customEventV8Wrapper) constructor(info *v8go.FunctionCallbackInfo) (*v8go.Value, error) {
	host := w.scriptHost
	iso := host.iso
	ctx := host.mustGetContext(info.Context())
	args := info.Args()
	if len(args) < 1 {
		return nil, v8go.NewTypeError(iso, "Must have at least one constructor argument")
	}
	data := event.CustomEventInit{}
	e := &event.Event{
		Type: args[0].String(),
	}
	if len(args) > 1 {
		if options, err := args[1].AsObject(); err == nil {
			bubbles, err1 := options.Get("bubbles")
			cancelable, err2 := options.Get("cancelable")
			detail, err3 := options.Get("detail")
			err = errors.Join(err1, err2, err3)
			if err != nil {
				return nil, err
			}
			e.Bubbles = bubbles.Boolean()
			e.Cancelable = cancelable.Boolean()
			data.Detail = detail
		}
	}
	e.Data = data
	handle := cgo.NewHandle(e)
	ctx.addDisposer(handleDisposable(handle))
	info.This().SetInternalField(0, v8go.NewValueExternalHandle(iso, handle))
	return v8go.Undefined(iso), nil
}

func (w customEventV8Wrapper) detail(info *v8go.FunctionCallbackInfo) (*v8go.Value, error) {
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	if data, ok := instance.Data.(event.CustomEventInit); ok {
		detail, _ := data.Detail.(*v8go.Value)
		return detail, nil
	}
	return nil, fmt.Errorf(
		"Data for custom event was not a CustomEventInit. %s",
		constants.BUG_USSUE_URL,
	)
}
