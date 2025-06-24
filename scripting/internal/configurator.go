package internal

import (
	_ "embed"

	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/html"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/scripting/internal/uievents"
	"github.com/gost-dom/browser/scripting/internal/url"
	"github.com/gost-dom/browser/scripting/internal/xhr"
)

func Configure[T any](host js.ScriptEngine[T]) {
	dom.Configure(host)
	fetch.Configure(host)
}

func Bootstrap[T any](reg js.ClassBuilder[T]) {
	dom.Bootstrap(reg)
	html.InitBuilder(reg)
	xhr.Bootstrap(reg)
	url.Bootstrap(reg)
	uievents.Bootstrap(reg)
	fetch.Bootstrap(reg)

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

//go:embed polyfills/xpath/xpath.js
var xpath []byte

func InstallPolyfills[T any](host js.ScriptEngine[T]) {
	host.RunScript(`
		FormData.prototype.forEach = function(cb) {
			return Array.from(this).forEach(([k,v]) => { cb(v,k) })
		}
	`, "gost-dom/polyfills/formdata.js")
	host.RunScript(`
		Node.ELEMENT_NODE = 1;
		Node.ATTRIBUTE_NODE = 2;
		Node.TEXT_NODE = 3;
		Node.CDATA_SECTION_NODE = 4;
		Node.ENTITY_REFERENCE_NODE = 5;
		Node.ENTITY_NODE = 6;
		Node.PROCESSING_INSTRUCTION_NODE = 7;
		Node.COMMENT_NODE = 8;
		Node.DOCUMENT_NODE = 9;
		Node.DOCUMENT_TYPE_NODE = 10;
		Node.DOCUMENT_FRAGMENT_NODE = 11;
		Node.NOTATION_NODE = 12;
		Node.DOCUMENT_POSITION_DISCONNECTED = 0x01;
		Node.DOCUMENT_POSITION_PRECEDING = 0x02;
		Node.DOCUMENT_POSITION_FOLLOWING = 0x04;
		Node.DOCUMENT_POSITION_CONTAINS = 0x08;
		Node.DOCUMENT_POSITION_CONTAINED_BY = 0x10;
		Node.DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 0x20;
	`, "gost-dom/polyfills/node.js")

	host.RunScript(string(xpath), "gost-dom/polyfills/xpath-jsdom.js")
	host.RunScript(`
			const { XPathExpression, XPathResult } = window;
			const evaluate = XPathExpression.prototype.evaluate;
			XPathExpression.prototype.evaluate = function (context, type, res) {
				return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
			};
			Element.prototype.scrollIntoView = function() {};

	`, "gost-dom/polyfills/xpath-custom.js")
}
