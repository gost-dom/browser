package v8host

import (
	"errors"
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func initDOMParser[T any](ft js.Constructor[T]) {
	ft.CreatePrototypeMethod("parseFromString",
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
			window := cbCtx.Scope().Window()
			html, err0 := consumeArgument(cbCtx, "html", nil, decodeString)
			contentType, err1 := consumeArgument(cbCtx, "contentType", nil, decodeString)
			if err := errors.Join(err0, err1); err != nil {
				return nil, err
			}
			if contentType != "text/html" {
				return nil, cbCtx.ValueFactory().NewTypeError(
					"DOMParser.parseFromString only supports text/html yet",
				)
			}
			domParser := NewDOMParser()
			var doc dom.Document
			if err := domParser.ParseReader(window, &doc, strings.NewReader(html)); err == nil {
				return encodeEntity[T](cbCtx, doc)
			} else {
				return nil, err
			}
		})

}

func installDOMParser[T any](host js.ScriptEngine[T]) {
	ctor := host.CreateClass(
		"DOMParser",
		nil,
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return nil, nil },
	)
	initDOMParser(ctor)
}

func init() {
	initializers = append(initializers, installDOMParser[jsTypeParam])
}
