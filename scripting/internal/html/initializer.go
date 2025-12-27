package html

import (
	"time"

	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type unconstructable[T any] struct{}

func NewUnconstructable[T any](host js.ScriptEngine[T]) unconstructable[T] {
	return unconstructable[T]{}
}

func (w unconstructable[T]) Initialize(c js.Class[T]) {}

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	Bootstrap(e)

	w := NewWindow(e)
	eventTarget, _ := e.Class("EventTarget")
	window := e.ConfigureGlobalScope("Window", eventTarget)
	w.Initialize(window)
	js.RegisterClass(e, "DOMStringMap", "", NewDOMStringMap, DOMStringMapConstructor)
	installEventLoopGlobals(window)

	// HTMLDocument exists as a separate class for historical reasons, but it
	// can be treated merely as an alias for Document. In Firefox, there is an
	// inheritance relationship between the two, which is modelled here.
	//
	// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLDocument
	js.RegisterClass(e, "HTMLDocument", "Document", NewHTMLDocument, HTMLDocumentConstructor)
	for _, cls := range codec.HtmlElements {
		if _, ok := e.Class(cls); !ok && cls != "HTMLElement" {
			js.RegisterClass(e, cls, "HTMLElement", NewUnconstructable, js.IllegalConstructor)
		}
	}
}

func installEventLoopGlobals[T any](global js.Class[T]) {
	global.CreateOperation("requestAnimationFrame", requestAnimationFrame)
}

func requestAnimationFrame[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	f, err := js.ConsumeArgument(cbCtx, "fn", nil, codec.DecodeFunction)
	if err != nil {
		return nil, err
	}
	cbCtx.Clock().AddSafeTask(
		func() {
			if _, err := f.Call(cbCtx.GlobalThis()); err != nil {
				dom.HandleJSCallbackError(cbCtx, "requestAnimationFrame", err)
			}
		}, 10*time.Millisecond)
	return nil, err
}
