package v8host

import (
	"errors"
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/html"
)

func initDOMParser(ft jsConstructor) {
	ft.CreatePrototypeMethod("parseFromString",
		func(cbCtx jsCallbackContext) (jsValue, error) {
			ctx := cbCtx.ScriptCtx()
			window := ctx.window
			html, err0 := consumeArgument(cbCtx, "html", nil, decodeString)
			contentType, err1 := consumeArgument(cbCtx, "contentType", nil, decodeString)
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
		})

}

func installDOMParser(host *V8ScriptHost) {
	ctor := host.CreateClass(
		"DOMParser",
		nil,
		func(cbCtx jsCallbackContext) (jsValue, error) { return nil, nil },
	)
	initDOMParser(ctor)
}

func init() {
	initializers = append(initializers, installDOMParser)
}
