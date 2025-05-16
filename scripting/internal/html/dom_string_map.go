package html

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

// func createDOMStringMapPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
// 	iso := scriptHost.iso
// 	wrapper := NewDOMStringMap(scriptHost)
// 	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)
//
// 	instanceTmpl := constructor.InstanceTemplate()
// 	instanceTmpl.SetInternalFieldCount(1)
// 	instanceTmpl.SetNamedHandler(wrapper)
//
// 	// wrapper.installPrototype(constructor.PrototypeTemplate())
//
// 	return constructor
// }

type DOMStringMap[T any] struct{}

func NewDOMStringMap[T any](host js.ScriptEngine[T]) DOMStringMap[T] {
	return DOMStringMap[T]{}
}

func (w DOMStringMap[T]) Constructor(info js.CallbackContext[T]) (js.Value[T], error) {
	return info.ReturnWithTypeError("Illegal Constructor")
}

func (w DOMStringMap[T]) NamedPropertyGet(
	info js.GetterCallbackContext[T, js.Value[T]],
) (js.Value[T], error) {
	instance, err := js.As[*html.DOMStringMap](info.Instance())
	if err != nil {
		return nil, err
	}
	if !info.Key().IsString() { // Don't intercept symbol properties
		return nil, js.NotIntercepted
	}
	if val, found := instance.Get(info.Key().String()); found {
		return info.ValueFactory().NewString(val), nil
	}
	return nil, nil
}

func (w DOMStringMap[T]) NamedPropertyEnumerator(
	info js.GetterCallbackContext[T, js.Value[T]],
) ([]js.Value[T], error) {
	instance, err := js.As[*html.DOMStringMap](info.Instance())
	if err != nil {
		return nil, err
	}
	keys := instance.Keys()
	retVal := make([]js.Value[T], len(keys))
	fact := info.ValueFactory()
	for i, key := range keys {
		retVal[i] = fact.NewString(key)
	}
	return retVal, nil
}
