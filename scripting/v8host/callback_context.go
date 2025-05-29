package v8host

import (
	"errors"
	"fmt"
	"log/slog"
	"runtime/debug"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/constants"
	"github.com/gost-dom/v8go"
	v8 "github.com/gost-dom/v8go"
)

// type jsCallbackContext = js.CallbackContext[*v8Value]
type jsCallbackContext = *v8CallbackContext
type jsValueFactory = v8ValueFactory

var (
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)

type v8CallbackContext struct {
	v8Info            *v8.FunctionCallbackInfo
	host              *V8ScriptHost
	noOfReadArguments int
	currentIndex      int
}

func newCallbackContext(host *V8ScriptHost, info *v8.FunctionCallbackInfo) jsCallbackContext {
	return &v8CallbackContext{v8Info: info, host: host}
}

func (h v8CallbackContext) Global() jsObject {
	return h.ScriptCtx().global
}

func (h v8CallbackContext) This() jsObject       { return newV8Object(h.iso(), h.v8Info.This()) }
func (h v8CallbackContext) iso() *v8.Isolate     { return h.ScriptCtx().host.iso }
func (h v8CallbackContext) v8ctx() *v8.Context   { return h.ScriptCtx().v8ctx }
func (h v8CallbackContext) logger() *slog.Logger { return h.ScriptCtx().host.Logger() }

func (h *v8CallbackContext) ScriptCtx() *V8ScriptContext {
	return h.host.mustGetContext(h.v8Info.Context())
}

func (c *v8CallbackContext) ValueFactory() jsValueFactory { return v8ValueFactory{c.host, c} }

func (h *v8CallbackContext) ReturnWithValue(val jsValue) (jsValue, error) {
	return val, nil
}

func (h *v8CallbackContext) ReturnWithJSValue(val jsValue) (jsValue, error) {
	return val, nil
}

func (h *v8CallbackContext) ReturnWithValueErr(val jsValue, err error) (jsValue, error) {
	return val, err
}

func (h *v8CallbackContext) ReturnWithJSValueErr(val jsValue, err error) (jsValue, error) {
	return val, err
}

func (h *v8CallbackContext) getInstanceForNode(node dom.Node) (jsValue, error) {
	return h.ScriptCtx().getJSInstance(node)
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

// acceptIndex informs argumentHelper that argument at index was accepted. This
// affects when we query about how many arguments were read, in order to
// determine which overload to call in the system.
func (args *v8CallbackContext) acceptIndex(index int) {
	if args.noOfReadArguments == index {
		args.noOfReadArguments = index + 1
	}
}

// consumeValue works like [argumentHelper.consumeArg], but returns undefined
// instead of nil if the value doesn't exist.
func (h *v8CallbackContext) consumeValue() jsValue {
	if arg := h.ConsumeArg(); arg != nil {
		return arg
	}
	return &v8Value{h.iso(), v8.Undefined(h.iso())}
}

func (h *v8CallbackContext) consumeFunction() (jsFunction, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return nil, ErrWrongNoOfArguments
	}
	if f, ok := arg.AsFunction(); ok {
		return f, nil
	}
	return nil, h.newTypeError("Expected function")
}

func (h *v8CallbackContext) consumeInt32() (int32, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return 0, ErrWrongNoOfArguments
	}
	return arg.Int32(), nil
}

func (h *v8CallbackContext) consumeString() (string, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return "", ErrWrongNoOfArguments
	}
	return arg.String(), nil
}

func (h *v8CallbackContext) assertIndex(index int) {
	if index != h.currentIndex {
		panic(fmt.Sprintf("Bad index: %v (expected %v)", index, h.currentIndex))
	}
	h.currentIndex++
}

func (h *v8CallbackContext) ConsumeArg() jsValue {
	index := h.currentIndex
	h.assertIndex(index)
	args := h.v8Info.Args()
	if len(args) <= index {
		return nil
	}
	arg := args[index]
	if arg.IsUndefined() {
		return nil
	}
	h.acceptIndex(index)
	return &v8Value{h.iso(), arg}
}

func (h *v8CallbackContext) consumeRest() []*v8.Value {
	index := h.currentIndex
	// h.assertIndex(index)
	args := h.v8Info.Args()
	if len(args) <= index {
		return nil
	}
	h.currentIndex = len(args)
	return args[index:]
}

func (h *v8CallbackContext) newTypeError(msg string) error {
	return v8.NewTypeError(h.iso(), fmt.Sprintf(msg))
}

type v8ValueFactory struct {
	host *V8ScriptHost
	ctx  *v8CallbackContext
}

func (f v8ValueFactory) iso() *v8go.Isolate { return f.host.iso }
func (f v8ValueFactory) Null() jsValue      { return f.toVal(v8go.Null(f.iso())) }

func (f v8ValueFactory) NewString(val string) jsValue { return f.newV8Value(val) }
func (f v8ValueFactory) NewInt32(val int32) jsValue   { return f.newV8Value(val) }
func (f v8ValueFactory) NewUint32(val uint32) jsValue { return f.newV8Value(val) }
func (f v8ValueFactory) NewInt64(val int64) jsValue   { return f.newV8Value(val) }
func (f v8ValueFactory) NewBoolean(val bool) jsValue  { return f.newV8Value(val) }

func (f v8ValueFactory) NewArray(values ...jsValue) jsValue {
	// Total hack, v8go doesn't expose Array values, so we polyfill the engine
	var err error
	arrayOf, err := f.ctx.v8ctx().RunScript("Array.of", "gost-polyfills-array")
	if err != nil {
		panic(err)
	}
	arrVal := newV8Value(f.ctx.iso(), arrayOf)
	if fn, ok := arrVal.AsFunction(); ok {
		res, err := fn.Call(f.ctx.Global(), values...)
		if err != nil {
			panic(err)
		}
		return res
	} else {
		panic("Array.of is not a function")
	}
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
	return &v8Value{f.iso(), res}
}

func (f v8ValueFactory) toVal(val *v8go.Value) jsValue {
	return newV8Value(f.iso(), val)
}

type internalCallback func(jsCallbackContext) (jsValue, error)

func wrapV8Callback(
	host *V8ScriptHost,
	callback internalCallback,
) *v8go.FunctionTemplate {
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
			val := assertV8Value(result)
			return val.v8Value(), err
		},
	)
}

func wrapV8CallbackFn(
	host *V8ScriptHost,
	callback internalCallback,
) v8go.FunctionCallbackWithError {
	return func(info *v8go.FunctionCallbackInfo) (res *v8go.Value, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("PANIC in callback: %v\n%s", r, debug.Stack())
			}
		}()
		cbCtx := newCallbackContext(host, info)
		result, err := callback(cbCtx)
		val := assertV8Value(result)
		return val.v8Value(), err
	}
}

/* -------- Decoders -------- */

func decodeInt32(cbCtx jsCallbackContext, val jsValue) (int32, error) {
	return val.Int32(), nil
}
