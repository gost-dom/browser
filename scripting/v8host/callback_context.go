package v8host

import (
	"errors"
	"fmt"
	"iter"
	"log/slog"
	"runtime/debug"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/v8go"
	v8 "github.com/gost-dom/v8go"
)

type jsCallbackContext = js.CallbackContext[*v8Value]

var (
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)

type v8CallbackInfo interface {
	Context() *v8go.Context
	This() *v8go.Object
}

type v8CallbackScope struct {
	v8Scope
	v8Info v8CallbackInfo
}

func newV8CallbackScope(host *V8ScriptHost, info v8CallbackInfo) v8CallbackScope {
	return v8CallbackScope{newV8Scope(host.mustGetContext(info.Context())), info}
}

func (h v8CallbackScope) This() jsObject {
	return newV8Object(h.ScriptCtx(), h.v8Info.This())
}

func (h v8CallbackScope) iso() *v8.Isolate { return h.ScriptCtx().host.iso }

func (h v8CallbackScope) ScriptCtx() *V8ScriptContext {
	return h.host.mustGetContext(h.v8Info.Context())
}

func (h v8CallbackScope) Instance() (any, error) {
	if h.v8Info.This().InternalFieldCount() >= 1 {
		handle := h.v8Info.This().GetInternalField(0).ExternalHandle()
		if handle != 0 {
			return handle.Value(), nil
		}
	}
	return nil, v8go.NewTypeError(h.iso(), "No internal instance")
}

type v8CallbackContext struct {
	v8CallbackScope
	v8Info       *v8.FunctionCallbackInfo
	host         *V8ScriptHost
	currentIndex int
}

func newCallbackContext(host *V8ScriptHost, info *v8.FunctionCallbackInfo) jsCallbackContext {
	return &v8CallbackContext{
		v8CallbackScope: newV8CallbackScope(host, info),
		v8Info:          info,
		host:            host,
	}
}

func (h *v8CallbackContext) ReturnWithTypeError(msg string) (jsValue, error) {
	return nil, v8.NewTypeError(h.iso(), msg)
}

func (h *v8CallbackContext) Args() []jsValue {
	ctx := h.host.mustGetContext(h.v8Info.Context())
	v8Args := h.v8Info.Args()
	res := make([]jsValue, len(v8Args))
	for i, arg := range v8Args {
		res[i] = newV8Value(ctx, arg)
	}
	return res
}

func (h *v8CallbackContext) ConsumeArg() (jsValue, bool) {
	index := h.currentIndex
	h.currentIndex++
	args := h.v8Info.Args()
	if len(args) <= index {
		return nil, false
	}
	arg := args[index]
	if arg.IsUndefined() {
		return nil, true
	}
	return h.toJSValue(arg), true
}

/* -------- v8Scope -------- */

type v8Scope struct {
	*V8ScriptContext
}

func newV8Scope(ctx *V8ScriptContext) v8Scope {
	return v8Scope{ctx}
}

func (s v8Scope) Logger() *slog.Logger { return s.host.Logger() }
func (s v8Scope) Window() html.Window  { return s.window }
func (s v8Scope) GlobalThis() jsObject { return s.global }
func (s v8Scope) Clock() *clock.Clock  { return s.clock }

func (f v8Scope) iso() *v8go.Isolate { return f.host.iso }
func (f v8Scope) Undefined() jsValue { return f.toJSValue(v8go.Undefined(f.iso())) }
func (f v8Scope) Null() jsValue      { return f.toJSValue(v8go.Null(f.iso())) }

func (f v8Scope) NewString(val string) jsValue { return f.newV8Value(val) }

func (f v8Scope) NewObject() jsObject {
	val, err := f.V8ScriptContext.v8ctx.RunScript("({})", "gost-dom/object")
	if err != nil {
		panic(fmt.Sprintf("cannot create object: %v", err))
	}
	obj, err := val.AsObject()
	if err != nil {
		panic(fmt.Sprintf("cannot evaluate: %v", err))
	}
	return newV8Object(f.V8ScriptContext, obj)
}

func (f v8Scope) NewUint8Array(data []byte) jsValue {
	bytes := make([]jsValue, len(data))
	for i, b := range data {
		bytes[i] = f.NewInt32(int32(b))
	}
	byteArray := f.NewArray(bytes...)
	from, err := f.v8ctx.RunScript("(data) => Uint8Array.from(data)", "gost-dom/v8host/uint8array")
	if err != nil {
		panic(fmt.Sprintf("gost-dom/v8host: Uint8Array.from: %v", err))
	}
	fn, err := from.AsFunction()
	if err != nil {
		panic(
			fmt.Sprintf(
				"gost-dom/v8host: Uint8Array.from: asFunction: %v (%v)",
				err, from,
			),
		)
	}
	res, err := fn.Call(f.v8ctx.Global(), byteArray.Self().Value)
	if err != nil {
		panic(fmt.Sprintf("gost-dom/v8host: Uint8Array.from: call: %v", err))
	}
	return f.toJSValue(res)
}

func (f v8Scope) NewInt32(val int32) jsValue   { return f.newV8Value(val) }
func (f v8Scope) NewUint32(val uint32) jsValue { return f.newV8Value(val) }
func (f v8Scope) NewInt64(val int64) jsValue   { return f.newV8Value(val) }
func (f v8Scope) NewBoolean(val bool) jsValue  { return f.newV8Value(val) }
func (s v8Scope) NewPromise() js.Promise[jsTypeParam] {
	return newV8Promise(s.V8ScriptContext)
}

func (s v8Scope) NewError(
	err error,
) js.Error[jsTypeParam] {
	return newV8Error(s.V8ScriptContext, err)
}

func (f v8Scope) JSONStringify(val jsValue) string {
	r, err := v8.JSONStringify(f.v8ctx, toV8Value(val))
	if err != nil {
		panic(fmt.Sprintf("JSONStringify: unexpected error: %v. %s", err, constants.BUG_ISSUE_URL))
	}
	return r
}

func (f v8Scope) JSONParse(val string) (jsValue, error) {
	v, err := v8.JSONParse(f.v8ctx, val)
	return f.toJSValue(v), err

}

func (f v8Scope) NewArray(values ...jsValue) jsValue {
	// Total hack, v8go doesn't expose Array values, so we polyfill the engine
	var err error
	arrayOf, err := f.v8ctx.RunScript("Array.of", "gost-polyfills-array")
	if err != nil {
		panic(err)
	}
	arrVal := f.toJSValue(arrayOf)
	if fn, ok := arrVal.AsFunction(); ok {
		res, err := fn.Call(f.global, values...)
		if err != nil {
			panic(err)
		}
		return res
	} else {
		panic("Array.of is not a function")
	}
}

func (f v8Scope) NewIterator(
	i iter.Seq2[js.Value[jsTypeParam], error],
) js.Value[jsTypeParam] {
	return f.host.iterator.newIterator(newV8Scope(f.V8ScriptContext), i)
}

func (f v8Scope) NewTypeError(msg string) error {
	return v8go.NewTypeError(f.iso(), msg)
}

func (f v8Scope) toJSValue(val *v8go.Value) jsValue {
	return newV8Value(f.V8ScriptContext, val)
}

// Creates a value in V8 from any value. This variant is hidden, as not all
// types are valid, and for type safety reasons, only valid types are exposed.
func (f v8Scope) newV8Value(val any) jsValue {
	// I'm unsure _when_ this could fail. AFAIK, v8 could throw an error if there is
	// currently an uncaught exception; but that scenario shouldn't occur from v8go.
	//
	// Maybe integer overflows?
	res, err := v8go.NewValue(f.iso(), val)
	if err != nil {
		panic(
			fmt.Sprintf(
				"v8 value conversion. value construction was assumed to succeed but failed: %v. %s",
				err,
				constants.BUG_ISSUE_URL,
			),
		)
	}
	return f.toJSValue(res)
}

func wrapV8Callback(
	host *V8ScriptHost,
	callback js.FunctionCallback[jsTypeParam],
) *v8go.FunctionTemplate {
	if callback == nil {
		return nil
	}
	return v8go.NewFunctionTemplateWithError(
		host.iso,
		func(info *v8go.FunctionCallbackInfo) (res *v8go.Value, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("PANIC in callback: %v\n%s", r, debug.Stack())
				}
			}()
			cbCtx := newCallbackContext(host, info)
			result, err := callback(cbCtx)
			return toV8Value(result), err
		},
	)
}

/* -------- v8Constructable -------- */

type v8Constructable struct {
	ctx  *V8ScriptContext
	ctor v8Class
}

func (c v8Constructable) NewInstance(nativeValue any) (jsObject, error) {
	val, err := c.ctor.ft.InstanceTemplate().NewInstance(c.ctx.v8ctx)
	obj := newV8Object(c.ctx, val).(*v8Object)
	if err == nil {
		obj.SetNativeValue(nativeValue)
		c.ctx.addDisposer(obj)
	}
	return obj, err
}
