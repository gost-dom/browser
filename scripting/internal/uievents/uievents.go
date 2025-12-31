package uievents

import (
	"github.com/gost-dom/browser/internal/uievents"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func decodePointerEventInit[T any](
	_ js.Scope[T], _ js.Object[T], _ *uievents.PointerEventInit,
) error {
	return nil
}

func decodeMouseEventInit[T any](_ js.Scope[T], _ js.Object[T], _ *uievents.MouseEventInit) error {
	return nil
}

func decodeUIEventInit[T any](_ js.Scope[T], _ js.Object[T], _ *uievents.UIEventInit) error {
	return nil
}
