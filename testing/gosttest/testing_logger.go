package gosttest

import (
	"context"
	"fmt"
	"log/slog"
	"os"
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

	allowErrors     bool
	minLogLevel     slog.Level
	fallbackHandler slog.Handler
	testCompleted   bool
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
		if l.fallbackHandler == nil {
			l.fallbackHandler = slog.NewTextHandler(os.Stderr, nil)
		}
		fmt.Fprintf(os.Stderr, "gost-dom/gosttest: write to test logger after close")
		return l.fallbackHandler.Handle(ctx, r)
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

// WithAttrs implements slog.Handler
func (l testingLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return l }

// WithGroup implements slog.Handler
func (l testingLogHandler) WithGroup(name string) slog.Handler { return l }

// close prevents further log messages from being written to the wrapped
// [testing.TB] instance. Once a test has completed, logging will cause a panic.
func (l *testingLogHandler) close() {
	l.testCompleted = true
}

// TestingLogHandlerOption is the type of for functional options when creating
// test loggers using [NewTestingLogger]
type TestingLogHandlerOption = func(*testingLogHandler)

// MinLogLevel controls which log level is written. Default is [slog.LevelInfo]
func MinLogLevel(lvl slog.Level) TestingLogHandlerOption {
	return func(h *testingLogHandler) { h.minLogLevel = lvl }
}

// AllowErrors prevents error level log records from automatically failing a
// test.
func AllowErrors() TestingLogHandlerOption {
	return func(h *testingLogHandler) { h.allowErrors = true }
}

// NewTestingLogger creates an [*slog.Logger] that sends log records to t.
//
// By default, log levels of [slog.LevelError] or above are reported using
// t.Error. Levels are reported using t.Log. Error reporting can be suppressed
// with the [AllowErrors] option.
//
// By default, log levels below [slog.LevelInfo] are supressed. This can be
// customized with the [MinLogLevel] option.
func NewTestingLogger(t testing.TB, opts ...TestingLogHandlerOption) *slog.Logger {
	handler := testingLogHandler{TB: t}
	for _, o := range opts {
		o(&handler)
	}
	t.Cleanup(func() {
		handler.close()
	})
	return slog.New(handler)
}

// attrWrite writes a single [slog.Attr] to a [strings.Builder] in the form
// "key=value". Groups are rendered in a hirarchical structure
type attrWriter struct {
	groups []string
}

func (w attrWriter) write(b *strings.Builder, a slog.Attr) {
	if a.Value.Kind() == slog.KindGroup {
		as := a.Value.Group()
		if len(as) == 0 {
			return
		}
		w.writeGroup(b, a)
		return
	}
	w.writePrefix(b)
	b.WriteString(a.Key)
	b.WriteString(": ")
	b.WriteString(a.Value.String())
	b.WriteString("\n")
}

func (w attrWriter) writePrefix(b *strings.Builder) {
	b.WriteString("\t")

	// An alternate implementation, writes groups as a dotted list, e.g.
	// instead of:
	//
	// 	res
	// 		- status=200
	// 		- header
	// 			- Set-Cookie=...
	//
	// It will write
	//
	// 	res.status=200
	// 	res.header.Set-Cookie=...

	// for _, g := range w.groups {
	// 	b.WriteString(g)
	// 	b.WriteByte('.')
	// }

	if len(w.groups) == 0 {
		return
	}
	for range w.groups {
		b.WriteString("\t")
	}
	b.WriteString("- ")
}

func (w attrWriter) writeGroup(b *strings.Builder, a slog.Attr) {
	// Write group header - remove if using the alternate version
	w.writePrefix(b)
	b.WriteString(a.Key)
	b.WriteString("\n")

	g := w.groups
	w.groups = append(w.groups, a.Key)
	defer func() { w.groups = g }()

	for _, a := range a.Value.Group() {
		w.write(b, a)
	}
}
