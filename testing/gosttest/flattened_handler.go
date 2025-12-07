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
	Handler     simpleLogHandler
	group       string
	groups      []string
	attrs       []slog.Attr
	replaceAttr ReplaceAttr
}

// Enabled implements [slog.Handler]
func (h flattenedHandler) Handle(ctx context.Context, r slog.Record) error {
	// if h.attrs == nil && h.group == "" {
	// 	return h.Handler.Handle(ctx, r)
	// }
	if len(h.attrs) > 0 {
		r = r.Clone()
		// Attrs have already been replaced at WithAttrs
		r.AddAttrs(h.attrs...)
	}

	if h.group == "" && h.replaceAttr == nil {
		return h.Handler.Handle(ctx, r)
	}

	attrs := getRecordAttrs(r)
	if h.replaceAttr != nil {
		for i, a := range attrs {
			attrs[i] = h.replaceAttr(h.groups, a)
		}
	}
	r = slog.NewRecord(r.Time, r.Level, r.Message, r.PC)
	if h.group == "" {
		r.AddAttrs(attrs...)
	} else {
		r.AddAttrs(slog.GroupAttrs(h.group, attrs...))
	}
	return h.Handler.Handle(ctx, r)
}

// Enabled implements [slog.Handler]
func (h flattenedHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if h.replaceAttr != nil {
		for i, a := range attrs {
			attrs[i] = h.replaceAttr(h.groups, a)
		}
	}
	// Take advantage of non-pointer receiver
	h.attrs = append(h.attrs, attrs...)
	return h
}

// Enabled implements [slog.Handler]
func (h flattenedHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return h.Handler.Enabled(ctx, lvl)
}

// Enabled implements [slog.Handler]
func (h flattenedHandler) WithGroup(name string) slog.Handler {
	grps := h.groups
	l := len(grps)
	h.Handler = h
	h.group = name
	h.groups = make([]string, l+1)
	copy(h.groups, grps)
	h.groups[l] = name
	return h
}

// getRecordAttrs iterates the attributes of an [slog.Record] and creates a new
// slice of the log attributes.
func getRecordAttrs(r slog.Record) []slog.Attr {
	attrs := make([]slog.Attr, r.NumAttrs())
	i := 0

	r.Attrs(func(a slog.Attr) bool {
		attrs[i] = a
		i++
		return true
	})
	return attrs
}
