package js

import "github.com/gost-dom/browser/dom/event"

func UnhandledError[T any](scope Scope[T], err error) {
	scope.Window().DispatchEvent(event.NewErrorEvent(err))
}
