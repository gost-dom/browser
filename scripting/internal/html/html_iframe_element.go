package html

import (
	"github.com/gost-dom/browser/html"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// installIFrameElement wires the HTMLIFrameElement.contentWindow and
// contentDocument accessors onto the placeholder class created by
// installHtmlElementTypes (which must run first). They expose the iframe's
// nested browsing context.
func installIFrameElement[T any](e js.ScriptEngine[T]) {
	if iframe, ok := e.Class("HTMLIFrameElement"); ok {
		iframe.CreateAttribute("contentWindow", iframeContentWindow, nil)
		iframe.CreateAttribute("contentDocument", iframeContentDocument, nil)
	}
}

// iframeContentWindow returns the iframe's nested-browsing-context global
// object. Because the child Window's JS global is cached when its script
// context is created, EncodeEntity hands back that very object, a value living
// in the child realm, so reads like contentWindow.String resolve to the child
// realm's genuine native constructors.
func iframeContentWindow[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	el, err := js.As[html.HTMLIFrameElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.Null(), nil
	}
	w := el.ContentWindow()
	if w == nil {
		return cbCtx.Null(), nil
	}
	return codec.EncodeEntity(cbCtx, w)
}

// iframeContentDocument returns the document of the iframe's nested browsing
// context (resolved within the same realm as contentWindow).
func iframeContentDocument[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	el, err := js.As[html.HTMLIFrameElement](cbCtx.Instance())
	if err != nil {
		return cbCtx.Null(), nil
	}
	w := el.ContentWindow()
	if w == nil {
		return cbCtx.Null(), nil
	}
	// Read `document` off the child realm's global so the returned value also
	// belongs to the child realm (rather than a parent-realm wrapper).
	cw, err := codec.EncodeEntity(cbCtx, w)
	if err != nil {
		return nil, err
	}
	obj, ok := cw.AsObject()
	if !ok {
		return cbCtx.Null(), nil
	}
	return obj.Get("document")
}
