package dom

import (
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/scripting/internal/codec"
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

func encodeAny[T any](s js.Scope[T], v any) (js.Value[T], error) {
	return codec.EncodeAny(s, v)
}
