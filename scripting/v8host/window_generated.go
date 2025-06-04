// This file is generated. Do not edit.

package v8host

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func init() {
	registerJSClass("Window", "EventTarget", createWindowPrototype)
}

type windowV8Wrapper struct {
	handleReffedObject[html.Window, jsTypeParam]
}

func newWindowV8Wrapper(scriptHost *V8ScriptHost) *windowV8Wrapper {
	return &windowV8Wrapper{newHandleReffedObject[html.Window](scriptHost)}
}

func createWindowPrototype(scriptHost *V8ScriptHost) v8Class {
	wrapper := newWindowV8Wrapper(scriptHost)
	constructor := wrapV8Callback(scriptHost, wrapper.constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	jsClass := newV8Class(scriptHost, constructor)
	wrapper.installPrototype(jsClass)

	return jsClass
}
func (wrapper windowV8Wrapper) initialize(jsClass v8Class) {
	wrapper.installPrototype(jsClass)
}

func (w windowV8Wrapper) installPrototype(jsClass v8Class) {
	jsClass.CreatePrototypeMethod("close", w.close)
	jsClass.CreatePrototypeMethod("stop", w.stop)
	jsClass.CreatePrototypeMethod("focus", w.focus)
	jsClass.CreatePrototypeMethod("blur", w.blur)
	jsClass.CreatePrototypeMethod("open", w.open)
	jsClass.CreatePrototypeMethod("alert", w.alert)
	jsClass.CreatePrototypeMethod("confirm", w.confirm)
	jsClass.CreatePrototypeMethod("print", w.print)
	jsClass.CreatePrototypeMethod("postMessage", w.postMessage)
	jsClass.CreatePrototypeAttribute("window", w.window, nil)
	jsClass.CreatePrototypeAttribute("self", w.self, nil)
	jsClass.CreatePrototypeAttribute("document", w.document, nil)
	jsClass.CreatePrototypeAttribute("name", w.name, w.setName)
	jsClass.CreatePrototypeAttribute("history", w.history, nil)
	jsClass.CreatePrototypeAttribute("navigation", w.navigation, nil)
	jsClass.CreatePrototypeAttribute("customElements", w.customElements, nil)
	jsClass.CreatePrototypeAttribute("locationbar", w.locationbar, nil)
	jsClass.CreatePrototypeAttribute("menubar", w.menubar, nil)
	jsClass.CreatePrototypeAttribute("personalbar", w.personalbar, nil)
	jsClass.CreatePrototypeAttribute("scrollbars", w.scrollbars, nil)
	jsClass.CreatePrototypeAttribute("statusbar", w.statusbar, nil)
	jsClass.CreatePrototypeAttribute("toolbar", w.toolbar, nil)
	jsClass.CreatePrototypeAttribute("status", w.status, w.setStatus)
	jsClass.CreatePrototypeAttribute("closed", w.closed, nil)
	jsClass.CreatePrototypeAttribute("frames", w.frames, nil)
	jsClass.CreatePrototypeAttribute("length", w.length, nil)
	jsClass.CreatePrototypeAttribute("top", w.top, nil)
	jsClass.CreatePrototypeAttribute("opener", w.opener, w.setOpener)
	jsClass.CreatePrototypeAttribute("frameElement", w.frameElement, nil)
	jsClass.CreatePrototypeAttribute("navigator", w.navigator, nil)
	jsClass.CreatePrototypeAttribute("clientInformation", w.clientInformation, nil)
	jsClass.CreatePrototypeAttribute("originAgentCluster", w.originAgentCluster, nil)
}

func (w windowV8Wrapper) constructor(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w windowV8Wrapper) close(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.close")
	return cbCtx.ReturnWithError(errors.New("Window.close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) stop(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.stop")
	return cbCtx.ReturnWithError(errors.New("Window.stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) focus(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.focus")
	return cbCtx.ReturnWithError(errors.New("Window.focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) blur(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.blur")
	return cbCtx.ReturnWithError(errors.New("Window.blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) open(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.open")
	return cbCtx.ReturnWithError(errors.New("Window.open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) alert(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.alert")
	return cbCtx.ReturnWithError(errors.New("Window.alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) confirm(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.confirm")
	return cbCtx.ReturnWithError(errors.New("Window.confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) print(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.print")
	return cbCtx.ReturnWithError(errors.New("Window.print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) postMessage(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.postMessage")
	return cbCtx.ReturnWithError(errors.New("Window.postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) self(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.self")
	return cbCtx.ReturnWithError(errors.New("Window.self: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) document(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.document")
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Document()
	return encodeEntity(cbCtx, result)
}

func (w windowV8Wrapper) name(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.name")
	return cbCtx.ReturnWithError(errors.New("Window.name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) setName(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setName")
	return cbCtx.ReturnWithError(errors.New("Window.setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) navigation(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.navigation")
	return cbCtx.ReturnWithError(errors.New("Window.navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) customElements(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.customElements")
	return cbCtx.ReturnWithError(errors.New("Window.customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) locationbar(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.locationbar")
	return cbCtx.ReturnWithError(errors.New("Window.locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) menubar(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.menubar")
	return cbCtx.ReturnWithError(errors.New("Window.menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) personalbar(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.personalbar")
	return cbCtx.ReturnWithError(errors.New("Window.personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) scrollbars(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.scrollbars")
	return cbCtx.ReturnWithError(errors.New("Window.scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) statusbar(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.statusbar")
	return cbCtx.ReturnWithError(errors.New("Window.statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) toolbar(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.toolbar")
	return cbCtx.ReturnWithError(errors.New("Window.toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) status(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.status")
	return cbCtx.ReturnWithError(errors.New("Window.status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) setStatus(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setStatus")
	return cbCtx.ReturnWithError(errors.New("Window.setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) closed(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.closed")
	return cbCtx.ReturnWithError(errors.New("Window.closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) frames(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.frames")
	return cbCtx.ReturnWithError(errors.New("Window.frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) length(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.length")
	return cbCtx.ReturnWithError(errors.New("Window.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) top(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.top")
	return cbCtx.ReturnWithError(errors.New("Window.top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) opener(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.opener")
	return cbCtx.ReturnWithError(errors.New("Window.opener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) setOpener(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setOpener")
	return cbCtx.ReturnWithError(errors.New("Window.setOpener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) frameElement(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.frameElement")
	return cbCtx.ReturnWithError(errors.New("Window.frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) navigator(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.navigator")
	return cbCtx.ReturnWithError(errors.New("Window.navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) clientInformation(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.clientInformation")
	return cbCtx.ReturnWithError(errors.New("Window.clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w windowV8Wrapper) originAgentCluster(cbCtx jsCallbackContext) (jsValue, error) {
	cbCtx.Logger().Debug("V8 Function call: Window.originAgentCluster")
	return cbCtx.ReturnWithError(errors.New("Window.originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
