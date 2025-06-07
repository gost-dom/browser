package internal

import (
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/scripting/internal/url"
	"github.com/gost-dom/browser/scripting/internal/xhr"
)

func Configure[T any](host js.ScriptEngine[T]) {
	dom.Configure(host)
}

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	dom.Bootstrap(reg)
	html.Bootstrap(reg)
	xhr.Bootstrap(reg)
	url.Bootstrap(reg)

	// HTMLDocument exists as a separate class for historical reasons, but it
	// can be treated merely as an alias for Document. In Firefox, there is an
	// inheritance relationship between the two, which is modelled here.
	//
	// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLDocument
	js.RegisterClass(reg, "HTMLDocument", "Document", html.NewHTMLDocumentV8Wrapper)
}
