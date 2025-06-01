package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/js"

	v8 "github.com/gost-dom/v8go"
)

type converters[T any] struct{}

type eventInitWrapper struct {
	bubbles    bool
	cancelable bool
	init       any
}

func (w converters[T]) decodeEventInit(
	ctx js.CallbackContext[T],
	val js.Value[T],
) (eventInitWrapper, error) {
	options, ok := val.AsObject()
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

func (w converters[T]) decodeString(ctx js.CallbackContext[T], val js.Value[T]) (string, error) {
	return val.String(), nil
}

func (w converters[T]) decodeBoolean(ctx js.CallbackContext[T], val js.Value[T]) (bool, error) {
	return val.Boolean(), nil
}

func (w converters[T]) decodeLong(ctx js.CallbackContext[T], val js.Value[T]) (int, error) {
	return int(val.Int32()), nil
}

func (w converters[T]) decodeUnsignedLong(ctx js.CallbackContext[T], val js.Value[T]) (int, error) {
	return int(val.Uint32()), nil
}

func (w converters[T]) decodeNode(ctx js.CallbackContext[T], val js.Value[T]) (dom.Node, error) {
	if obj, ok := val.AsObject(); ok {
		if node, ok := obj.NativeValue().(dom.Node); ok {
			return node, nil
		}
	}
	return nil, ctx.ValueFactory().NewTypeError("Value is not a node")
}

func (w converters[T]) decodeHTMLElement(
	cbCtx jsCallbackContext,
	val jsValue,
) (html.HTMLElement, error) {
	if o, ok := val.AsObject(); ok {
		if node, ok := o.NativeValue().(html.HTMLElement); ok {
			return node, nil
		}
	}
	return nil, cbCtx.ValueFactory().NewTypeError("Must be a node")
}

func (w converters[T]) decodeHTMLFormElement(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (html.HTMLFormElement, error) {
	var (
		res html.HTMLFormElement
		ok  bool
	)
	node, err := w.decodeNode(cbCtx, val)
	if err == nil {
		res, ok = node.(html.HTMLFormElement)
		if !ok {
			err = cbCtx.ValueFactory().NewTypeError("Not a form")
		}
	}
	return res, err
}

func (c converters[T]) defaultHTMLElement() html.HTMLElement { return nil }

func (w converters[T]) decodeNodeOrText(
	cbCtx js.CallbackContext[T],
	val js.Value[T],
) (dom.Node, error) {
	if val.IsString() {
		return cbCtx.Scope().Window().Document().CreateTextNode(val.String()), nil
	}
	return w.decodeNode(cbCtx, val)
}

func (w converters[T]) toNull(cbCtx jsCallbackContext) (jsValue, error) {
	return cbCtx.ValueFactory().Null(), nil
}

func (w converters[T]) toNullableString_(
	cbCtx jsCallbackContext,
	str *string,
) (jsValue, error) {
	if str == nil {
		return w.toNull(cbCtx)
	}
	return w.toString_(cbCtx, *str)
}

func (w converters[T]) toNillableString_(
	cbCtx jsCallbackContext,
	str string,
	hasVal bool,
) (jsValue, error) {
	if !hasVal {
		return w.toNull(cbCtx)
	}
	return w.toString_(cbCtx, str)
}

func (w converters[T]) toUnsignedLong(cbCtx jsCallbackContext, val int) (jsValue, error) {
	return cbCtx.ValueFactory().NewUint32(uint32(val)), nil
}

func (w converters[T]) toUnsignedShort(cbCtx jsCallbackContext, val int) (jsValue, error) {
	// TODO: This should be uint16 - but v8go doesn't support uint16
	return cbCtx.ValueFactory().NewUint32(uint32(val)), nil
}

func (w converters[T]) toLong(cbCtx jsCallbackContext, val int) (jsValue, error) {
	return cbCtx.ValueFactory().NewInt64(int64(val)), nil
}

func (w converters[T]) toAny(cbCtx jsCallbackContext, val string) (jsValue, error) {
	return w.toString_(cbCtx, val)
}

func (w converters[T]) toString_(cbCtx jsCallbackContext, val string) (jsValue, error) {
	return cbCtx.ValueFactory().NewString(val), nil
}

func (w converters[T]) toBoolean(cbCtx jsCallbackContext, val bool) (jsValue, error) {
	return cbCtx.ValueFactory().NewBoolean(val), nil
}

func (w converters[T]) toJSWrapper(
	cbCtx jsCallbackContext,
	val entity.ObjectIder,
) (jsValue, error) {
	return encodeEntity(cbCtx, val)
}

// handleReffedObject serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type handleReffedObject[T, U any] struct {
	scriptHost *V8ScriptHost
	converters[U]
}

func (w handleReffedObject[T, U]) logger(info *v8.FunctionCallbackInfo) log.Logger {
	ctx := w.mustGetContext(info)
	return ctx.window.Logger()
}

func (o handleReffedObject[T, U]) mustGetContext(info *v8.FunctionCallbackInfo) *V8ScriptContext {
	return o.scriptHost.mustGetContext(info.Context())
}

func newHandleReffedObject[T any](host *V8ScriptHost) handleReffedObject[T, jsTypeParam] {
	return handleReffedObject[T, jsTypeParam]{
		scriptHost: host,
	}
}

func (o handleReffedObject[T, U]) getInstance(info callbackInfo) (T, error) {
	return getWrappedInstance[T](info.This())
}

func (o handleReffedObject[T, U]) iso() *v8.Isolate { return o.scriptHost.iso }

// TODO: Return js.CallbackRVal
func (o handleReffedObject[T, U]) store(
	value any,
	cbCtx jsCallbackContext,
) (jsValue, error) {
	this := cbCtx.This()
	if e, ok := value.(entity.ObjectIder); ok {
		cbCtx.Scope().SetValue(e, this)
	}

	this.SetNativeValue(value)
	if d, ok := this.(js.Disposable); ok {
		cbCtx.Scope().AddDisposable(d)
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

func encodeInstance[T any](cbCtx js.CallbackContext[T], e entity.ObjectIder) (js.Value[T], error) {

	scope := cbCtx.Scope()
	if res, ok := cbCtx.Scope().GetValue(e); ok {
		return res, nil
	}
	prototypeName := lookupJSPrototype(e)
	ctor := cbCtx.Scope().Constructor(prototypeName)
	value, err := ctor.NewInstance(cbCtx.Scope(), e)
	if err == nil {
		scope.SetValue(e, value)
	}
	return value, err

}
