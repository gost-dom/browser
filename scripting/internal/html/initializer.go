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
	Bootstrap(e)

	eventTarget, _ := e.Class("EventTarget")
	window := e.ConfigureGlobalScope("Window", eventTarget)
	InitializeWindow(window)
	InitializeDOMStringMap(js.CreateClass(e, "DOMStringMap", "", js.IllegalConstructor))
	installEventLoopGlobals(window)

	// HTMLDocument exists as a separate class for historical reasons, but it
	// can be treated merely as an alias for Document. In Firefox, there is an
	// inheritance relationship between the two, which is modelled here.
	//
	// See also: https://developer.mozilla.org/en-US/docs/Web/API/HTMLDocument
	dom.InitializeDocument(js.CreateClass(e, "HTMLDocument", "Document", js.IllegalConstructor))
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
