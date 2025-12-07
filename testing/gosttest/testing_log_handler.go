package gosttest

import (
	"context"
	"log/slog"
	"strings"
	"testing"
)

// testingLogHandler provides an implementation of [slog.Handler] that writes
// log records to a [testing.TB] instance, keeping log records associated with
// the test in which they were generated.
//
// This type could reasonably be exported, but the design is likely to change,
// and you would (almost) always create a logger using NewTestingLogger.
type testingLogHandler struct {
	testing.TB

	allowErrors   bool
	minLogLevel   slog.Level
	testCompleted bool
}

// Enabled implements slog.Handler.
func (l testingLogHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= slog.Level(l.minLogLevel)
}

func (h testingLogHandler) testDone() bool {
	if h.testCompleted {
		return true
	}
	return h.TB.Context().Err() != nil
}

func (l testingLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if l.testDone() {
		return nil
	}
	l.TB.Helper()
	var b strings.Builder
	var w attrWriter
	r.Attrs(func(a slog.Attr) bool {
		if !a.Equal(slog.Attr{}) {
			w.write(&b, a)
		}
		return true
	})
	if r.Level < slog.LevelError || l.allowErrors {
		l.TB.Logf("%v: %s\n%s", r.Level, r.Message, b.String())
	} else {
		l.TB.Errorf("%v: %s\n%s", r.Level, r.Message, b.String())
	}
	return nil
}
