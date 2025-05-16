package v8host

import (
	"github.com/gost-dom/browser/html"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("DOMStringMap", "", createDOMStringMapPrototype)
}

func createDOMStringMapPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newDOMStringMap(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)
	instanceTmpl.SetNamedHandler(wrapper)

	// wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}

type domStringMapV8Wrapper struct {
	handleReffedObject[*html.DOMStringMap, jsTypeParam]
}

func (w domStringMapV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func newDOMStringMap(host *V8ScriptHost) domStringMapV8Wrapper {
	return domStringMapV8Wrapper{newHandleReffedObject[*html.DOMStringMap](host)}
}

func (w domStringMapV8Wrapper) NamedPropertyGet(
	property *v8.Value,
	info v8.PropertyCallbackInfo,
) (*v8.Value, error) {
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	if !property.IsString() { // Don't intercept symbol properties
		return nil, v8.NotIntercepted
	}
	if val, found := instance.Get(property.String()); found {
		return v8.NewValue(w.iso(), val)
	}
	return nil, v8.NotIntercepted
}

func (w domStringMapV8Wrapper) NamedPropertyEnumerator(
	info v8.PropertyCallbackInfo,
) ([]*v8.Value, error) {
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	keys := instance.Keys()
	retVal := make([]*v8.Value, len(keys))
	for i, key := range keys {
		retVal[i], err = v8.NewValue(w.iso(), key)
	}
	return retVal, nil
}
