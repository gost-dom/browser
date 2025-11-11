package fetch

import (
	"fmt"

	"github.com/gost-dom/browser/internal/fetch"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func (wrapper Response[T]) Constructor(cbCtx js.CallbackContext[T]) (js.Value[T], error) {
	return nil, fmt.Errorf("gost-dom/fetch: Response constructor not implemented")
}

func (wrapper Response[T]) toHeaders(
	ctx js.CallbackContext[T],
	headers fetch.Headers,
) (js.Value[T], error) {
	ctor := ctx.Constructor("Headers")
	return ctor.NewInstance(headers)
}
