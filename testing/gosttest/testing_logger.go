package gosttest

import (
	"log/slog"
	"testing"

	"github.com/gost-dom/browser/internal/log"
)

type ReplaceAttr func(groups []string, attr slog.Attr) slog.Attr

func (a ReplaceAttr) Chain(r ReplaceAttr) ReplaceAttr {
	if a == nil {
		return r
	}
	return func(grps []string, attr slog.Attr) slog.Attr {
		return a(grps, r(grps, attr))
	}
}

type testLogHandlerOptions struct {
	minLogLevel slog.Level
	allowErrors bool
	replaceAttr ReplaceAttr
}

// TestingLogHandlerOption is the type of for functional options when creating
// test loggers using [NewTestingLogger]
type TestingLogHandlerOption = func(*testLogHandlerOptions)

// MinLogLevel controls which log level is written. Default is [slog.LevelInfo]
func MinLogLevel(lvl slog.Level) TestingLogHandlerOption {
	return func(h *testLogHandlerOptions) { h.minLogLevel = lvl }
}

// AllowErrors prevents error level log records from automatically failing a
// test.
func AllowErrors() TestingLogHandlerOption {
	return func(h *testLogHandlerOptions) { h.allowErrors = true }
}

func WithReplaceAttr(r ReplaceAttr) TestingLogHandlerOption {
	return func(h *testLogHandlerOptions) { h.replaceAttr = h.replaceAttr.Chain(r) }
}

func WithDisableStack() TestingLogHandlerOption {
	return WithReplaceAttr(log.ReplaceStackAttr)
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
	var options testLogHandlerOptions
	for _, o := range opts {
		o(&options)
	}
	handler := testingLogHandler{TB: t}
	handler.minLogLevel = options.minLogLevel
	handler.allowErrors = options.allowErrors
	t.Cleanup(func() {
		handler.testCompleted = true
	})
	rootHandler := flattenedHandler{Handler: handler}
	rootHandler.replaceAttr = options.replaceAttr
	return slog.New(rootHandler)
}

// attrWrite writes a single [slog.Attr] to a [strings.Builder] in the form
// "key=value". Groups are rendered in a hirarchical structure
