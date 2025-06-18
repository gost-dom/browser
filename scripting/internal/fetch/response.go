package fetch

import "github.com/gost-dom/browser/scripting/internal/js"

func (wrapper Response[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return nil, nil
}
