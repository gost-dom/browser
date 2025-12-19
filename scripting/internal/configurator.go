package internal

import (
	"time"

	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/scripting/internal/mathml"
	"github.com/gost-dom/browser/scripting/internal/streams"
	"github.com/gost-dom/browser/scripting/internal/svg"
	"github.com/gost-dom/browser/scripting/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/url"
	"github.com/gost-dom/browser/scripting/internal/xhr"
)

func Configure[T any](host js.ScriptEngine[T]) {
	dom.Configure(host)
	fetch.Configure(host)
	configureConsole(host)

	host.CreateFunction(
		"requestAnimationFrame",
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
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
		},
	)
	host.SetUnhandledPromiseRejectionHandler(
		js.ErrorHandlerFunc[T](handleUnhandledPromiseRejection[T]),
	)
}

func handleUnhandledPromiseRejection[T any](scope js.Scope[T], err error) {
	dom.HandleJSCallbackError(scope, "promiseRejected", err)
}

func Bootstrap[T any](e js.ScriptEngine[T]) {
	dom.Register(e)
	html.InitBuilder(e)
	svg.Bootstrap(e)
	mathml.Bootstrap(e)
	xhr.Bootstrap(e)
	url.Bootstrap(e)
	uievents.Bootstrap(e)
	fetch.Bootstrap(e)
	streams.Bootstrap(e)

	js.RegisterClass(e, "File", "", dom.NewEvent)
	js.RegisterClass(e, "CustomEvent", "Event", dom.NewCustomEvent)

	// HTMLDocument exists as a separate class for historical reasons, but it
	// can be treated merely as an alias for Document. In Firefox, there is an
	// inheritance relationship between the two, which is modelled here.
	//
	// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLDocument
	js.RegisterClass(e, "HTMLDocument", "Document", html.NewHTMLDocument)

	js.RegisterClass(e, "ShadowRoot", "DocumentFragment", NewUnconstructable)
	for _, cls := range codec.HtmlElements {
		if _, ok := e.Class(cls); !ok && cls != "HTMLElement" {
			js.RegisterClass(e, cls, "HTMLElement", NewUnconstructable)
		}
	}
}
