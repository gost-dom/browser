// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
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
	iso := w.scriptHost.iso
	prototypeTmpl.Set("submit", v8.NewFunctionTemplateWithError(iso, w.submit))
	prototypeTmpl.Set("requestSubmit", v8.NewFunctionTemplateWithError(iso, w.requestSubmit))
	prototypeTmpl.Set("reset", v8.NewFunctionTemplateWithError(iso, w.reset))
	prototypeTmpl.Set("checkValidity", v8.NewFunctionTemplateWithError(iso, w.checkValidity))
	prototypeTmpl.Set("reportValidity", v8.NewFunctionTemplateWithError(iso, w.reportValidity))

	prototypeTmpl.SetAccessorProperty("acceptCharset",
		v8.NewFunctionTemplateWithError(iso, w.acceptCharset),
		v8.NewFunctionTemplateWithError(iso, w.setAcceptCharset),
		v8.None)
	prototypeTmpl.SetAccessorProperty("action",
		v8.NewFunctionTemplateWithError(iso, w.action),
		v8.NewFunctionTemplateWithError(iso, w.setAction),
		v8.None)
	prototypeTmpl.SetAccessorProperty("autocomplete",
		v8.NewFunctionTemplateWithError(iso, w.autocomplete),
		v8.NewFunctionTemplateWithError(iso, w.setAutocomplete),
		v8.None)
	prototypeTmpl.SetAccessorProperty("enctype",
		v8.NewFunctionTemplateWithError(iso, w.enctype),
		v8.NewFunctionTemplateWithError(iso, w.setEnctype),
		v8.None)
	prototypeTmpl.SetAccessorProperty("encoding",
		v8.NewFunctionTemplateWithError(iso, w.encoding),
		v8.NewFunctionTemplateWithError(iso, w.setEncoding),
		v8.None)
	prototypeTmpl.SetAccessorProperty("method",
		v8.NewFunctionTemplateWithError(iso, w.method),
		v8.NewFunctionTemplateWithError(iso, w.setMethod),
		v8.None)
	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, w.target),
		v8.NewFunctionTemplateWithError(iso, w.setTarget),
		v8.None)
	prototypeTmpl.SetAccessorProperty("rel",
		v8.NewFunctionTemplateWithError(iso, w.rel),
		v8.NewFunctionTemplateWithError(iso, w.setRel),
		v8.None)
	prototypeTmpl.SetAccessorProperty("relList",
		v8.NewFunctionTemplateWithError(iso, w.relList),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("elements",
		v8.NewFunctionTemplateWithError(iso, w.elements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, w.length),
		nil,
		v8.None)
}

func (w htmlFormElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w htmlFormElementV8Wrapper) submit(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.submit")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Submit()
	return nil, callErr
}

func (w htmlFormElementV8Wrapper) requestSubmit(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.requestSubmit")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	submitter, err1 := tryParseArgWithDefault(args, 0, w.defaultHTMLElement, w.decodeHTMLElement)
	if args.noOfReadArguments >= 1 {
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
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Action()
	return w.toString_(ctx, result)
}

func (w htmlFormElementV8Wrapper) setAction(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setAction")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeString)
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
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Method()
	return w.toString_(ctx, result)
}

func (w htmlFormElementV8Wrapper) setMethod(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.setMethod")
	ctx := w.mustGetContext(info)
	instance, err0 := w.getInstance(info)
	val, err1 := parseSetterArg(ctx, info, w.decodeString)
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
	ctx := w.mustGetContext(info)
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Elements()
	return w.toHTMLFormControlsCollection(ctx, result)
}

func (w htmlFormElementV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLFormElement.length")
	return nil, errors.New("HTMLFormElement.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
