package gosttest

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"
)

type TestingLogHandler struct {
	testing.TB
	AllowErrors bool
	MinLogLevel slog.Level

	fallbackHandler slog.Handler
	testCompleted   bool
}

func (l TestingLogHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= slog.Level(l.MinLogLevel)
}

func (l TestingLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if l.testCompleted {
		if l.fallbackHandler == nil {
			l.fallbackHandler = slog.NewTextHandler(os.Stderr, nil)
		}
		fmt.Fprintf(os.Stderr, "gost-dom/gosttest: write to test logger after close")
		return l.fallbackHandler.Handle(ctx, r)
	}
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

// close prevents further log messages from being written to the wrapped
// [testing.TB] instance. Once a test has completed, logging will cause a panic.
func (l *TestingLogHandler) close() { l.testCompleted = true }

type HandlerOption = func(*TestingLogHandler)

func MinLogLevel(lvl slog.Level) HandlerOption {
	return func(h *TestingLogHandler) { h.MinLogLevel = lvl }
}

func AllowErrors() HandlerOption { return func(h *TestingLogHandler) { h.AllowErrors = true } }

type TestLoggerOption = func(*TestingLogHandler)

func NewTestLogger(t testing.TB, opts ...TestLoggerOption) *slog.Logger {
	handler := TestingLogHandler{TB: t}
	for _, o := range opts {
		o(&handler)
	}
	t.Cleanup(func() {
		handler.close()
	})
	return slog.New(handler)
}
