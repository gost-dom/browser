package v8host

import (
	"errors"
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/html"
	v8 "github.com/gost-dom/v8go"
)

func createDOMParserPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	constructor := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, nil
		},
	)
	prototype := constructor.PrototypeTemplate()
	prototype.Set(
		"parseFromString",
		wrapV8Callback(host, func(cbCtx *argumentHelper) (jsValue, error) {
			ctx := cbCtx.ScriptCtx()
			window := ctx.window
			html, err0 := cbCtx.consumeString()
			contentType, err1 := cbCtx.consumeString()
			if err := errors.Join(err0, err1); err != nil {
				return cbCtx.ReturnWithError(err)
			}
			if contentType != "text/html" {
				return cbCtx.ReturnWithTypeError(
					"DOMParser.parseFromString only supports text/html yet",
				)
			}
			domParser := NewDOMParser()
			var doc dom.Document
			if err := domParser.ParseReader(window, &doc, strings.NewReader(html)); err == nil {
				v, err := ctx.getJSInstance(doc)
				return cbCtx.ReturnWithJSValueErr(v, err)
			} else {
				return cbCtx.ReturnWithError(err)
			}
		}))
	return constructor
}
