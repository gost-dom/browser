package v8host

import (
	"errors"
	"runtime/cgo"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"

	v8 "github.com/gost-dom/v8go"
)

type converters struct{}

type eventInitWrapper struct {
	bubbles    bool
	cancelable bool
	init       any
}

func (w converters) decodeEventInit(
	ctx *V8ScriptContext,
	v *v8.Value,
) (eventInitWrapper, error) {
	options, err0 := v.AsObject()

	bubbles, err1 := options.Get("bubbles")
	cancelable, err2 := options.Get("cancelable")
	err := errors.Join(err0, err1, err2)
	if err != nil {
		return eventInitWrapper{}, err
	}
	init := eventInitWrapper{
		bubbles:    bubbles.Boolean(),
		cancelable: cancelable.Boolean(),
	}
	return init, nil
}

func (w converters) decodeString(ctx *V8ScriptContext, val *v8.Value) (string, error) {
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

func (w converters) toNullableString_(
	cbCtx *argumentHelper,
	str *string,
) (*v8.Value, error) {
	if str == nil {
		return v8.Null(cbCtx.iso()), nil
	}
	return v8.NewValue(cbCtx.iso(), str)
}

func (w converters) toNillableString_(
	cbCtx *argumentHelper,
	str string,
	hasVal bool,
) (*v8.Value, error) {
	if !hasVal {
		return v8.Null(cbCtx.iso()), nil
	}
	return v8.NewValue(cbCtx.iso(), str)
}

func (w converters) toUnsignedLong(cbCtx *argumentHelper, val int) (*v8.Value, error) {
	return v8.NewValue(cbCtx.iso(), uint32(val))
}

func (w converters) toLong(cbCtx *argumentHelper, val int) (*v8.Value, error) {
	return v8.NewValue(cbCtx.iso(), int64(val))
}

func (w converters) toAny(cbCtx *argumentHelper, val string) (*v8.Value, error) {
	return v8.NewValue(cbCtx.iso(), val)
}

func (w converters) toString_(cbCtx *argumentHelper, str string) (*v8.Value, error) {
	return v8.NewValue(cbCtx.iso(), str)
}

func (w converters) toUnsignedShort(cbCtx *argumentHelper, val int) (*v8.Value, error) {
	return v8.NewValue(cbCtx.iso(), uint32(val))
}

func (w converters) toBoolean(cbCtx *argumentHelper, val bool) (*v8.Value, error) {
	return v8.NewValue(cbCtx.iso(), val)
}

func (w converters) toNodeList(cbCtx *argumentHelper, val dom.NodeList) (*v8.Value, error) {
	return cbCtx.ScriptCtx().getInstanceForNodeByName("NodeList", val)
}

func (w converters) toHTMLFormControlsCollection(
	cbCtx *argumentHelper,
	val dom.NodeList,
) (*v8.Value, error) {
	return w.toNodeList(cbCtx, val)
}

// handleReffedObject serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type handleReffedObject[T any] struct {
	scriptHost *V8ScriptHost
	converters
}

func (w handleReffedObject[T]) logger(info *v8.FunctionCallbackInfo) log.Logger {
	ctx := w.mustGetContext(info)
	return ctx.window.Logger()
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

func storeObjectHandleInV8Instance(
	value any,
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	handle := cgo.NewHandle(value)
	ctx.addDisposer(handleDisposable(handle))

	e, ok := value.(entity.ObjectIder)
	if ok {
		objectId := e.ObjectId()
		ctx.v8nodes[objectId] = this.Value
	}

	internalField := v8.NewValueExternalHandle(ctx.v8ctx.Isolate(), handle)
	this.SetInternalField(0, internalField)
	return this.Value, nil
}

func (o handleReffedObject[T]) store(
	value any,
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	return storeObjectHandleInV8Instance(value, ctx, this)
}

func getWrappedInstance[T any](object *v8.Object) (res T, err error) {
	field := object.GetInternalField(0)
	handle := field.ExternalHandle()
	var ok bool
	res, ok = handle.Value().(T)
	if !ok {
		err = errors.New("Not a valid type stored in the handle")
	}
	return
}

type callbackInfo interface {
	This() *v8.Object
}

func (o handleReffedObject[T]) getInstance(info callbackInfo) (T, error) {
	return getWrappedInstance[T](info.This())
}
