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
	registerJSClass("Window", "EventTarget", createWindowPrototype)
}

type windowV8Wrapper struct {
	handleReffedObject[html.Window]
}

func newWindowV8Wrapper(scriptHost *V8ScriptHost) *windowV8Wrapper {
	return &windowV8Wrapper{newHandleReffedObject[html.Window](scriptHost)}
}

func createWindowPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newWindowV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w windowV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	prototypeTmpl.Set("close", wrapV8Callback(w.scriptHost, w.close))
	prototypeTmpl.Set("stop", wrapV8Callback(w.scriptHost, w.stop))
	prototypeTmpl.Set("focus", wrapV8Callback(w.scriptHost, w.focus))
	prototypeTmpl.Set("blur", wrapV8Callback(w.scriptHost, w.blur))
	prototypeTmpl.Set("open", wrapV8Callback(w.scriptHost, w.open))
	prototypeTmpl.Set("alert", wrapV8Callback(w.scriptHost, w.alert))
	prototypeTmpl.Set("confirm", wrapV8Callback(w.scriptHost, w.confirm))
	prototypeTmpl.Set("print", wrapV8Callback(w.scriptHost, w.print))
	prototypeTmpl.Set("postMessage", wrapV8Callback(w.scriptHost, w.postMessage))

	prototypeTmpl.SetAccessorProperty("window",
		wrapV8Callback(w.scriptHost, w.window),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("self",
		wrapV8Callback(w.scriptHost, w.self),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("document",
		wrapV8Callback(w.scriptHost, w.document),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("name",
		wrapV8Callback(w.scriptHost, w.name),
		wrapV8Callback(w.scriptHost, w.setName),
		v8.None)
	prototypeTmpl.SetAccessorProperty("history",
		wrapV8Callback(w.scriptHost, w.history),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigation",
		wrapV8Callback(w.scriptHost, w.navigation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("customElements",
		wrapV8Callback(w.scriptHost, w.customElements),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("locationbar",
		wrapV8Callback(w.scriptHost, w.locationbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("menubar",
		wrapV8Callback(w.scriptHost, w.menubar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("personalbar",
		wrapV8Callback(w.scriptHost, w.personalbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("scrollbars",
		wrapV8Callback(w.scriptHost, w.scrollbars),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("statusbar",
		wrapV8Callback(w.scriptHost, w.statusbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("toolbar",
		wrapV8Callback(w.scriptHost, w.toolbar),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("status",
		wrapV8Callback(w.scriptHost, w.status),
		wrapV8Callback(w.scriptHost, w.setStatus),
		v8.None)
	prototypeTmpl.SetAccessorProperty("closed",
		wrapV8Callback(w.scriptHost, w.closed),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("frames",
		wrapV8Callback(w.scriptHost, w.frames),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("length",
		wrapV8Callback(w.scriptHost, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("top",
		wrapV8Callback(w.scriptHost, w.top),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("opener",
		wrapV8Callback(w.scriptHost, w.opener),
		wrapV8Callback(w.scriptHost, w.setOpener),
		v8.None)
	prototypeTmpl.SetAccessorProperty("frameElement",
		wrapV8Callback(w.scriptHost, w.frameElement),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("navigator",
		wrapV8Callback(w.scriptHost, w.navigator),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("clientInformation",
		wrapV8Callback(w.scriptHost, w.clientInformation),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("originAgentCluster",
		wrapV8Callback(w.scriptHost, w.originAgentCluster),
		nil,
		v8.None)
}

func (w windowV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.Constructor")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w windowV8Wrapper) close(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.close")
	return nil, errors.New("Window.close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) stop(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.stop")
	return nil, errors.New("Window.stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) focus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.focus")
	return nil, errors.New("Window.focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) blur(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.blur")
	return nil, errors.New("Window.blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) open(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.open")
	return nil, errors.New("Window.open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) alert(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.alert")
	return nil, errors.New("Window.alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) confirm(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.confirm")
	return nil, errors.New("Window.confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) print(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.print")
	return nil, errors.New("Window.print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) postMessage(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.postMessage")
	return nil, errors.New("Window.postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) self(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.self")
	return nil, errors.New("Window.self: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) document(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.document")
	cbCtx := newArgumentHelper(w.scriptHost, info)
	instance, err := abstraction.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return cbCtx.Context().getInstanceForNode(result)
}

func (w windowV8Wrapper) name(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.name")
	return nil, errors.New("Window.name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) setName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.setName")
	return nil, errors.New("Window.setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) navigation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.navigation")
	return nil, errors.New("Window.navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) customElements(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.customElements")
	return nil, errors.New("Window.customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) locationbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.locationbar")
	return nil, errors.New("Window.locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) menubar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.menubar")
	return nil, errors.New("Window.menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) personalbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.personalbar")
	return nil, errors.New("Window.personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) scrollbars(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.scrollbars")
	return nil, errors.New("Window.scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) statusbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.statusbar")
	return nil, errors.New("Window.statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) toolbar(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.toolbar")
	return nil, errors.New("Window.toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) status(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.status")
	return nil, errors.New("Window.status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) setStatus(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.setStatus")
	return nil, errors.New("Window.setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) closed(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.closed")
	return nil, errors.New("Window.closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) frames(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.frames")
	return nil, errors.New("Window.frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.length")
	return nil, errors.New("Window.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) top(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.top")
	return nil, errors.New("Window.top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) opener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.opener")
	return nil, errors.New("Window.opener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) setOpener(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.setOpener")
	return nil, errors.New("Window.setOpener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) frameElement(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.frameElement")
	return nil, errors.New("Window.frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) navigator(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.navigator")
	return nil, errors.New("Window.navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) clientInformation(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.clientInformation")
	return nil, errors.New("Window.clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w windowV8Wrapper) originAgentCluster(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug(w.logger(info), "V8 Function call: Window.originAgentCluster")
	return nil, errors.New("Window.originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
