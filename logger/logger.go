// Package logger provides the basic functionality of supplying a custom logger.
//
// gost-dom will write log messages to an [slog.Logger] instance, but the default
// logger discards all messages. You can provide your own logger, allowing
// control over where log messages appear.
package logger

import (
	"log/slog"

	"github.com/gost-dom/browser/internal/log"
)

// SetDefault sets the [slog.Logger] that will receive log messages from the
// server.
func SetDefault(logger *slog.Logger) {
	log.SetDefault(logger)
}
