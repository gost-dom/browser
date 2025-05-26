package v8host

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
	v8 "github.com/gost-dom/v8go"
)

type argumentHelper struct {
	*v8.FunctionCallbackInfo
	host              *V8ScriptHost
	noOfReadArguments int
	currentIndex      int
}

func newArgumentHelper(host *V8ScriptHost, info *v8.FunctionCallbackInfo) *argumentHelper {
	return &argumentHelper{info, host, 0, 0}
}

func (h argumentHelper) iso() *v8.Isolate     { return h.FunctionCallbackInfo.Context().Isolate() }
func (h argumentHelper) logger() *slog.Logger { return h.ScriptCtx().host.Logger() }
func (h *argumentHelper) ScriptCtx() *V8ScriptContext {
	return h.host.mustGetContext(h.FunctionCallbackInfo.Context())
}

func (h *argumentHelper) ReturnWithValue(val *v8.Value) js.CallbackRVal {
	return v8CallbackRVal{rtnVal: val}
}
func (h *argumentHelper) ReturnWithValueErr(val *v8.Value, err error) js.CallbackRVal {
	return v8CallbackRVal{val, err}
}

func (h *argumentHelper) getInstanceForNode(node dom.Node) js.CallbackRVal {
	val, err := h.ScriptCtx().getInstanceForNode(node)
	if err != nil {
		return h.ReturnWithError(err)
	} else {
		return h.ReturnWithValue(val)
	}
}

func (h *argumentHelper) ReturnWithError(
	err error,
) js.CallbackRVal {
	return v8CallbackRVal{err: err}
}

func (h *argumentHelper) ReturnWithTypeError(msg string) js.CallbackRVal {
	return v8CallbackRVal{err: v8.NewTypeError(h.iso(), msg)}
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
	if args.noOfReadArguments == index {
		args.noOfReadArguments = index + 1
	}
}

// consumeValue works like [argumentHelper.consumeArg], but returns undefined
// instead of nil if the value doesn't exist.
func (h *argumentHelper) consumeValue() *v8.Value {
	if arg := h.ConsumeArg(); arg != nil {
		return arg
	}
	return v8.Undefined(h.FunctionCallbackInfo.Context().Isolate())
}

func (h *argumentHelper) consumeFunction() (*v8.Function, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return nil, ErrWrongNoOfArguments
	}
	if arg.IsFunction() {
		return arg.AsFunction()
	}
	return nil, h.newTypeError("Expected function", arg)
}

func (h *argumentHelper) consumeInt32() (int32, error) {
	arg := h.ConsumeArg()
	if arg == nil {
		return 0, ErrWrongNoOfArguments
	}
	if arg.IsNumber() {
		return arg.Int32(), nil
	}
	return 0, h.newTypeError("Expected int32", arg)
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

func (h *argumentHelper) ConsumeArg() *v8.Value {
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
	h.acceptIndex(index)
	return arg
}

func (h *argumentHelper) consumeRest() []*v8.Value {
	index := h.currentIndex
	// h.assertIndex(index)
	args := h.FunctionCallbackInfo.Args()
	if len(args) <= index {
		return nil
	}
	h.currentIndex = len(args)
	return args[index:]
}

func (h *argumentHelper) newTypeError(msg string, v *v8.Value) error {
	json, _ := v8.JSONStringify(h.FunctionCallbackInfo.Context(), v)
	return v8.NewTypeError(
		h.iso(),
		fmt.Sprintf("TypeError: %s\n%s", msg, json),
	)
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
