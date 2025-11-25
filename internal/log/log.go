package log

import (
	"errors"
	"io"
	"log/slog"

	"github.com/gost-dom/v8go"
)

type LogSource interface{ Logger() *slog.Logger }
type Logger = *slog.Logger

type LoggerLogSource struct{ L *slog.Logger }

func (s LoggerLogSource) Logger() Logger { return s.L }

var nullLogger = slog.New(slog.NewTextHandler(io.Discard, nil))
var defaultLogger = nullLogger

// Set a default [slog/Logger] instance to use in contexts where a specific
// logger has not been set. If no default logger is set, the default will
// discard all log messages.
func SetDefault(logger *slog.Logger) {
	if logger == nil {
		logger = nullLogger
	}
	defaultLogger = logger
}

// Default returns the default logger. This method is guaranteed to always
// return a value, even if the default has explicitly beed set to nil. If no
// default has been configured, or overriden by a nil value, the logger will
// discard all logged messages.
func Default() *slog.Logger {
	return defaultLogger
}

// ErrAttr creates a log record attribute representing an error.
//
// If the error originates from V8, the relevant JavaScript location and stack
// trace are included in the log record.
func ErrAttr(err error) slog.Attr {
	var jsError *v8go.JSError
	if errors.As(err, &jsError) {
		return slog.Group("err",
			"message", jsError.Message,
			"location", jsError.Location,
			"stackTrace", jsError.StackTrace,
		)
	}
	var exception *v8go.Exception
	if errors.As(err, &exception) {
		obj, isObj := exception.Value.AsObject()
		if isObj == nil {
			attrs := make([]any, 1, 8)
			attrs[0] = slog.Any("error", exception.Error())
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

			return slog.Group("err", attrs...)
		}
	}
	return slog.Any("err", err)
}
