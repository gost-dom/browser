// This file is generated. Do not edit.

package html

import (
	"errors"
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Window[T any] struct{}

func NewWindow[T any](scriptHost js.ScriptEngine[T]) *Window[T] {
	return &Window[T]{}
}

func (wrapper Window[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Window[T]) installPrototype(jsClass js.Class[T]) {
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

func (w Window[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Window[T]) close(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.close")
	return nil, errors.New("Window.close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) stop(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.stop")
	return nil, errors.New("Window.stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) focus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.focus")
	return nil, errors.New("Window.focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) blur(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.blur")
	return nil, errors.New("Window.blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) open(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.open")
	return nil, errors.New("Window.open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) alert(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.alert")
	return nil, errors.New("Window.alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) confirm(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.confirm")
	return nil, errors.New("Window.confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) print(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.print")
	return nil, errors.New("Window.print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) postMessage(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.postMessage")
	return nil, errors.New("Window.postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) self(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.self")
	return nil, errors.New("Window.self: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) document(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.document")
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Window[T]) name(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.name")
	return nil, errors.New("Window.name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) setName(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setName")
	return nil, errors.New("Window.setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) navigation(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.navigation")
	return nil, errors.New("Window.navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) customElements(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.customElements")
	return nil, errors.New("Window.customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) locationbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.locationbar")
	return nil, errors.New("Window.locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) menubar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.menubar")
	return nil, errors.New("Window.menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) personalbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.personalbar")
	return nil, errors.New("Window.personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) scrollbars(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.scrollbars")
	return nil, errors.New("Window.scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) statusbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.statusbar")
	return nil, errors.New("Window.statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) toolbar(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.toolbar")
	return nil, errors.New("Window.toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) status(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.status")
	return nil, errors.New("Window.status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) setStatus(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setStatus")
	return nil, errors.New("Window.setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) closed(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.closed")
	return nil, errors.New("Window.closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) frames(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.frames")
	return nil, errors.New("Window.frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) length(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.length")
	return nil, errors.New("Window.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) top(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.top")
	return nil, errors.New("Window.top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) opener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.opener")
	return nil, errors.New("Window.opener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) setOpener(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.setOpener")
	return nil, errors.New("Window.setOpener: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) frameElement(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.frameElement")
	return nil, errors.New("Window.frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) navigator(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.navigator")
	return nil, errors.New("Window.navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) clientInformation(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.clientInformation")
	return nil, errors.New("Window.clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) originAgentCluster(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("V8 Function call: Window.originAgentCluster")
	return nil, errors.New("Window.originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
