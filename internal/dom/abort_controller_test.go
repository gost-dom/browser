package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/dom"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	"github.com/stretchr/testify/assert"
)

func TestAbortControllerAbort(t *testing.T) {
	var eventDispatched bool
	ac := dom.NewAbortController()
	signal := ac.Signal()
	signal.AddEventListener("abort", eventtest.NewTestHandler(func(e *event.Event) {
		eventDispatched = true
	}))

	assert.False(t, eventDispatched, "event dispatched before Abort()")
	assert.False(t, signal.Aborted(), "signal aborted before Abort()")
	assert.Nil(t, signal.Reason(), "abort reason before Abort()")

	ac.Abort("Dummy reason")

	assert.True(t, eventDispatched, "event dispatched after Abort()")
	assert.True(t, signal.Aborted(), "signal aborted after Abort()")
	assert.Equal(t, "Dummy reason", signal.Reason(), "abort reason after Abort()")
}
