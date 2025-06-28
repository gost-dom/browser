package codec

import (
	"fmt"

	"github.com/gost-dom/browser/scripting/internal/js"
)

type OptionDecoder[T, U any] = func(js.Value[T]) (U, error)

type Options[T, U any] map[string]OptionDecoder[T, U]

func OptDecoder[T, U, V any](f func(U) V) OptionDecoder[T, V] {
	return func(val js.Value[T]) (res V, err error) {
		obj, ok := val.AsObject()
		if !ok {
			err = fmt.Errorf("Not an object")
			return
		}
		if opt, ok := obj.NativeValue().(U); ok {
			return f(opt), nil
		}
		err = fmt.Errorf("Option not compatible")
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
		if err != nil {
			return nil, err
		}
		o, err := v(opt)
		if err != nil {
			return nil, err
		}
		opts = append(opts, o)
	}
	return
}
