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
	registerJSClass("HTMLAnchorElement", "HTMLElement", createHTMLAnchorElementPrototype)
}

type htmlAnchorElementV8Wrapper struct {
	handleReffedObject[html.HTMLAnchorElement]
	htmlHyperlinkElementUtils *htmlHyperlinkElementUtilsV8Wrapper
}

func newHTMLAnchorElementV8Wrapper(scriptHost *V8ScriptHost) *htmlAnchorElementV8Wrapper {
	return &htmlAnchorElementV8Wrapper{
		newHandleReffedObject[html.HTMLAnchorElement](scriptHost),
		newHTMLHyperlinkElementUtilsV8Wrapper(scriptHost),
	}
}

func createHTMLAnchorElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLAnchorElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlAnchorElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso

	prototypeTmpl.SetAccessorProperty("target",
		v8.NewFunctionTemplateWithError(iso, w.target),
		v8.NewFunctionTemplateWithError(iso, w.setTarget),
		v8.None)
	w.htmlHyperlinkElementUtils.installPrototype(prototypeTmpl)
}

func (w htmlAnchorElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLAnchorElement.Constructor")
	args := newArgumentHelper(w.scriptHost, info)
	return args.ReturnWithTypeError("Illegal constructor")
}

func (w htmlAnchorElementV8Wrapper) target(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLAnchorElement.target")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.HTMLAnchorElement](args.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Target()
	return w.toString_(args.Context(), result)
}

func (w htmlAnchorElementV8Wrapper) setTarget(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLAnchorElement.setTarget")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := abstraction.As[html.HTMLAnchorElement](args.Instance())
	val, err1 := parseSetterArg(args.Context(), info, w.decodeString)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	instance.SetTarget(val)
	return nil, nil
}
