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
	jsClass.CreateOperation("close", Window_close)
	jsClass.CreateOperation("stop", Window_stop)
	jsClass.CreateOperation("focus", Window_focus)
	jsClass.CreateOperation("blur", Window_blur)
	jsClass.CreateOperation("open", Window_open)
	jsClass.CreateOperation("alert", Window_alert)
	jsClass.CreateOperation("confirm", Window_confirm)
	jsClass.CreateOperation("print", Window_print)
	jsClass.CreateOperation("postMessage", Window_postMessage)
	jsClass.CreateAttribute("window", Window_window, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("self", Window_self, nil)
	jsClass.CreateAttribute("document", Window_document, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("name", Window_name, Window_setName)
	jsClass.CreateAttribute("location", Window_location, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("history", Window_history, nil)
	jsClass.CreateAttribute("navigation", Window_navigation, nil)
	jsClass.CreateAttribute("customElements", Window_customElements, nil)
	jsClass.CreateAttribute("locationbar", Window_locationbar, nil)
	jsClass.CreateAttribute("menubar", Window_menubar, nil)
	jsClass.CreateAttribute("personalbar", Window_personalbar, nil)
	jsClass.CreateAttribute("scrollbars", Window_scrollbars, nil)
	jsClass.CreateAttribute("statusbar", Window_statusbar, nil)
	jsClass.CreateAttribute("toolbar", Window_toolbar, nil)
	jsClass.CreateAttribute("status", Window_status, Window_setStatus)
	jsClass.CreateAttribute("closed", Window_closed, nil)
	jsClass.CreateAttribute("frames", Window_frames, nil)
	jsClass.CreateAttribute("length", Window_length, nil)
	jsClass.CreateAttribute("top", Window_top, nil, js.LegacyUnforgeable())
	jsClass.CreateAttribute("opener", Window_opener, Window_setOpener)
	jsClass.CreateAttribute("parent", Window_parent, nil)
	jsClass.CreateAttribute("frameElement", Window_frameElement, nil)
	jsClass.CreateAttribute("navigator", Window_navigator, nil)
	jsClass.CreateAttribute("clientInformation", Window_clientInformation, nil)
	jsClass.CreateAttribute("originAgentCluster", Window_originAgentCluster, nil)
	w.windowOrWorkerGlobalScope.installPrototype(jsClass)
}

func WindowConstructor[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func Window_close[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_stop[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_focus[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_blur[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_open[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_alert[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_confirm[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_print[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_postMessage[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_document[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return codec.EncodeEntity(cbCtx, result)
}

func Window_name[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_setName[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_location[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Location()
	return codec.EncodeEntity(cbCtx, result)
}

func Window_navigation[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_customElements[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_locationbar[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_menubar[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_personalbar[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_scrollbars[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_statusbar[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_toolbar[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_status[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_setStatus[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_closed[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_frames[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_length[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_top[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_frameElement[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_navigator[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_clientInformation[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func Window_originAgentCluster[T any](cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	return codec.EncodeCallbackErrorf(cbCtx, "Window.Window_originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
