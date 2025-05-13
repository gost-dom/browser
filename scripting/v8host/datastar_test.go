package v8host_test

import (
	"log/slog"
	"testing"

	"github.com/gost-dom/browser"
	app "github.com/gost-dom/browser/internal/test/integration/test-app"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

func TestDatastar(t *testing.T) {
	server := app.CreateServer()
	b := htmltest.NewBrowserHelper(t, browser.New(
		browser.WithHandler(server),
		browser.WithLogger(gosttest.NewTestLogger(t, gosttest.MinLogLevel(slog.LevelInfo))),
	))
	win := b.OpenWindow("https://example.com/ds")
	win.HTMLDocument().GetHTMLElementById("fetch-events-button").Click()
	t.Log(win.HTMLDocument().Body().OuterHTML())
}
