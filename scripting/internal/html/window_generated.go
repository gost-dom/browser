// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type WindowV8Wrapper[T any] struct{}

func NewWindowV8Wrapper[T any](scriptHost js.ScriptEngine[T]) *WindowV8Wrapper[T] {
	return &WindowV8Wrapper[T]{}
}

func (wrapper WindowV8Wrapper[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w WindowV8Wrapper[T]) installPrototype(jsClass js.Class[T]) {
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

func (w WindowV8Wrapper[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w WindowV8Wrapper[T]) close(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.close")
	return cbCtx.ReturnWithError(errors.New("Window.close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) stop(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.stop")
	return cbCtx.ReturnWithError(errors.New("Window.stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) focus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.focus")
	return cbCtx.ReturnWithError(errors.New("Window.focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) blur(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.blur")
	return cbCtx.ReturnWithError(errors.New("Window.blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) open(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.open")
	return cbCtx.ReturnWithError(errors.New("Window.open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) alert(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.alert")
	return cbCtx.ReturnWithError(errors.New("Window.alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) confirm(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.confirm")
	return cbCtx.ReturnWithError(errors.New("Window.confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) print(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.print")
	return cbCtx.ReturnWithError(errors.New("Window.print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) postMessage(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.postMessage")
	return cbCtx.ReturnWithError(errors.New("Window.postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) self(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.self")
	return cbCtx.ReturnWithError(errors.New("Window.self: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) document(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.document")
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return cbCtx.ReturnWithError(err)
	}
	result := instance.Document()
	return codec.EncodeEntity(cbCtx, result)
}

func (w WindowV8Wrapper[T]) name(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.name")
	return cbCtx.ReturnWithError(errors.New("Window.name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) setName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setName")
	return cbCtx.ReturnWithError(errors.New("Window.setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) navigation(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.navigation")
	return cbCtx.ReturnWithError(errors.New("Window.navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) customElements(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.customElements")
	return cbCtx.ReturnWithError(errors.New("Window.customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) locationbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.locationbar")
	return cbCtx.ReturnWithError(errors.New("Window.locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) menubar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.menubar")
	return cbCtx.ReturnWithError(errors.New("Window.menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) personalbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.personalbar")
	return cbCtx.ReturnWithError(errors.New("Window.personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) scrollbars(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.scrollbars")
	return cbCtx.ReturnWithError(errors.New("Window.scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) statusbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.statusbar")
	return cbCtx.ReturnWithError(errors.New("Window.statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) toolbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.toolbar")
	return cbCtx.ReturnWithError(errors.New("Window.toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) status(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.status")
	return cbCtx.ReturnWithError(errors.New("Window.status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) setStatus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setStatus")
	return cbCtx.ReturnWithError(errors.New("Window.setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) closed(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.closed")
	return cbCtx.ReturnWithError(errors.New("Window.closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) frames(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.frames")
	return cbCtx.ReturnWithError(errors.New("Window.frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.length")
	return cbCtx.ReturnWithError(errors.New("Window.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) top(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.top")
	return cbCtx.ReturnWithError(errors.New("Window.top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) opener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.opener")
	return cbCtx.ReturnWithError(errors.New("Window.opener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) setOpener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setOpener")
	return cbCtx.ReturnWithError(errors.New("Window.setOpener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) frameElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.frameElement")
	return cbCtx.ReturnWithError(errors.New("Window.frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) navigator(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.navigator")
	return cbCtx.ReturnWithError(errors.New("Window.navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) clientInformation(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.clientInformation")
	return cbCtx.ReturnWithError(errors.New("Window.clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}

func (w WindowV8Wrapper[T]) originAgentCluster(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.originAgentCluster")
	return cbCtx.ReturnWithError(errors.New("Window.originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues"))
}
