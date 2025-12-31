package html

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeDOMStringMap[T any](class js.Class[T]) {
	class.CreateNamedHandler(
		js.WithGetterCallback(DOMStringMap_NamedPropertyGet[T]),
		js.WithSetterCallback(DOMStringMap_NamedPropertySet[T]),
		js.WithDeleterCallback(DOMStringMap_NamedPropertyDelete[T]),
		js.WithEnumeratorCallback(DOMStringMap_NamedPropertyEnumerator[T]),
	)
}

func DOMStringMap_NamedPropertyGet[T any](
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

func DOMStringMap_NamedPropertySet[T any](
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

func DOMStringMap_NamedPropertyDelete[T any](
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

func DOMStringMap_NamedPropertyEnumerator[T any](info js.CallbackScope[T]) ([]js.Value[T], error) {
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
