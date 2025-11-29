// This file is generated. Do not edit.

package html

import (
	html "github.com/gost-dom/browser/html"
	log "github.com/gost-dom/browser/internal/log"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
	"log/slog"
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
	jsClass.CreatePrototypeAttribute("parent", w.parent, nil)
	jsClass.CreatePrototypeAttribute("frameElement", w.frameElement, nil)
	jsClass.CreatePrototypeAttribute("navigator", w.navigator, nil)
	jsClass.CreatePrototypeAttribute("clientInformation", w.clientInformation, nil)
	jsClass.CreatePrototypeAttribute("originAgentCluster", w.originAgentCluster, nil)
}

func (w Window[T]) Constructor(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "Constructor"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}

func (w Window[T]) close(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "close"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.close: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) stop(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "stop"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.stop: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) focus(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "focus"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.focus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) blur(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "blur"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.blur: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) open(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "open"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.open: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) alert(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "alert"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.alert: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) confirm(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "confirm"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.confirm: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) print(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "print"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.print: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) postMessage(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "postMessage"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.postMessage: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) document(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "document"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	instance, err := js.As[html.Window](cbCtx.Instance())
	if err != nil {
		return nil, err
	}
	result := instance.Document()
	return codec.EncodeEntity(cbCtx, result)
}

func (w Window[T]) name(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "name"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.name: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) setName(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "setName"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.setName: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) navigation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "navigation"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.navigation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) customElements(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "customElements"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.customElements: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) locationbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "locationbar"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.locationbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) menubar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "menubar"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.menubar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) personalbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "personalbar"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.personalbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) scrollbars(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "scrollbars"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.scrollbars: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) statusbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "statusbar"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.statusbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) toolbar(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "toolbar"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.toolbar: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) status(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "status"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.status: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) setStatus(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "setStatus"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.setStatus: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) closed(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "closed"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.closed: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) frames(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "frames"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.frames: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) length(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "length"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.length: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) top(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "top"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.top: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) frameElement(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "frameElement"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.frameElement: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) navigator(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "navigator"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.navigator: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) clientInformation(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "clientInformation"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.clientInformation: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w Window[T]) originAgentCluster(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	l := cbCtx.Logger().With(slog.String("IdlInterface", "Window"), slog.String("Method", "originAgentCluster"))
	l.Debug("JS function callback enter", js.ThisLogAttr(cbCtx), js.ArgsLogAttr(cbCtx))
	defer func() {
		l.Debug("JS function callback exit", js.LogAttr("res", res), log.ErrAttr(err))
	}()
	return codec.EncodeCallbackErrorf(cbCtx, "Window.originAgentCluster: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
