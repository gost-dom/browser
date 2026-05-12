package dom

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/html/document"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func DOMImplementation_createHTMLDocument[T any](
	cbCtx js.CallbackContext[T],
) (res js.Value[T], err error) {
	title, found, err := js.ConsumeOptionalArg(cbCtx, "title", codec.DecodeString)
	if err != nil {
		return nil, err
	}

	var options []document.NewDocumentOption
	if found {
		options = append(options, document.WithTitle(title))
	}

	doc := html.NewValidHTMLDocument(nil, options...)
	return codec.EncodeEntity(cbCtx, doc)
}

func DOMImplementation_createDocument[T any](
	cbCtx js.CallbackContext[T],
) (res js.Value[T], err error) {
	doc := html.NewEmptyHtmlDocument(nil)
	return codec.EncodeEntity(cbCtx, doc)
}
