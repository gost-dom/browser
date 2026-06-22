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
	host.InstallPolyfill(`
		installNode = (n) => {
			n.ELEMENT_NODE = 1;
			n.ATTRIBUTE_NODE = 2;
			n.TEXT_NODE = 3;
			n.CDATA_SECTION_NODE = 4;
			n.ENTITY_REFERENCE_NODE = 5;
			n.ENTITY_NODE = 6;
			n.PROCESSING_INSTRUCTION_NODE = 7;
			n.COMMENT_NODE = 8;
			n.DOCUMENT_NODE = 9;
			n.DOCUMENT_TYPE_NODE = 10;
			n.DOCUMENT_FRAGMENT_NODE = 11;
			n.NOTATION_NODE = 12;
			n.DOCUMENT_POSITION_DISCONNECTED = 0x01;
			n.DOCUMENT_POSITION_PRECEDING = 0x02;
			n.DOCUMENT_POSITION_FOLLOWING = 0x04;
			n.DOCUMENT_POSITION_CONTAINS = 0x08;
			n.DOCUMENT_POSITION_CONTAINED_BY = 0x10;
			n.DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 0x20;
	    }
		installNode(Node)
		installNode(Node.prototype)
	`, "gost-dom/polyfills/node.js")

	host.InstallPolyfill(string(xpath), "gost-dom/polyfills/xpath-jsdom.js")
	host.InstallPolyfill(string(textEncoderDecoder), "gost-dom/polyfills/text-encoder-decoder.js")
	host.InstallPolyfill(`
			const { XPathExpression, XPathResult } = window;
			const evaluate = XPathExpression.prototype.evaluate;
			XPathExpression.prototype.evaluate = function (context, type, res) {
				return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
			};
			Element.prototype.scrollIntoView = function() {};

	`, "gost-dom/polyfills/xpath-custom.js")
	host.InstallPolyfill(`
		Object.setPrototypeOf(DOMException, Error)
		Object.setPrototypeOf(DOMException.prototype, Error.prototype)
	`, "gost-dom/polyfills/errors.js")
}
