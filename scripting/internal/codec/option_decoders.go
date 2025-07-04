package codec

import (
	"fmt"

	"github.com/gost-dom/browser/scripting/internal/js"
)

type OptionDecoder[T, U any] = func(js.Value[T]) (U, error)

type Options[T, U any] map[string]OptionDecoder[T, U]

// OptDecoder extracts the "native value" of a JavaScript object
func OptDecoder[T, U, V any](f func(U) V) OptionDecoder[T, V] {
	return func(val js.Value[T]) (res V, err error) {
		obj, ok := val.AsObject()
		if !ok {
			err = fmt.Errorf("gost-dom/codec: option not an object: %v", val)
			return
		}
		optVal := obj.NativeValue()
		if opt, ok := optVal.(U); ok {
			return f(opt), nil
		}
		err = fmt.Errorf("gost-dom/codec: option not of type %T: %v", res, optVal)
		return
	}
}

func DecodeOptions[T, U any](
	scope js.Scope[T], val js.Value[T], specs Options[T, U],
) (opts []U, err error) {
	var obj js.Object[T]
	obj, ok := val.AsObject()
	if !ok {
		return nil, nil
	}
	for k, v := range specs {
		opt, err := obj.Get(k)
		if !opt.IsUndefined() {
			if err != nil {
				return nil, err
			}
			o, err := v(opt)
			if err != nil {
				return nil, err
			}
			opts = append(opts, o)
		}
	}
	return
}

func NewOptionsDecoder[T, U any](
	opts Options[T, U],
) func(js.CallbackContext[T], js.Value[T]) ([]U, error) {
	return func(scope js.CallbackContext[T], v js.Value[T]) ([]U, error) {
		return DecodeOptions(scope, v, opts)
	}
}
