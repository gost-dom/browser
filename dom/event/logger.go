package event

import (
	"log/slog"
)

func newLogAttrForEvent(e *Event) slog.Attr {
	return slog.Group("event",
		slog.String("type", e.Type),
		slog.Bool("cancelable", e.Cancelable),
		slog.Bool("bubbles", e.Bubbles),
		slog.String("phase", e.EventPhase.String()),
	)
}

func newLogAttrForListener(l EventListener) slog.Attr {
	return slog.Group("listener",
		slog.Bool("capture", l.Capture),
		slog.Bool("once", l.Once),
	)
}
