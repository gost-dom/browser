package html

import (
	"fmt"
	"time"

	"github.com/gost-dom/browser/internal/constants"
	codec "github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func ConfigureScriptEngine[T any](e js.ScriptEngine[T]) {
	InitializeDOMStringMap(js.CreateClass(e, "DOMStringMap", "", js.IllegalConstructor))
	window, _ := e.Class("Window")
	installEventLoopGlobals(window)

	// HTMLDocument exists as a separate class for historical reasons, but it
	// can be treated merely as an alias for Document. In Firefox, there is an
	// inheritance relationship between the two, which is modelled here.
	//
	// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLDocument
	js.CreateClass(e, "HTMLDocument", "Document", js.IllegalConstructor)
	installHtmlElementTypes(e)
}

// installHtmlElementTypes adds classes for all the HTML element types work which
// there isn't yet an explicit implementation. This enables JavaScript code to
// type check on element types even though there is no explicit implementation.
//
// E.g., there isn't an explicit implementation of heading elements, so the
// following code would fail if these elements are not created (true at the time of writing this):
//
//	const e = document.getElementById("heading")
//	if (e instanceof HTMLHeadingElement) { /* ... */ }
func installHtmlElementTypes[T any](e js.ScriptEngine[T]) {
	element, ok := e.Class("HTMLElement")
	if !ok {
		panic(fmt.Sprintf("HTMLElement not configured: %s", constants.BUG_ISSUE_URL))
	}
	for _, cls := range codec.HtmlElements {
		if _, ok := e.Class(cls); !ok && cls != "HTMLElement" {
			e.CreateClass(cls, element, js.IllegalConstructor)
		}
	}
}

func installEventLoopGlobals[T any](global js.Class[T]) {
	global.CreateOperation("requestAnimationFrame", requestAnimationFrame)
}

func requestAnimationFrame[T any](cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	f, err := js.ConsumeArgument(cbCtx, "fn", nil, codec.DecodeFunction)
	if err != nil {
		return nil, err
	}
	cbCtx.Clock().AddSafeTask(
		func() {
			if _, err := f.Call(cbCtx.GlobalThis()); err != nil {
				dom.HandleJSCallbackError(cbCtx, "requestAnimationFrame", err)
			}
		}, 10*time.Millisecond)
	return nil, err
}
