package html

import (
	"errors"

	"github.com/gost-dom/browser/scripting/internal/js"
)

func (e HTMLTemplateElement[T]) CreateInstance(
	ctx js.ScriptEngine[T],
	this js.Object[T],
) (js.Value[T], error) {
	return nil, errors.New("TODO")
}
