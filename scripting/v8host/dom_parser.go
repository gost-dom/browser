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
		v8.NewFunctionTemplateWithError(
			iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.mustGetContext(info.Context())
				window := ctx.window
				args := newArgumentHelper(host, info)
				html, err0 := args.consumeString()
				contentType, err1 := args.consumeString()
				if err := errors.Join(err0, err1); err != nil {
					return nil, err
				}
				if contentType != "text/html" {
					return nil, v8.NewTypeError(host.iso,
						"DOMParser.parseFromString only supports text/html yet",
					)
				}
				domParser := NewDOMParser()
				var doc dom.Document
				if err := domParser.ParseReader(window, &doc, strings.NewReader(html)); err == nil {
					return ctx.getJSInstance(doc)
				} else {
					return nil, err
				}
			},
		),
	)
	return constructor
}
