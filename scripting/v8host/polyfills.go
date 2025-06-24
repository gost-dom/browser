package v8host

import (
	_ "embed"
	"errors"
)

//go:embed polyfills/xpath/xpath.js
var xpath []byte

func installPolyfills(context *V8ScriptContext) error {
	errs := []error{
		context.Run(string(xpath)),
		context.Run(`
				const { XPathExpression, XPathResult } = window;
				const evaluate = XPathExpression.prototype.evaluate;
				XPathExpression.prototype.evaluate = function (context, type, res) {
					return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
				};
				Element.prototype.scrollIntoView = function() {};

		`),
	}
	return errors.Join(errs...)
}
