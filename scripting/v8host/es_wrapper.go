package v8host

import (
	"errors"
	"runtime/cgo"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"

	v8 "github.com/gost-dom/v8go"
)

type converters struct{}

func (w converters) decodeEventInit(
	ctx *V8ScriptContext,
	v *v8.Value,
) (event.EventInit, error) {
	// var eventOptions []event.EventOption
	options, err0 := v.AsObject()

	bubbles, err1 := options.Get("bubbles")
	cancelable, err2 := options.Get("cancelable")
	err := errors.Join(err0, err1, err2)
	if err != nil {
		return event.EventInit{}, err
	}
	init := event.EventInit{
		Bubbles:    bubbles.Boolean(),
		Cancelable: cancelable.Boolean(),
	}
	// err := errors.Join(err0, err1, err2)
	// if err == nil {
	// 	eventOptions = []event.EventOption{
	// 		event.EventBubbles(bubbles.Boolean()),
	// 		event.EventCancelable(cancelable.Boolean()),
	// 	}
	// }
	// return event.EventOptions(eventOptions...), nil
	return init, nil
}

func (w converters) decodeUSVString(ctx *V8ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeByteString(ctx *V8ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeDOMString(ctx *V8ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeBoolean(ctx *V8ScriptContext, val *v8.Value) (bool, error) {
	return val.Boolean(), nil
}

func (w converters) decodeLong(ctx *V8ScriptContext, val *v8.Value) (int, error) {
	return int(val.Int32()), nil
}

func (w converters) decodeUnsignedLong(ctx *V8ScriptContext, val *v8.Value) (int, error) {
	return int(val.Uint32()), nil
}

func (w converters) decodeNode(ctx *V8ScriptContext, val *v8.Value) (dom.Node, error) {
	if val.IsObject() {
		o := val.Object()
		cached, ok_1 := ctx.getCachedNode(o)
		if node, ok_2 := cached.(dom.Node); ok_1 && ok_2 {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(ctx.host.iso, "Must be a node")
}

func (w converters) decodeHTMLElement(
	ctx *V8ScriptContext,
	val *v8.Value,
) (html.HTMLElement, error) {
	if val.IsObject() {
		o := val.Object()
		cached, ok_1 := ctx.getCachedNode(o)
		if node, ok_2 := cached.(html.HTMLElement); ok_1 && ok_2 {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(ctx.host.iso, "Must be a node")
}
func (c converters) defaultHTMLElement() html.HTMLElement { return nil }

func (w converters) decodeNodeOrText(ctx *V8ScriptContext, val *v8.Value) (dom.Node, error) {
	if val.IsString() {
		return ctx.window.Document().CreateText(val.String()), nil
	}
	return w.decodeNode(ctx, val)
}

func (w converters) toNullableByteString(ctx *V8ScriptContext, str *string) (*v8.Value, error) {
	if str == nil {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, *str)
}

func (w converters) toByteString(ctx *V8ScriptContext, str string) (*v8.Value, error) {
	if str == "" {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toDOMString(ctx *V8ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toNullableDOMString(ctx *V8ScriptContext, str *string) (*v8.Value, error) {
	if str == nil {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toUnsignedLong(ctx *V8ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, uint32(val))
}

func (w converters) toAny(ctx *V8ScriptContext, val string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w converters) toUSVString(ctx *V8ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toUnsignedShort(ctx *V8ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, uint32(val))
}

func (w converters) toBoolean(ctx *V8ScriptContext, val bool) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w converters) toNodeList(ctx *V8ScriptContext, val NodeList) (*v8.Value, error) {
	return ctx.getInstanceForNodeByName("NodeList", val)
}

func (w converters) toHTMLFormControlsCollection(
	ctx *V8ScriptContext,
	val NodeList,
) (*v8.Value, error) {
	return w.toNodeList(ctx, val)
}

// handleReffedObject serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type handleReffedObject[T any] struct {
	scriptHost *V8ScriptHost
	converters
}

func (o handleReffedObject[T]) mustGetContext(info *v8.FunctionCallbackInfo) *V8ScriptContext {
	return o.scriptHost.mustGetContext(info.Context())
}

func newHandleReffedObject[T any](host *V8ScriptHost) handleReffedObject[T] {
	return handleReffedObject[T]{
		scriptHost: host,
	}
}

func (o handleReffedObject[T]) iso() *v8.Isolate { return o.scriptHost.iso }

func (o handleReffedObject[T]) store(
	value any,
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	handle := cgo.NewHandle(value)
	ctx.addDisposer(handleDisposable(handle))

	e, ok := value.(entity.Entity)
	if ok {
		objectId := e.ObjectId()
		ctx.v8nodes[objectId] = this.Value
		ctx.domNodes[objectId] = e
	}

	internalField := v8.NewValueExternalHandle(o.scriptHost.iso, handle)
	this.SetInternalField(0, internalField)
	return this.Value, nil
}

func getWrappedInstance[T any](object *v8.Object) (res T, err error) {
	field := object.GetInternalField(0)
	handle := field.ExternalHandle()
	var ok bool
	res, ok = handle.Value().(T)
	if !ok {
		panic("Foo")
		err = errors.New("Not a valid type stored in the handle")
	}
	return
}

func (o handleReffedObject[T]) getInstance(info *v8.FunctionCallbackInfo) (T, error) {
	return getWrappedInstance[T](info.This())
}
