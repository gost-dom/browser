package gosttest_test

import (
	"context"
	"log/slog"
	"maps"
	"slices"
	"testing"

	"github.com/gost-dom/browser/testing/gosttest"
	"github.com/stretchr/testify/assert"
)

type logRecordRecorder struct {
	records []slog.Record
}

func (l *logRecordRecorder) Handle(ctx context.Context, r slog.Record) error {
	l.records = append(l.records, r)
	return nil
}

func (r *logRecordRecorder) Enabled(ctx context.Context, lvl slog.Level) bool { return true }

func TestWithAttributes(t *testing.T) {
	recorder := &logRecordRecorder{}
	handler := &gosttest.FlattenedHandler{Handler: recorder}
	l := slog.New(handler)
	l = l.With(slog.String("foo", "foo-value"))
	l.Info("Hello", "bar", "bar-value")

	assert.Equal(t, 1, len(recorder.records))
	m := make(map[string]slog.Value)
	recorder.records[0].Attrs(func(a slog.Attr) bool {
		m[a.Key] = a.Value
		return true
	})

	assert.Equal(t, "foo-value", m["foo"].String())
	assert.Equal(t, "bar-value", m["bar"].String())
}

func TestWithGroup(t *testing.T) {
	recorder := &logRecordRecorder{}
	handler := &gosttest.FlattenedHandler{Handler: recorder}
	l := slog.New(handler)
	l = l.WithGroup("grp")
	l.Info("Hello", "bar", "bar-value")

	assert.Equal(t, 1, len(recorder.records))
	m := make(map[string]slog.Value)
	recorder.records[0].Attrs(func(a slog.Attr) bool {
		m[a.Key] = a.Value
		return true
	})

	t.Log("Keys", slices.Collect(maps.Keys(m)))
	grpAttr, ok := m["grp"]
	assert.True(t, ok, "grp exists")

	assert.Equal(t, slog.KindGroup, grpAttr.Kind())
	grp := grpAttr.Group()
	assert.Equal(t, 1, len(grp))
	assert.Equal(t, "bar", grp[0].Key)
	assert.Equal(t, "bar-value", grp[0].Value.String())
}
