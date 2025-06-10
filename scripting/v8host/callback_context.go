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

type jsScriptEngine = js.ScriptEngine[jsTypeParam]
type jsCallbackContext = js.CallbackContext[*v8Value]
type jsValueFactory = js.ValueFactory[*v8Value]

var (
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)

type v8CallbackContext struct {
	v8Info       *v8.FunctionCallbackInfo
	host         *V8ScriptHost
	currentIndex int
}

func newCallbackContext(host *V8ScriptHost, info *v8.FunctionCallbackInfo) jsCallbackContext {
	return &v8CallbackContext{v8Info: info, host: host}
}

func (h v8CallbackContext) This() jsObject {
	return newV8Object(h.ScriptCtx(), h.v8Info.This())
}
func (h v8CallbackContext) iso() *v8.Isolate   { return h.ScriptCtx().host.iso }
func (h v8CallbackContext) v8ctx() *v8.Context { return h.ScriptCtx().v8ctx }

func (h *v8CallbackContext) ScriptCtx() *V8ScriptContext {
	return h.host.mustGetContext(h.v8Info.Context())
}

func (c *v8CallbackContext) Scope() js.Scope[jsTypeParam] { return newV8Scope(c.ScriptCtx()) }

func (c *v8CallbackContext) ValueFactory() jsValueFactory {
	return v8ValueFactory{c.host, c.ScriptCtx()}
}

func (h *v8CallbackContext) ReturnWithValue(val jsValue) (jsValue, error) {
	return val, nil
}

func (h *v8CallbackContext) ReturnWithError(err error) (jsValue, error) {
	return nil, err
}

func (h *v8CallbackContext) ReturnWithTypeError(msg string) (jsValue, error) {
	return nil, v8.NewTypeError(h.iso(), msg)
}

func (h *v8CallbackContext) Instance() (any, error) {
	if h.v8Info.This().InternalFieldCount() < 1 {
		return h.ReturnWithTypeError("No internal instance")
	}
	return h.v8Info.This().GetInternalField(0).ExternalHandle().Value(), nil
}

func (c *v8CallbackContext) Logger() *slog.Logger {
	return c.host.Logger()
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
	return newV8Value(h.ScriptCtx(), arg), true
}

/* -------- v8ValueFactory -------- */

type v8ValueFactory struct {
	host *V8ScriptHost
	ctx  *V8ScriptContext
}

func (f v8ValueFactory) iso() *v8go.Isolate { return f.host.iso }
func (f v8ValueFactory) Null() jsValue      { return newV8Value(f.ctx, v8go.Null(f.iso())) }

func (f v8ValueFactory) NewString(val string) jsValue { return f.newV8Value(val) }
func (f v8ValueFactory) NewInt32(val int32) jsValue   { return f.newV8Value(val) }
func (f v8ValueFactory) NewUint32(val uint32) jsValue { return f.newV8Value(val) }
func (f v8ValueFactory) NewInt64(val int64) jsValue   { return f.newV8Value(val) }
func (f v8ValueFactory) NewBoolean(val bool) jsValue  { return f.newV8Value(val) }

func (f v8ValueFactory) JSONStringify(val jsValue) string {
	r, err := v8.JSONStringify(f.ctx.v8ctx, toV8Value(val))
	if err != nil {
		panic(fmt.Sprintf("JSONStringify: unexpected error: %v. %s", err, constants.BUG_ISSUE_URL))
	}
	return r
}

func (f v8ValueFactory) JSONParse(val string) (jsValue, error) {
	v, err := v8.JSONParse(f.ctx.v8ctx, val)
	return newV8Value(f.ctx, v), err

}

func (f v8ValueFactory) NewArray(values ...jsValue) jsValue {
	// Total hack, v8go doesn't expose Array values, so we polyfill the engine
	var err error
	arrayOf, err := f.ctx.v8ctx.RunScript("Array.of", "gost-polyfills-array")
	if err != nil {
		panic(err)
	}
	arrVal := newV8Value(f.ctx, arrayOf)
	if fn, ok := arrVal.AsFunction(); ok {
		res, err := fn.Call(f.ctx.global, values...)
		if err != nil {
			panic(err)
		}
		return res
	} else {
		panic("Array.of is not a function")
	}
}

func (f v8ValueFactory) NewIterator(
	i iter.Seq2[js.Value[jsTypeParam], error],
) js.Value[jsTypeParam] {
	return f.host.iterator.newIterator(v8Scope{f.host, f.ctx}, i)
}

func (f v8ValueFactory) NewTypeError(msg string) error {
	return v8go.NewTypeError(f.iso(), msg)
}

// Creates a value in V8 from any value. This variant is hidden, as not all
// types are valid, and for type safety reasons, only valid types are exposed.
func (f v8ValueFactory) newV8Value(val any) jsValue {
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
	return newV8Value(f.ctx, res)
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

/* -------- v8Scope -------- */

type v8Scope struct {
	host *V8ScriptHost
	*V8ScriptContext
}

func newV8Scope(ctx *V8ScriptContext) v8Scope { return v8Scope{ctx.host, ctx} }

func (s v8Scope) Window() html.Window  { return s.window }
func (s v8Scope) GlobalThis() jsObject { return s.global }
func (s v8Scope) Clock() *clock.Clock  { return s.clock }

func (s v8Scope) ValueFactory() js.ValueFactory[jsTypeParam] {
	return v8ValueFactory{host: s.host, ctx: s.V8ScriptContext}
}

//
// func (c v8Scope) Constructor(name string) js.Constructor[jsTypeParam] {
// 	return
// }

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

/* -------- v8IndexedHandlers -------- */

type indexedGetterCallback = func(js.GetterCallbackContext[jsTypeParam, int]) (jsValue, error)

type v8GetterCallbackContext struct {
	*v8CallbackContext
}

func newIndexedGetterCallbackContext(
	host *V8ScriptHost,
	info *v8.FunctionCallbackInfo,
) js.GetterCallbackContext[jsTypeParam, int] {
	return v8GetterCallbackContext{&v8CallbackContext{v8Info: info, host: host}}
}

func (c v8GetterCallbackContext) Key() int {
	return int(c.v8Info.Index())
}
