package gosttest

import (
	"context"
	"log/slog"
)

var _ slog.Handler = flattenedHandler{}

// simpleLogHandler defines a subset of [slog.Handler] where WithAttrs and
// WithGroup aren't included.
type simpleLogHandler interface {
	Enabled(context.Context, slog.Level) bool
	Handle(context.Context, slog.Record) error
}

// flattenedHandler is a an [slog.Handler] implementation that is a proxy for
// [simpleLogHandler]. flattenedHandler implements WithAttrs and WithGroup in a
// way that they will be remembered and merged with the final log records
// handled by the wrapped handler.
//
// The downside of this is that it bypass a performance optimization of slog,
// that permits the handler to pre-format attributes shared by many log
// statements.
type flattenedHandler struct {
	Handler simpleLogHandler
	group   string
	attrs   []slog.Attr
}

// Enabled implements [slog.Handler]
func (l flattenedHandler) Handle(ctx context.Context, r slog.Record) error {
	if len(l.attrs) > 0 {
		r = r.Clone()
		r.AddAttrs(l.attrs...)
	}

	if l.group != "" {
		attrs := make([]slog.Attr, r.NumAttrs())
		i := 0
		r.Attrs(func(a slog.Attr) bool {
			attrs[i] = a
			i++
			return true
		})
		r = slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
		r.Add(slog.GroupAttrs(l.group, attrs...))
	}
	return l.Handler.Handle(ctx, r)
}

// Enabled implements [slog.Handler]
func (l flattenedHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return flattenedHandler{
		l.Handler,
		l.group,
		append(l.attrs, attrs...),
	}
}

// Enabled implements [slog.Handler]
func (h flattenedHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return h.Handler.Enabled(ctx, lvl)
}

// Enabled implements [slog.Handler]
func (l flattenedHandler) WithGroup(name string) slog.Handler {
	return flattenedHandler{l, name, nil}
}
