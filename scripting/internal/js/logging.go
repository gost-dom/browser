package js

import (
	"fmt"
	"log/slog"
	"strconv"
)

type jsValueLogger[T any] struct{ v Value[T] }

func (l jsValueLogger[T]) LogValue() slog.Value {
	val := l.v
	if val == nil || val.IsUndefined() {
		return slog.StringValue("undefined")
	}
	if val.IsNull() {
		return slog.StringValue("null")
	}
	if obj, ok := val.AsObject(); ok {
		if native := obj.NativeValue(); native != nil {
			return slog.GroupValue(
				slog.String("type", fmt.Sprintf("%T", native)),
				slog.Any("val", native),
			)
		}
		// Don't log the object directly, as custom toString() functions will
		// stack overflow.
		return slog.StringValue("JS Object")
	}
	return slog.AnyValue(val)
}

// thisLogValuer implements [slog.Valuer] rendering the "this" argument
type thisLogValuer[T any] struct{ ctx CallbackContext[T] }

func (l thisLogValuer[T]) LogValue() slog.Value {
	return jsValueLogger[T]{l.ctx.This()}.LogValue()
}

// argsLogValuer implements [slog.Valuer] rendering the called arguments
type argsLogValuer[T any] struct{ ctx CallbackContext[T] }

func (l argsLogValuer[T]) LogValue() slog.Value {
	args := l.ctx.Args()
	loggers := make([]slog.Attr, len(args))
	for i, a := range args {
		loggers[i] = slog.Any(strconv.Itoa(i), jsValueLogger[T]{a})
	}
	return slog.GroupValue(loggers...)
}

func LogAttr[T any](key string, val Value[T]) slog.Attr {
	return slog.Any(key, jsValueLogger[T]{val})
}

// ThisLogAttr creates an [slog.Attr] representing the JavaScript this value in
// a callback.
func ThisLogAttr[T any](cbCtx CallbackContext[T]) slog.Attr {
	return slog.Any("this", thisLogValuer[T]{cbCtx})
}

func ArgsLogAttr[T any](cbCtx CallbackContext[T]) slog.Attr {
	return slog.Any("args", argsLogValuer[T]{cbCtx})
}
