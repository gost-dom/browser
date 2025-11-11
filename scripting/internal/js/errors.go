package js

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/log"
)

// HandleJSCallbackError is to be called when calling into a JS callback function
// results in an error. Argument cbType represent the type of callback, e.g.,
// event handler, mutation observer, interval, etc.
func HandleJSCallbackError[T any](scope Scope[T], cbType string, err error) {
	scope.Logger().Error("Callback error", "callback-type", cbType, log.ErrAttr(err))
	scope.Window().DispatchEvent(event.NewErrorEvent(err))
}
