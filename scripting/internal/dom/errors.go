package dom

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/log"
	"github.com/gost-dom/browser/scripting/internal/codec"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

// HandleJSCallbackError is to be called when calling into a JS callback function
// results in an error. Argument cbType represent the type of callback, e.g.,
// event handler, mutation observer, interval, etc.
func HandleJSCallbackError[T any](scope js.Scope[T], cbType string, err error) {
	scope.Logger().Error("Callback error", "callback-type", cbType, log.ErrAttr(err))

	if target, err := codec.GetWindow(scope); err == nil {
		target.DispatchEvent(event.NewErrorEvent(err))
	}
}
