package codec

import (
	"github.com/gost-dom/browser/scripting/internal/js"
)

type OptionDecoder[T, U any] = func(js.Scope[T], js.Value[T]) (U, error)

type Options[T, U any] map[string]OptionDecoder[T, U]

func OptDecoder[T, U, V any](
	decode func(scope js.Scope[T], val js.Value[T]) (U, error),
	f func(U) V,
) OptionDecoder[T, V] {
	return func(scope js.Scope[T], val js.Value[T]) (res V, err error) {
		var tmp U
		tmp, err = decode(scope, val)
		if err == nil {
			res = f(tmp)
		}
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
			o, err := v(scope, opt)
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
