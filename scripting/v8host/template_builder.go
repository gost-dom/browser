package v8host

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/scripting/internal/js"
)

type unconstructableV8Wrapper[T any] struct{}

func newUnconstructableV8Wrapper(
	host js.ScriptEngine[jsTypeParam],
) unconstructableV8Wrapper[jsTypeParam] {
	return unconstructableV8Wrapper[jsTypeParam]{}
}

func (w unconstructableV8Wrapper[T]) Constructor(cb js.CallbackContext[T]) (js.Value[T], error) {
	return nil, cb.ValueFactory().NewTypeError("Illegal constructor")
}
func (w unconstructableV8Wrapper[T]) Initialize(c jsClass) {}

func zeroValue[T any]() (res T) { return }

func consumeOptionalArg[T, U any](
	cbCtx js.CallbackContext[T],
	name string,
	decoders ...func(js.CallbackContext[T], js.Value[T]) (U, error),
) (result U, found bool, err error) {
	value, _ := cbCtx.ConsumeArg()
	if value == nil {
		return
	}
	found = true
	errs := make([]error, len(decoders))
	for i, parser := range decoders {
		result, errs[i] = parser(cbCtx, value)
		if errs[i] == nil {
			return
		}
	}
	// TODO: This should eventually become a TypeError in JS
	err = fmt.Errorf("tryParseArg: %s: %w", name, errors.Join(errs...))
	return
}
