package svg

import (
	"github.com/gost-dom/browser/scripting/internal/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type SVGElement[T any] struct {
	htmlOrSVGElement html.HTMLOrSVGElement[T]
}

func NewSVGElement[T any](scriptHost js.ScriptEngine[T]) SVGElement[T] {
	return SVGElement[T]{html.NewHTMLOrSVGElement(scriptHost)}
}

func (wrapper SVGElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w SVGElement[T]) installPrototype(jsClass js.Class[T]) {
	w.htmlOrSVGElement.InstallPrototype(jsClass)
}

func (w SVGElement[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: SVGElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}
