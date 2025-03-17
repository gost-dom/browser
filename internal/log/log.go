package log

import (
	"context"
	"log/slog"
)

type LogSource interface{ Logger() *slog.Logger }
type Logger = *slog.Logger

type LoggerLogSource struct{ L *slog.Logger }

func (s LoggerLogSource) Logger() Logger { return s.L }

var defaultLogger Logger

func SetDefault(logger *slog.Logger) {
	defaultLogger = logger
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
