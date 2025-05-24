// This file is generated. Do not edit.

package v8host

import (
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	abstraction "github.com/gost-dom/browser/scripting/v8host/internal/abstraction"
	v8 "github.com/gost-dom/v8go"
)

func init() {
	registerJSClass("HTMLElement", "Element", createHTMLElementPrototype)
}

type htmlElementV8Wrapper struct {
	handleReffedObject[html.HTMLElement]
	htmlOrSVGElement *htmlOrSVGElementV8Wrapper
}

func newHTMLElementV8Wrapper(scriptHost *V8ScriptHost) *htmlElementV8Wrapper {
	return &htmlElementV8Wrapper{
		newHandleReffedObject[html.HTMLElement](scriptHost),
		newHTMLOrSVGElementV8Wrapper(scriptHost),
	}
}

func createHTMLElementPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHTMLElementV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w htmlElementV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("click", v8.NewFunctionTemplateWithError(iso, w.click))
	w.htmlOrSVGElement.installPrototype(prototypeTmpl)
}

func (w htmlElementV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w htmlElementV8Wrapper) click(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: HTMLElement.click")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.HTMLElement](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	instance.Click()
	return nil, nil
}
