package v8host

import (
	"errors"

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

func (w converters) decodeEventInit(ctx jsCallbackContext, v jsValue) (eventInitWrapper, error) {
	options, ok := v.AsObject()
	if !ok {
		return eventInitWrapper{}, errors.New("Not an event init object")
	}

	bubbles, err1 := options.Get("bubbles")
	cancelable, err2 := options.Get("cancelable")
	err := errors.Join(err1, err2)
	if err != nil {
		return eventInitWrapper{}, err
	}
	init := eventInitWrapper{
		bubbles:    bubbles.Boolean(),
		cancelable: cancelable.Boolean(),
	}
	return init, nil
}

func (w converters) decodeString(ctx jsCallbackContext, val jsValue) (string, error) {
	return val.String(), nil
}

func (w converters) decodeBoolean(ctx jsCallbackContext, val jsValue) (bool, error) {
	return val.Boolean(), nil
}

func (w converters) decodeLong(ctx jsCallbackContext, val jsValue) (int, error) {
	return int(val.Int32()), nil
}

func (w converters) decodeUnsignedLong(ctx jsCallbackContext, val jsValue) (int, error) {
	return int(val.Uint32()), nil
}

func (w converters) decodeNode(ctx jsCallbackContext, val jsValue) (dom.Node, error) {
	if obj, ok := val.AsObject(); ok {
		if node, ok := obj.NativeValue().(dom.Node); ok {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(ctx.host.iso, "Must be a node")
}

func (w converters) decodeHTMLElement(
	cbCtx jsCallbackContext,
	val jsValue,
) (html.HTMLElement, error) {
	if o, ok := val.AsObject(); ok {
		if node, ok := o.NativeValue().(html.HTMLElement); ok {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(cbCtx.iso(), "Must be a node")
}
func (w converters) decodeHTMLFormElement(
	ctx jsCallbackContext,
	val jsValue,
) (html.HTMLFormElement, error) {
	var (
		res html.HTMLFormElement
		ok  bool
	)
	node, err := w.decodeNode(ctx, val)
	if err == nil {
		res, ok = node.(html.HTMLFormElement)
		if !ok {
			err = v8.NewTypeError(ctx.host.iso, "Not a form")
		}
	}
	return res, err
}

func (c converters) defaultHTMLElement() html.HTMLElement { return nil }

func (w converters) decodeNodeOrText(cbCtx jsCallbackContext, val jsValue) (dom.Node, error) {
	if val.IsString() {
		return cbCtx.ScriptCtx().window.Document().CreateTextNode(val.String()), nil
	}
	return w.decodeNode(cbCtx, val)
}

func (w converters) toNull(cbCtx jsCallbackContext) (jsValue, error) {
	return cbCtx.ValueFactory().Null(), nil
}

func (w converters) toNullableString_(
	cbCtx jsCallbackContext,
	str *string,
) (jsValue, error) {
	if str == nil {
		return w.toNull(cbCtx)
	}
	return w.toString_(cbCtx, *str)
}

func (w converters) toNillableString_(
	cbCtx jsCallbackContext,
	str string,
	hasVal bool,
) (jsValue, error) {
	if !hasVal {
		return w.toNull(cbCtx)
	}
	return w.toString_(cbCtx, str)
}

func (w converters) toUnsignedLong(cbCtx jsCallbackContext, val int) (jsValue, error) {
	return cbCtx.ValueFactory().NewUint32(uint32(val)), nil
}

func (w converters) toUnsignedShort(cbCtx jsCallbackContext, val int) (jsValue, error) {
	// TODO: This should be uint16 - but v8go doesn't support uint16
	return cbCtx.ValueFactory().NewUint32(uint32(val)), nil
}

func (w converters) toLong(cbCtx jsCallbackContext, val int) (jsValue, error) {
	return cbCtx.ValueFactory().NewInt64(int64(val)), nil
}

func (w converters) toAny(cbCtx jsCallbackContext, val string) (jsValue, error) {
	return w.toString_(cbCtx, val)
}

func (w converters) toString_(cbCtx jsCallbackContext, val string) (jsValue, error) {
	return cbCtx.ValueFactory().NewString(val), nil
}

func (w converters) toBoolean(cbCtx jsCallbackContext, val bool) (jsValue, error) {
	return cbCtx.ValueFactory().NewBoolean(val), nil
}

func (w converters) toJSWrapper(cbCtx jsCallbackContext, val entity.ObjectIder) (jsValue, error) {
	return cbCtx.ReturnWithJSValueErr(cbCtx.ScriptCtx().getJSInstance(val))
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

// TODO: Return js.CallbackRVal
func (o handleReffedObject[T]) store(
	value any,
	ctx *V8ScriptContext,
	this jsObject,
) (jsValue, error) {
	ctx.cacheNode(this, value)

	this.SetNativeValue(value)
	if d, ok := this.(disposable); ok {
		ctx.addDisposer(d)
	}
	return this, nil
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
