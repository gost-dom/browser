package internal

import (
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
}

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	dom.Bootstrap(reg)
	html.InitBuilder(reg)
	svg.Bootstrap(reg)
	mathml.Bootstrap(reg)
	xhr.Bootstrap(reg)
	url.Bootstrap(reg)
	uievents.Bootstrap(reg)
	fetch.Bootstrap(reg)
	streams.Bootstrap(reg)

	js.RegisterClass(reg, "File", "", dom.NewEvent)
	js.RegisterClass(reg, "CustomEvent", "Event", dom.NewCustomEvent)

	// HTMLDocument exists as a separate class for historical reasons, but it
	// can be treated merely as an alias for Document. In Firefox, there is an
	// inheritance relationship between the two, which is modelled here.
	//
	// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLDocument
	js.RegisterClass(reg, "HTMLDocument", "Document", html.NewHTMLDocument)

	js.RegisterClass(reg, "ShadowRoot", "DocumentFragment", NewUnconstructable)
	for _, cls := range codec.HtmlElements {
		if !reg.HasClass(cls) && cls != "HTMLElement" {
			js.RegisterClass(reg, cls, "HTMLElement", NewUnconstructable)
		}
	}
}
