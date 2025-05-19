package gosttest

import (
	"context"
	"log/slog"
)

// LogRecorder is an implementation of slog.Handler recording all log entries in
// a slice.
type LogRecorder struct {
	Records []slog.Record
}

func (r *LogRecorder) Handle(_ context.Context, rec slog.Record) error {
	r.Records = append(r.Records, rec)
	return nil
}
func (r *LogRecorder) Enabled(context.Context, slog.Level) bool { return true }
func (r *LogRecorder) WithAttrs(attrs []slog.Attr) slog.Handler { return r }
func (r *LogRecorder) WithGroup(name string) slog.Handler       { return r }

// FilterLevel returns a slice of log records with the specified log level
func (r *LogRecorder) FilterLevel(lvl slog.Level) (ret []slog.Record) {
	for _, rec := range r.Records {
		if rec.Level == lvl {
			ret = append(ret, rec)
		}
	}
	return
}

func (r *LogRecorder) Reset() { r.Records = nil }
