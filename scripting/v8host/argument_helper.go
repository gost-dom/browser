package v8host

import (
	"errors"
	"fmt"

	v8 "github.com/gost-dom/v8go"
)

type argumentHelper struct {
	*v8.FunctionCallbackInfo
	ctx               *V8ScriptContext
	noOfReadArguments int
	currentIndex      int
}

func newArgumentHelper(host *V8ScriptHost, info *v8.FunctionCallbackInfo) *argumentHelper {
	ctx := host.mustGetContext(info.Context())
	return &argumentHelper{info, ctx, 0, 0}
}

func (h argumentHelper) iso() *v8.Isolate           { return h.ctx.host.iso }
func (h *argumentHelper) Context() *V8ScriptContext { return h.ctx }

func (h *argumentHelper) ReturnWithTypeError(msg string) (*v8.Value, error) {
	return nil, v8.NewTypeError(h.iso(), msg)
}

func (h *argumentHelper) Instance() (any, error) {
	if h.This().InternalFieldCount() < 1 {
		// TODO: Create a type error
		return nil, errors.New("TypeError")
	}
	return h.This().GetInternalField(0).ExternalHandle().Value(), nil
}

// acceptIndex informs argumentHelper that argument at index was accepted. This
// affects when we query about how many arguments were read, in order to
// determine which overload to call in the system.
func (args *argumentHelper) acceptIndex(index int) {
	if args.noOfReadArguments <= index {
		args.noOfReadArguments = index + 1
	}
}

// consumeValue works like [argumentHelper.consumeArg], but returns undefined
// instead of nil if the value doesn't exist.
func (h *argumentHelper) consumeValue() *v8.Value {
	if arg := h.consumeArg(); arg != nil {
		return arg
	}
	return v8.Undefined(h.ctx.host.iso)
}

func (h *argumentHelper) consumeFunction() (*v8.Function, error) {
	arg := h.consumeArg()
	if arg == nil {
		return nil, ErrWrongNoOfArguments
	}
	if arg.IsFunction() {
		return arg.AsFunction()
	}
	return nil, h.newTypeError("Expected function", arg)
}

func (h *argumentHelper) consumeInt32() (int32, error) {
	arg := h.consumeArg()
	if arg == nil {
		return 0, ErrWrongNoOfArguments
	}
	if arg.IsNumber() {
		return arg.Int32(), nil
	}
	return 0, h.newTypeError("Expected int32", arg)
}

func (h *argumentHelper) consumeString() (string, error) {
	arg := h.consumeArg()
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

func (h *argumentHelper) consumeArg() *v8.Value {
	index := h.currentIndex
	h.assertIndex(index)
	args := h.FunctionCallbackInfo.Args()
	if len(args) <= index {
		return nil
	}
	arg := args[index]
	if arg.IsUndefined() {
		return nil
	}
	if h.noOfReadArguments <= index {
		h.noOfReadArguments = index + 1
	}
	return arg
}

func (h *argumentHelper) newTypeError(msg string, v *v8.Value) error {
	json, _ := v8.JSONStringify(h.ctx.v8ctx, v)
	return v8.NewTypeError(
		h.ctx.host.iso,
		fmt.Sprintf("TypeError: %s\n%s", msg, json),
	)
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
