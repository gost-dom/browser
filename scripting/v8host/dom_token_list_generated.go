// This file is generated. Do not edit.

package v8host

import (
	"errors"
	dom "github.com/gost-dom/browser/dom"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("DOMTokenList", "", createDOMTokenListPrototype)
}

type domTokenListV8Wrapper struct {
	handleReffedObject[dom.DOMTokenList]
}

func newDOMTokenListV8Wrapper(scriptHost *V8ScriptHost) *domTokenListV8Wrapper {
	return &domTokenListV8Wrapper{newHandleReffedObject[dom.DOMTokenList](scriptHost)}
}

func createDOMTokenListPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newDOMTokenListV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w domTokenListV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("item", v8.NewFunctionTemplateWithError(iso, w.item))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, w.contains))
	prototypeTmpl.Set("add", v8.NewFunctionTemplateWithError(iso, w.add))
	prototypeTmpl.Set("remove", v8.NewFunctionTemplateWithError(iso, w.remove))
	prototypeTmpl.Set("toggle", v8.NewFunctionTemplateWithError(iso, w.toggle))
	prototypeTmpl.Set("replace", v8.NewFunctionTemplateWithError(iso, w.replace))
	prototypeTmpl.Set("supports", v8.NewFunctionTemplateWithError(iso, w.supports))

	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("value",
		v8.NewFunctionTemplateWithError(iso, w.value),
		v8.NewFunctionTemplateWithError(iso, w.setValue),
		v8.None)
	prototypeTmpl.Set("toString", v8.NewFunctionTemplateWithError(iso, w.value))
}

func (w domTokenListV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w domTokenListV8Wrapper) item(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.item")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	index, err1 := tryParseArg(args, 0, w.decodeUnsignedLong)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Item(index)
		return w.toNullableString_(cbCtx.Context(), result)
	}
	return nil, errors.New("DOMTokenList.item: Missing arguments")
}

func (w domTokenListV8Wrapper) contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.contains")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	token, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(token)
		return w.toBoolean(cbCtx.Context(), result)
	}
	return nil, errors.New("DOMTokenList.contains: Missing arguments")
}

func (w domTokenListV8Wrapper) add(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.add")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	tokens, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.Add(tokens)
		return nil, callErr
	}
	return nil, errors.New("DOMTokenList.add: Missing arguments")
}

func (w domTokenListV8Wrapper) remove(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.remove")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	tokens, err1 := tryParseArg(args, 0, w.decodeString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Remove(tokens)
		return nil, nil
	}
	return nil, errors.New("DOMTokenList.remove: Missing arguments")
}

func (w domTokenListV8Wrapper) replace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.replace")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	token, err1 := tryParseArg(args, 0, w.decodeString)
	newToken, err2 := tryParseArg(args, 1, w.decodeString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result := instance.Replace(token, newToken)
		return w.toBoolean(cbCtx.Context(), result)
	}
	return nil, errors.New("DOMTokenList.replace: Missing arguments")
}

func (w domTokenListV8Wrapper) supports(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.supports")
	return nil, errors.New("DOMTokenList.supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w domTokenListV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.length")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return w.toUnsignedLong(cbCtx.Context(), result)
}

func (w domTokenListV8Wrapper) value(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.value")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return w.toString_(cbCtx.Context(), result)
}

func (w domTokenListV8Wrapper) setValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: DOMTokenList.setValue")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[dom.DOMTokenList](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.Context(), info, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
