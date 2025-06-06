package html

import (
	"errors"

	"github.com/gost-dom/browser/scripting/internal/js"
)

func (e HTMLTemplateElementV8Wrapper[T]) CreateInstance(
	ctx js.ScriptEngine[T],
	this js.Object[T],
) (js.Value[T], error) {
	return nil, errors.New("TODO")
}
