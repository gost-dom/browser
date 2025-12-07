package gosttest

import (
	"log/slog"
	"testing"
)

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
		handler.testCompleted = true
	})
	return slog.New(flattenedHandler{handler, "", nil})
}

// attrWrite writes a single [slog.Attr] to a [strings.Builder] in the form
// "key=value". Groups are rendered in a hirarchical structure
