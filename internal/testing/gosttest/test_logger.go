package gosttest

import (
	"context"
	"log/slog"
	"strings"
	"testing"
)

type TestingLogHandler struct {
	testing.TB
	AllowErrors bool
	MinLogLevel int
}

func (l TestingLogHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= slog.Level(l.MinLogLevel)
}

func (l TestingLogHandler) Handle(_ context.Context, r slog.Record) error {
	l.TB.Helper()
	var b strings.Builder
	r.Attrs(func(a slog.Attr) bool {
		b.WriteString("\t")
		b.WriteString(a.String())
		b.WriteString("\n")
		return true
	})
	if r.Level < slog.LevelError || l.AllowErrors {
		l.TB.Logf("%v: %s\n%s", r.Level, r.Message, b.String())
	} else {
		l.TB.Errorf("%v: %s\n%s", r.Level, r.Message, b.String())
	}
	return nil
}

func (l TestingLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return l }

func (l TestingLogHandler) WithGroup(name string) slog.Handler { return l }

type HandlerOption = func(*TestingLogHandler)

func MinLogLevel(lvl int) HandlerOption {
	return func(h *TestingLogHandler) { h.MinLogLevel = lvl }
}

func AllowErrors() HandlerOption { return func(h *TestingLogHandler) { h.AllowErrors = true } }

type TestLoggerOption = func(*TestingLogHandler)

func NewTestLogger(t testing.TB, opts ...TestLoggerOption) *slog.Logger {
	handler := TestingLogHandler{TB: t}
	for _, o := range opts {
		o(&handler)
	}
	return slog.New(handler)
}
