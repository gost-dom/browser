// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("HTMLFormElement", "HTMLElement", createHTMLFormElementPrototype)
}

type htmlFormElementV8Wrapper struct {
	handleReffedObject[html.HTMLFormElement]
}

func newHTMLFormElementV8Wrapper(scriptHost *V8ScriptHost) *htmlFormElementV8Wrapper {
	return &htmlFormElementV8Wrapper{newHandleReffedObject[html.HTMLFormElement](scriptHost)}
}

func createHTMLFormElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLFormElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlFormElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("submit", wrapV8Callback(w.scriptHost, w.submit))
	prototypeTmpl.Set("requestSubmit", wrapV8Callback(w.scriptHost, w.requestSubmit))
	prototypeTmpl.Set("reset", wrapV8Callback(w.scriptHost, w.reset))
	prototypeTmpl.Set("checkValidity", wrapV8Callback(w.scriptHost, w.checkValidity))
	prototypeTmpl.Set("reportValidity", wrapV8Callback(w.scriptHost, w.reportValidity))

	prototypeTmpl.SetAccessorProperty("acceptCharset",
		wrapV8Callback(w.scriptHost, w.acceptCharset),
		wrapV8Callback(w.scriptHost, w.setAcceptCharset),
		v8.None)
	prototypeTmpl.SetAccessorProperty("action",
		wrapV8Callback(w.scriptHost, w.action),
		wrapV8Callback(w.scriptHost, w.setAction),
		v8.None)
	prototypeTmpl.SetAccessorProperty("autocomplete",
		wrapV8Callback(w.scriptHost, w.autocomplete),
		wrapV8Callback(w.scriptHost, w.setAutocomplete),
		v8.None)
	prototypeTmpl.SetAccessorProperty("enctype",
		wrapV8Callback(w.scriptHost, w.enctype),
		wrapV8Callback(w.scriptHost, w.setEnctype),
		v8.None)
	prototypeTmpl.SetAccessorProperty("encoding",
		wrapV8Callback(w.scriptHost, w.encoding),
		wrapV8Callback(w.scriptHost, w.setEncoding),
		v8.None)
	prototypeTmpl.SetAccessorProperty("method",
		wrapV8Callback(w.scriptHost, w.method),
		wrapV8Callback(w.scriptHost, w.setMethod),
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		wrapV8Callback(w.scriptHost, w.target),
		wrapV8Callback(w.scriptHost, w.setTarget),
		v8.None)
	prototypeTmpl.SetAccessorProperty("rel",
		wrapV8Callback(w.scriptHost, w.rel),
		wrapV8Callback(w.scriptHost, w.setRel),
		v8.None)
	prototypeTmpl.SetAccessorProperty("relList",
		wrapV8Callback(w.scriptHost, w.relList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("elements",
		wrapV8Callback(w.scriptHost, w.elements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
}

func (w htmlFormElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w htmlFormElementV8Wrapper) submit(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.submit")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	callErr := instance.Submit()
	return nil, callErr
}

func (w htmlFormElementV8Wrapper) requestSubmit(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.requestSubmit")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.HTMLFormElement](cbCtx.Instance())
	submitter, err1 := consumeArgument(cbCtx, "submitter", w.defaultHTMLElement, w.decodeHTMLElement)
	if cbCtx.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.RequestSubmit(submitter)
		return nil, callErr
	}
	return nil, errors.New("HTMLFormElement.requestSubmit: Missing arguments")
}

func (w htmlFormElementV8Wrapper) reset(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.reset")
	return nil, errors.New("HTMLFormElement.reset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) checkValidity(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.checkValidity")
	return nil, errors.New("HTMLFormElement.checkValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) reportValidity(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.reportValidity")
	return nil, errors.New("HTMLFormElement.reportValidity: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) acceptCharset(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.acceptCharset")
	return nil, errors.New("HTMLFormElement.acceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) setAcceptCharset(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setAcceptCharset")
	return nil, errors.New("HTMLFormElement.setAcceptCharset: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) action(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.action")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Action()
	return w.toString_(cbCtx.Context(), result)
}

func (w htmlFormElementV8Wrapper) setAction(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setAction")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.Context(), info, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetAction(val)
	return nil, nil
}

func (w htmlFormElementV8Wrapper) autocomplete(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.autocomplete")
	return nil, errors.New("HTMLFormElement.autocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) setAutocomplete(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setAutocomplete")
	return nil, errors.New("HTMLFormElement.setAutocomplete: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) enctype(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.enctype")
	return nil, errors.New("HTMLFormElement.enctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) setEnctype(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setEnctype")
	return nil, errors.New("HTMLFormElement.setEnctype: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) encoding(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.encoding")
	return nil, errors.New("HTMLFormElement.encoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) setEncoding(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setEncoding")
	return nil, errors.New("HTMLFormElement.setEncoding: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) method(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.method")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Method()
	return w.toString_(cbCtx.Context(), result)
}

func (w htmlFormElementV8Wrapper) setMethod(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setMethod")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.HTMLFormElement](cbCtx.Instance())
	val, err1 := parseSetterArg(cbCtx.Context(), info, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetMethod(val)
	return nil, nil
}

func (w htmlFormElementV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.target")
	return nil, errors.New("HTMLFormElement.target: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) setTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setTarget")
	return nil, errors.New("HTMLFormElement.setTarget: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) rel(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.rel")
	return nil, errors.New("HTMLFormElement.rel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) setRel(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setRel")
	return nil, errors.New("HTMLFormElement.setRel: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) relList(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.relList")
	return nil, errors.New("HTMLFormElement.relList: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w htmlFormElementV8Wrapper) elements(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.elements")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.HTMLFormElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Elements()
	return w.toHTMLFormControlsCollection(cbCtx.Context(), result)
}

func (w htmlFormElementV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.length")
	return nil, errors.New("HTMLFormElement.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
