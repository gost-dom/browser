package html

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

type DOMStringMap[T any] struct{}

func NewDOMStringMap[T any](host js.ScriptEngine[T]) DOMStringMap[T] {
	return DOMStringMap[T]{}
}

func (w DOMStringMap[T]) Constructor(info js.CallbackContext[T]) (js.Value[T], error) {
	return info.ReturnWithTypeError("Illegal Constructor")
}

func (w DOMStringMap[T]) Initialize(class js.Class[T]) {
	class.CreateNamedHandler(
		js.WithGetterCallback(w.NamedPropertyGet),
		js.WithSetterCallback(w.NamedPropertySet),
		js.WithDeleterCallback(w.NamedPropertyDelete),
		js.WithEnumeratorCallback(w.NamedPropertyEnumerator),
	)
}

func (w DOMStringMap[T]) NamedPropertyGet(
	info js.CallbackScope[T], key js.Value[T],
) (js.Value[T], error) {
	instance, err := js.As[*html.DOMStringMap](info.Instance())
	if err != nil {
		return nil, err
	}
	if !key.IsString() { // Don't intercept symbol properties
		return nil, js.NotIntercepted
	}
	if val, found := instance.Get(key.String()); found {
		return info.NewString(val), nil
	}
	return nil, js.NotIntercepted
}

func (w DOMStringMap[T]) NamedPropertySet(
	info js.CallbackScope[T], key, value js.Value[T],
) error {
	instance, err := js.As[*html.DOMStringMap](info.Instance())
	if err != nil {
		return err
	}
	if !key.IsString() { // Don't intercept symbol properties
		return js.NotIntercepted
	}
	instance.Set(key.String(), value.String())
	return nil
}

func (w DOMStringMap[T]) NamedPropertyDelete(
	info js.CallbackScope[T], key js.Value[T],
) (bool, error) {
	instance, err := js.As[*html.DOMStringMap](info.Instance())
	if err != nil {
		return false, err
	}
	if !key.IsString() { // Don't intercept symbol properties
		return false, js.NotIntercepted
	}
	instance.Delete(key.String())
	return true, nil
}

func (w DOMStringMap[T]) NamedPropertyEnumerator(info js.CallbackScope[T]) ([]js.Value[T], error) {
	instance, err := js.As[*html.DOMStringMap](info.Instance())
	if err != nil {
		return nil, err
	}
	keys := instance.Keys()
	retVal := make([]js.Value[T], len(keys))
	for i, key := range keys {
		retVal[i] = info.NewString(key)
	}
	return retVal, nil
}
