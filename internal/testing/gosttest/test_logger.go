package gosttest

import (
	"context"
	"log/slog"
	"strings"
	"testing"
)

type TestingLogHandler struct {
	testing.TB
	allowErrors bool
}

func (l TestingLogHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= slog.LevelInfo
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
	if r.Level < slog.LevelError || l.allowErrors {
		l.TB.Logf("%v: %s\n%s", r.Level, r.Message, b.String())
	} else {
		l.TB.Errorf("%v: %s\n%s", r.Level, r.Message, b.String())
	}
	return nil
}

func (l TestingLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return l }
func (l TestingLogHandler) WithGroup(name string) slog.Handler       { return l }

func NewTestLogger(t testing.TB) *slog.Logger {
	return slog.New(TestingLogHandler{TB: t})
}
