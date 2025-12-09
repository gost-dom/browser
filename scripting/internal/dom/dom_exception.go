package dom

import (
	"errors"

	dom "github.com/gost-dom/browser/dom"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

type domException[T any] struct{}

func newDOMException[T any](scriptHost js.ScriptEngine[T]) *domException[T] {
	return &domException[T]{}
}

func (w domException[T]) Constructor(info js.CallbackContext[T]) (js.Value[T], error) {
	return nil, nil
}

func (w domException[T]) Initialize(jsClass js.Class[T]) {
	jsClass.CreateInstanceAttribute("code", w.code, nil)
}
func (w domException[T]) code(cbCtx js.CallbackContext[T]) (res js.Value[T], err error) {
	val, err := js.As[error](cbCtx.Instance())
	if err != nil {
		return
	}
	var domErr dom.DOMError
	if !errors.As(val, &domErr) {
		return nil, nil
	}
	return cbCtx.NewInt32(int32(domErr.Code)), nil
}
