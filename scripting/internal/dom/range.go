package dom

import (
	dominterfaces "github.com/gost-dom/browser/internal/interfaces/dom-interfaces"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func RangeConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return nil, nil
}

func decodeRange[T any](s js.Scope[T], v js.Value[T]) (dominterfaces.Range, error) {
	if js.IsNullish(v) {
		return nil, nil
	}
	if obj, ok := v.AsObject(); ok {
		if r, ok := obj.NativeValue().(dominterfaces.Range); ok {
			return r, nil
		}
	}
	return nil, s.NewTypeError("Value is not a Range")
}

func encodeRange[T any](s js.Scope[T], r dominterfaces.Range) (js.Value[T], error) {
	return s.Constructor("Range").NewInstance(r)
}
