package v8engine

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/gost-dom/v8go"
)

type V8Error struct {
	*v8Value
	exception *v8go.Exception
	error
}

func newV8Error(ctx *V8ScriptContext, err error) *V8Error {
	exc := v8go.NewError(ctx.iso(), err.Error())
	val := &v8Value{ctx, exc.Value}
	return &V8Error{val, exc, err}
}

func (e V8Error) LogValue() slog.Value {
	err := e.error
	var jsError *v8go.JSError
	var errType = fmt.Sprintf("%T", err)
	if errors.As(err, &jsError) {
		return slog.GroupValue(
			slog.String("message", jsError.Message),
			slog.String("location", jsError.Location),
			slog.String("stackTrace", jsError.StackTrace),
			slog.String("errType", errType),
		)
	}
	var exception *v8go.Exception
	if errors.As(err, &exception) {
		obj, isObj := exception.Value.AsObject()
		if isObj == nil {
			attrs := make([]slog.Attr, 2, 9)
			attrs[0] = slog.Any("message", exception.Error())
			attrs[1] = slog.String("errType", errType)
			addValue := func(key string) {
				if val, err := obj.Get(key); err == nil {
					attrs = append(attrs, slog.Any(key, val))
				}
			}
			addValue("message")
			addValue("cause")
			addValue("fileName")
			addValue("lineNumber")
			addValue("columnNumber")
			addValue("name")
			addValue("stack")

			return slog.GroupValue(attrs...)
		}
	}
	return slog.GroupValue(slog.Any(
		"message", err),
		slog.String("errType", errType),
	)

}
