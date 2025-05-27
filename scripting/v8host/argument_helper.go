package v8host

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/v8go"
	v8 "github.com/gost-dom/v8go"
)

type argumentHelper struct {
	v8Info            *v8.FunctionCallbackInfo
	host              *V8ScriptHost
	noOfReadArguments int
	currentIndex      int
}

type jsCallbackContext = *argumentHelper

func newArgumentHelper(host *V8ScriptHost, info *v8.FunctionCallbackInfo) *argumentHelper {
	return &argumentHelper{info, host, 0, 0}
}

func (h argumentHelper) Global() jsObject {
	return h.ScriptCtx().global
}

func (h argumentHelper) This() jsObject       { return newV8Object(h.iso(), h.v8Info.This()) }
func (h argumentHelper) iso() *v8.Isolate     { return h.ScriptCtx().host.iso }
func (h argumentHelper) v8ctx() *v8.Context   { return h.ScriptCtx().v8ctx }
func (h argumentHelper) logger() *slog.Logger { return h.ScriptCtx().host.Logger() }
func (h *argumentHelper) ScriptCtx() *V8ScriptContext {
	return h.host.mustGetContext(h.v8Info.Context())
}

func (h *argumentHelper) ReturnWithValue(val *v8go.Value) (jsValue, error) {
	return h.ReturnWithJSValue(newV8Value(h.iso(), val))
}

func (h *argumentHelper) ReturnWithJSValue(val jsValue) (jsValue, error) {
	return val, nil
}

func (h *argumentHelper) ReturnWithValueErr(val *v8go.Value, err error) (jsValue, error) {
	return newV8Value(h.iso(), val), err
}

func (h *argumentHelper) ReturnWithJSValueErr(val jsValue, err error) (jsValue, error) {
	return val, err
}

func (h *argumentHelper) getInstanceForNode(node dom.Node) (jsValue, error) {
	v, err := h.ScriptCtx().getJSInstance(node)
	return h.ReturnWithJSValueErr(v, err)
}

func (h *argumentHelper) ReturnWithError(err error) (jsValue, error) {
	return nil, err
}

func (h *argumentHelper) ReturnWithTypeError(msg string) (jsValue, error) {
	return nil, v8.NewTypeError(h.iso(), msg)
}

func (h *argumentHelper) Instance() (any, error) {
	if h.v8Info.This().InternalFieldCount() < 1 {
		// TODO: Create a type error
		return nil, errors.New("TypeError")
	}
	return h.v8Info.This().GetInternalField(0).ExternalHandle().Value(), nil
}

// acceptIndex informs argumentHelper that argument at index was accepted. This
// affects when we query about how many arguments were read, in order to
// determine which overload to call in the system.
func (args *argumentHelper) acceptIndex(index int) {
	if args.noOfReadArguments == index {
		args.noOfReadArguments = index + 1
	}
}

// consumeValue works like [argumentHelper.consumeArg], but returns undefined
// instead of nil if the value doesn't exist.
func (h *argumentHelper) consumeValue() jsValue {
	if arg := h.ConsumeArg(); arg != nil {
		return arg
	}
	return &v8Value{h.iso(), v8.Undefined(h.iso())}
}

func (h *argumentHelper) consumeFunction() (jsFunction, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return nil, ErrWrongNoOfArguments
	}
	if f, ok := arg.AsFunction(); ok {
		return f, nil
	}
	return nil, h.newTypeError("Expected function")
}

func (h *argumentHelper) consumeInt32() (int32, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return 0, ErrWrongNoOfArguments
	}
	return arg.Int32(), nil
}

func (h *argumentHelper) consumeString() (string, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return "", ErrWrongNoOfArguments
	}
	return arg.String(), nil
}

func (h *argumentHelper) assertIndex(index int) {
	if index != h.currentIndex {
		panic(fmt.Sprintf("Bad index: %v (expected %v)", index, h.currentIndex))
	}
	h.currentIndex++
}

func (h *argumentHelper) ConsumeArg() jsValue {
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

func (h *argumentHelper) consumeRest() []*v8.Value {
	index := h.currentIndex
	// h.assertIndex(index)
	args := h.v8Info.Args()
	if len(args) <= index {
		return nil
	}
	h.currentIndex = len(args)
	return args[index:]
}

func (h *argumentHelper) newTypeError(msg string) error {
	return v8.NewTypeError(h.iso(), fmt.Sprintf(msg))
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
