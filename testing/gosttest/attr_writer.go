package gosttest

import (
	"log/slog"
	"strings"
)

type attrWriter struct {
	groups []string
}

func (w attrWriter) write(b *strings.Builder, a slog.Attr) {
	value := a.Value.Resolve()
	if value.Kind() == slog.KindGroup {
		as := value.Group()
		if len(as) == 0 {
			return
		}
		w.writeGroup(b, a)
		return
	}
	w.writePrefix(b)
	b.WriteString(a.Key)
	b.WriteString(": ")
	b.WriteString(value.String())
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

	value := a.Value.Resolve()
	for _, a := range value.Group() {
		w.write(b, a)
	}
}
