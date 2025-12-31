package svg

import (
	"github.com/gost-dom/browser/scripting/internal/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func InitializeSvgElement[T any](jsClass js.Class[T]) {
	html.InitializeHTMLOrSVGElement(jsClass)
}
