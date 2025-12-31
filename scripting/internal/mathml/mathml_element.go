package mathml

import (
	"github.com/gost-dom/browser/scripting/internal/html"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type MathMLElement[T any] struct {
	htmlOrSVGElement html.HTMLOrSVGElement[T]
}

func NewMathMLElement[T any](scriptHost js.ScriptEngine[T]) MathMLElement[T] {
	return MathMLElement[T]{html.NewHTMLOrSVGElement(scriptHost)}
}

func (wrapper MathMLElement[T]) Initialize(jsClass js.Class[T]) {
	wrapper.installPrototype(jsClass)
}

func (w MathMLElement[T]) installPrototype(jsClass js.Class[T]) {
	w.htmlOrSVGElement.InstallPrototype(jsClass)
}

func (w MathMLElement[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	cbCtx.Logger().Debug("JS Function call: MathMLElement.Constructor")
	return cbCtx.ReturnWithTypeError("Illegal constructor")
}
