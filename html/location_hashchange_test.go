package html_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

// Setting location.hash is a same-document navigation: it must normalize the
// fragment (strip one leading '#') and fire a hashchange event at the window.
func TestLocationHashSetterFiresHashChange(t *testing.T) {
	b := browsertest.InitBrowser(t, nil, nil)
	win := b.OpenWindow("http://example.com/index")

	fired := 0
	win.AddEventListener("hashchange", event.NewEventHandlerFunc(func(e *event.Event) error {
		fired++
		return nil
	}))

	win.Location().SetHash("#/login")

	// Normalized: reads back as "#/login", not "##/login".
	assert.Equal(t, "#/login", win.Location().Hash())
	assert.Equal(t, "http://example.com/index#/login", win.Location().Href(), "single # in href")
	assert.Equal(t, 1, fired, "hashchange should fire once")

	// Setting the same hash again fires nothing (no change).
	win.Location().SetHash("#/login")
	assert.Equal(t, 1, fired, "identical hash must not re-fire")

	// A different hash fires again.
	win.Location().SetHash("#/home")
	assert.Equal(t, 2, fired)
}

// The event carries oldURL/newURL per the HashChangeEvent interface.
func TestHashChangeEventURLs(t *testing.T) {
	b := browsertest.InitBrowser(t, nil, nil)
	win := b.OpenWindow("http://example.com/index")

	var got html.HashChangeEventInit
	win.AddEventListener("hashchange", event.NewEventHandlerFunc(func(e *event.Event) error {
		got, _ = e.Data.(html.HashChangeEventInit)
		return nil
	}))

	win.Location().SetHash("#/next")
	assert.Equal(t, "http://example.com/index", got.OldURL)
	assert.Equal(t, "http://example.com/index#/next", got.NewURL)
}
