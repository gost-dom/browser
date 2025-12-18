package codec

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func GetWindow[T any](scope js.Scope[T]) (html.Window, error) {
	win, err := DecodeNativeValue[T, html.Window](scope, scope.GlobalThis())
	return win, err
}
