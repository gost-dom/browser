package log

import (
	"context"
	"errors"
	"io"
	"log/slog"

	"github.com/gost-dom/v8go"
)

type LogSource interface{ Logger() *slog.Logger }
type Logger = *slog.Logger

type LoggerLogSource struct{ L *slog.Logger }

func (s LoggerLogSource) Logger() Logger { return s.L }

var defaultLogger Logger

func SetDefault(logger *slog.Logger) {
	defaultLogger = logger
}

func Default() *slog.Logger {
	if defaultLogger == nil {
		defaultLogger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	return defaultLogger
}

type nullHandler struct{}

func (_ nullHandler) Enabled(context.Context, slog.Level) bool  { return false }
func (_ nullHandler) Handle(context.Context, slog.Record) error { return nil }
func (_ nullHandler) WithAttrs([]slog.Attr) slog.Handler        { return nullHandler{} }
func (_ nullHandler) WithGroup(name string) slog.Handler        { return nullHandler{} }

func init() {
	defaultLogger = slog.New(nullHandler{})
}

func logger(source Logger) *slog.Logger {
	if source != nil {
		return source
	} else {
		return defaultLogger
	}
}

func Info(source Logger, msg string, args ...any) {
	logger(source).Info(msg, args...)
}

func Warn(source Logger, msg string, args ...any) {
	logger(source).Warn(msg, args...)
}

func Debug(source Logger, msg string, args ...any) {
	logger(source).Debug(msg, args...)
}

func Error(source Logger, msg string, args ...any) {
	logger(source).Error(msg, args...)
}

// ErrAttr creates a log record attribute representing an error. If the error
// originates from JavaScript, location and stack trace are included in the log
// record.
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

			slog.Group("err", attrs...)
		}
	}
	return slog.Any("err", err)
}
