package v8host

import (
	"github.com/gost-dom/browser/internal/uievents"
	js "github.com/gost-dom/browser/scripting/internal/js"
)

func (w uIEventV8Wrapper[T]) decodeMouseEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (eventInitWrapper, error) {
	return w.decodeUIEventInit(cbCtx, v)
}

func (w uIEventV8Wrapper[T]) decodePointerEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (eventInitWrapper, error) {
	return w.decodeMouseEventInit(cbCtx, v)
}

func (w uIEventV8Wrapper[T]) CreateInstance(
	cbCtx js.CallbackContext[T],
	type_ string,
) (js.Value[T], error) {
	return w.CreateInstanceEventInitDict(cbCtx, type_)
}

func (w uIEventV8Wrapper[T]) CreateInstanceEventInitDict(
	cbCtx js.CallbackContext[T],
	type_ string,
	options ...interface{}) (js.Value[T], error) {
	e := uievents.NewUIEvent(type_)
	return w.store(e, cbCtx)
}

func (w uIEventV8Wrapper[T]) decodeUIEventInit(
	cbCtx js.CallbackContext[T],
	v js.Value[T],
) (eventInitWrapper, error) {
	return w.decodeEventInit(cbCtx, v)
}

type mouseEventV8Wrapper[T any] struct {
	uIEventV8Wrapper[T]
}

type pointerEventV8Wrapper[T any] struct {
	mouseEventV8Wrapper[T]
}

func newMouseEventV8Wrapper[T any](host js.ScriptEngine[T]) mouseEventV8Wrapper[T] {
	return mouseEventV8Wrapper[T]{*newUIEventV8Wrapper(host)}
}

func newPointerEventV8Wrapper[T any](host js.ScriptEngine[T]) pointerEventV8Wrapper[T] {
	return pointerEventV8Wrapper[T]{newMouseEventV8Wrapper(host)}
}
