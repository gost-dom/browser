package internal

import (
	_ "embed"

	"github.com/gost-dom/browser/scripting/internal/js"
)

//go:embed polyfills/xpath/xpath.js
var xpath []byte

//go:embed polyfills/FastestSmallestTextEncoderDecoder/EncoderDecoderTogether.min.js
var textEncoderDecoder []byte

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
	host.RunScript(string(textEncoderDecoder), "gost-dom/polyfills/text-encoder-decoder.js")
	host.RunScript(`
			const { XPathExpression, XPathResult } = window;
			const evaluate = XPathExpression.prototype.evaluate;
			XPathExpression.prototype.evaluate = function (context, type, res) {
				return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
			};
			Element.prototype.scrollIntoView = function() {};

	`, "gost-dom/polyfills/xpath-custom.js")
	host.RunScript(`
		Object.setPrototypeOf(DOMException, Error)
		Object.setPrototypeOf(DOMException.prototype, Error.prototype)
	`, "gost-dom/polyfills/errors.js")
}
