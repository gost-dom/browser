package codec

import (
	"errors"
	"fmt"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func ZeroValue[T any]() (res T) { return }

func DecodeByteString[T any](s js.Scope[T], v js.Value[T]) (string, error) {
	if v == nil {
		v = s.Undefined()
	}
	res := v.String()
	b := []byte(res)
	for _, bb := range b {
		if bb < 0x20 || bb > 0x7E {
			return "", s.NewTypeError("cannot decode bytestring")
		}
	}
	return res, nil
}

func DecodeString[T any](s js.Scope[T], v js.Value[T]) (string, error) {
	if v == nil {
		v = s.Undefined()
	}
	return v.String(), nil
}

func DecodeBoolean[T any](_ js.Scope[T], val js.Value[T]) (bool, error) {
	return val.Boolean(), nil
}

func DecodeInt[T any](_ js.Scope[T], v js.Value[T]) (int, error) {
	return int(v.Int32()), nil
}

func DecodeNode[T any](s js.Scope[T], v js.Value[T]) (dom.Node, error) {
	if v.IsNull() {
		return nil, nil
	}
	if obj, ok := v.AsObject(); ok {
		if node, ok := obj.NativeValue().(dom.Node); ok {
			return node, nil
		}
	}
	return nil, s.NewTypeError("Value is not a node")
}

func DecodeAs[T, U any](s js.Scope[U], v js.Value[U]) (res T, err error) {
	if js.IsNullish(v) {
		return
	}
	if obj, ok := v.AsObject(); ok {
		if res, ok = obj.NativeValue().(T); ok {
			return
		}
	}
	err = s.NewTypeError(fmt.Sprintf("JavaScript value does not wrap an instance of %T", res))
	return
}

func DecodeHTMLElement[T any](s js.Scope[T], v js.Value[T]) (html.HTMLElement, error) {
	if obj, ok := v.AsObject(); ok {
		if res, ok := obj.NativeValue().(html.HTMLElement); ok {
			return res, nil
		}
	}
	return nil, s.NewTypeError("Value is not a node")
}

type EventInit struct {
	Bubbles    bool
	Cancelable bool
	Init       any
}

func DecodeEventInit[T any](
	_ js.Scope[T],
	val js.Value[T],
) (EventInit, error) {
	options, ok := val.AsObject()
	if !ok {
		return EventInit{}, errors.New("Not an event init object")
	}

	bubbles, err1 := options.Get("bubbles")
	cancelable, err2 := options.Get("cancelable")
	err := errors.Join(err1, err2)
	if err != nil {
		return EventInit{}, err
	}
	init := EventInit{
		Bubbles:    bubbles.Boolean(),
		Cancelable: cancelable.Boolean(),
	}
	return init, nil
}

func DecodeFunction[T any](s js.Scope[T], v js.Value[T]) (js.Function[T], error) {
	if f, ok := v.AsFunction(); ok {
		return f, nil
	}
	return nil, s.NewTypeError("Must be a function")
}

func DecodeInnerObject[T, U any](s js.Scope[T], v js.Value[T]) (res U, err error) {
	obj, ok := v.AsObject()
	if !ok {
		err = fmt.Errorf("gost-dom/codec: option not an object: %v", v)
		return
	}
	optVal := obj.NativeValue()
	if opt, ok := optVal.(U); ok {
		return opt, nil
	}
	err = fmt.Errorf("gost-dom/codec: option not of type %T: %v", res, optVal)
	return
}
