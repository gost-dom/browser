package v8host

import (
	_ "embed"
	"errors"
)

//go:embed polyfills/xpath/xpath.js
var xpath []byte

//go:embed polyfills/whatwg-fetch/fetch.js
var fetch []byte

//go:embed polyfills/abortcontroller/polyfill-patch-fetch.js
var abortController []byte

//go:embed polyfills/FastestSmallestTextEncoderDecoder/EncoderDecoderTogether.min.js
var encoding []byte

func installPolyfills(context *V8ScriptContext) error {
	installer := (*installer)(context)
	errs := []error{
		installer.installFormData(),
		installer.polyfillAnchor(),
		context.Run(string(xpath)),
		context.Run(string(fetch)),
		context.Run(string(abortController)),
		context.Run(string(encoding)),
		context.Run(`
				const { XPathExpression, XPathResult } = window;
				const evaluate = XPathExpression.prototype.evaluate;
				XPathExpression.prototype.evaluate = function (context, type, res) {
					return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
				};
				Element.prototype.scrollIntoView = function() {};

		`),
		installer.polyfillNode(),
	}
	return errors.Join(errs...)
}

type installer V8ScriptContext

func (i *installer) run(script string) error {
	return (*V8ScriptContext)(i).Run(script)
}

func (i *installer) polyfillAnchor() error {
	// TODO: This should eventually be generated
	return i.run(`
		HTMLAnchorElement.prototype.toString = function() { return this.href }
	`)
}

func (i *installer) installFormData() error {
	return i.run(`
		FormData.prototype.forEach = function(cb) {
			return Array.from(this).forEach(([k,v]) => { cb(v,k) })
		}
	`)
}

func (i *installer) polyfillNode() error {
	return i.run(`
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
	`)
}
