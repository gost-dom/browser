package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper Response[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return nil, fmt.Errorf("gost-dom/fetch: Response constructor not implemented")
}
