package html

import (
	"github.com/gost-dom/browser/internal/clock"
	htmlinterfaces "github.com/gost-dom/browser/internal/interfaces/html-interfaces"
	"github.com/gost-dom/browser/scripting/internal/codec"
	"github.com/gost-dom/browser/scripting/internal/dom"
	"github.com/gost-dom/browser/scripting/internal/js"
)

func decodeTimerHandler[T any](
	scope js.Scope[T], v js.Value[T],
) (res htmlinterfaces.TimerHandler, err error) {
	f, err := codec.DecodeFunction(scope, v)
	if err != nil {
		return nil, err
	}
	res = func() error {
		if _, err := f.Call(scope.GlobalThis()); err != nil {
			dom.HandleJSCallbackError(scope, "TimerHandler", err)
		}
		return nil
	}
	return res, nil
}

func decodeVoidFunction[T any](
	scope js.Scope[T], v js.Value[T],
) (res htmlinterfaces.TimerHandler, err error) {
	return decodeTimerHandler(scope, v)
}

func decodeTaskHandle[T any](
	s js.Scope[T], v js.Value[T],
) (clock.TaskHandle, error) {
	i, err := codec.DecodeInt(s, v)
	return clock.TaskHandle(i), err
}

func encodeTaskHandle[T any](s js.Scope[T], v clock.TaskHandle) (js.Value[T], error) {
	return codec.EncodeInt(s, int(v))
}
