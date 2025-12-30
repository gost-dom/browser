// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type Window[T any] struct {
	windowOrWorkerGlobalScope *WindowOrWorkerGlobalScope[T]
}

func NewWindow[T any](scriptHost js.ScriptEngine[T]) *Window[T] {
	return &Window[T]{NewWindowOrWorkerGlobalScope(scriptHost)}
}

func (wrapper Window[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w Window[T]) installPrototype(jsClass js.Class[T]) {
	jsClass.CreateOperation("close", w.close)
	jsClass.CreateOperation("stop", w.stop)
	jsClass.CreateOperation("focus", w.focus)
	jsClass.CreateOperation("blur", w.blur)
	jsClass.CreateOperation("open", w.open)
	jsClass.CreateOperation("alert", w.alert)
	jsClass.CreateOperation("confirm", w.confirm)
	jsClass.CreateOperation("print", w.print)
	jsClass.CreateOperation("postMessage", w.postMessage)
	jsClass.CreateAttribute("window", w.window, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("self", w.self, nil)
	jsClass.CreateAttribute("document", w.document, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("name", w.name, w.setName)
	jsClass.CreateAttribute("location", w.location, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("history", w.history, nil)
	jsClass.CreateAttribute("navigation", w.navigation, nil)
	jsClass.CreateAttribute("customElements", w.customElements, nil)
	jsClass.CreateAttribute("locationbar", w.locationbar, nil)
	jsClass.CreateAttribute("menubar", w.menubar, nil)
	jsClass.CreateAttribute("personalbar", w.personalbar, nil)
	jsClass.CreateAttribute("scrollbars", w.scrollbars, nil)
	jsClass.CreateAttribute("statusbar", w.statusbar, nil)
	jsClass.CreateAttribute("toolbar", w.toolbar, nil)
	jsClass.CreateAttribute("status", w.status, w.setStatus)
	jsClass.CreateAttribute("closed", w.closed, nil)
	jsClass.CreateAttribute("frames", w.frames, nil)
	jsClass.CreateAttribute("length", w.length, nil)
	jsClass.CreateAttribute("top", w.top, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("opener", w.opener, w.setOpener)
	jsClass.CreateAttribute("parent", w.parent, nil)
	jsClass.CreateAttribute("frameElement", w.frameElement, nil)
	jsClass.CreateAttribute("navigator", w.navigator, nil)
	jsClass.CreateAttribute("clientInformation", w.clientInformation, nil)
	jsClass.CreateAttribute("originAgentCluster", w.originAgentCluster, nil)
	w.windowOrWorkerGlobalScope.installPrototype(jsClass)
}

func WindowConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Window[T]) close(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) stop(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) focus(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) blur(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) open(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) alert(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) confirm(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) print(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) postMessage(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) document(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Window[T]) name(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) setName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) location(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Location()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Window[T]) navigation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) customElements(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) locationbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) menubar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) personalbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) scrollbars(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) statusbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) toolbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) status(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) setStatus(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) closed(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) frames(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) top(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) frameElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) navigator(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) clientInformation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) originAgentCluster(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
