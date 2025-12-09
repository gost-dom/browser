package log

import (
	"io"
	"log/slog"
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
// Different error types can implement their own LogValue() implementation,
// e.g., the V8 error does that, attempting to extract call stact etc.
func ErrAttr(err error) slog.Attr {
	// TODO: This used to contain logic, but not anymore. Candidate for
	// deletion. Only value now is consistent record key.
	return slog.Any("err", err)
}

// ReplaceStackAttr removes "stack" entries from log output. While stack output
// can be beneficial in some scenarios, it's very verbose.
//
// There is also a security consideration, stack output can expose details about
// the inner workings of the system, that can be exploited to find weaknesses.
// The intended use case of this library is not production use, there are still
// risks log messages leak, e.g., build logs.
func ReplaceStackAttr(grps []string, attr slog.Attr) slog.Attr {
	if attr.Key == "stack" {
		return slog.Attr{}
	}
	return attr
}
