package dom

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func decodeCustomEventInit[T any](
	_ js.Scope[T],
	v js.Value[T],
	init *event.CustomEventInit,
) (err error) {
	if obj, ok := v.AsObject(); ok {
		init.Detail, err = obj.Get("detail")
	}
	return
}
