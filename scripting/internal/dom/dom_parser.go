package dom

import (
	"errors"
	"strings"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func initDOMParser[T any](ft js.Class[T]) {
	ft.CreatePrototypeMethod("parseFromString", domParserParseFromString)
}
func domParserParseFromString[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	win := cbCtx.Window()
	doc := html.NewEmptyHtmlDocument(win)
	html, err0 := js.ConsumeArgument(cbCtx, "html", nil, codec.DecodeString)
	contentType, err1 := js.ConsumeArgument(cbCtx, "contentType", nil, codec.DecodeString)
	if err := errors.Join(err0, err1); err != nil {
		return nil, err
	}
	if contentType != "text/html" {
		return nil, cbCtx.NewTypeError(
			"DOMParser.parseFromString only supports text/html yet",
		)
	}
	if err := dom.ParseDocument(doc, strings.NewReader(html)); err == nil {
		return codec.EncodeEntity(cbCtx, doc)
	} else {
		return nil, err
	}
}

func installDOMParser[T any](host js.ScriptEngine[T]) {
	ctor := host.CreateClass(
		"DOMParser",
		nil,
		func(cbCtx js.CallbackContext[T]) (js.Value[T], error) { return nil, nil },
	)
	initDOMParser(ctor)
}
