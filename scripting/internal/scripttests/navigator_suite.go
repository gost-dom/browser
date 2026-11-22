package scripttests

import (
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

// testNavigator covers the configurable navigator surface: the default profile
// reports sane values, and the browser-level WithNavigator option overrides
// them.
func testNavigator(t *testing.T, e html.ScriptEngine) {
	t.Run("default profile", func(t *testing.T) {
		win := browsertest.InitWindow(t, e)
		eq := func(name, script string, want any) {
			t.Helper()
			assert.Equal(t, want, win.MustEval(script), name)
		}
		eq("userAgent non-empty", `navigator.userAgent.length > 0`, true)
		eq("platform non-empty", `navigator.platform.length > 0`, true)
		eq("vendor non-empty", `navigator.vendor.length > 0`, true)
		eq("language non-empty", `navigator.language.length > 0`, true)
		eq("languages is array", `Array.isArray(navigator.languages)`, true)
		eq("language matches languages[0]", `navigator.language === navigator.languages[0]`, true)
		eq("hardwareConcurrency positive",
			`typeof navigator.hardwareConcurrency === "number" && navigator.hardwareConcurrency > 0`, true)
		eq("webdriver false", `navigator.webdriver`, false)
	})

	t.Run("WithNavigator override", func(t *testing.T) {
		win := browsertest.InitWindow(t, e, browsertest.WithBrowserOption(
			browser.WithNavigator(html.NavigatorProfile{
				UserAgent: "Custom/1.0",
				Platform:  "TestOS",
				Languages: []string{"de-DE", "de"},
			}),
		))
		eq := func(name, script string, want any) {
			t.Helper()
			assert.Equal(t, want, win.MustEval(script), name)
		}
		eq("userAgent overridden", `navigator.userAgent`, "Custom/1.0")
		eq("platform overridden", `navigator.platform`, "TestOS")
		eq("languages overridden", `navigator.languages.join(",")`, "de-DE,de")
		eq("language defaults to languages[0]", `navigator.language`, "de-DE")
		// Unset fields fall back to the default profile.
		eq("vendor falls back to default", `navigator.vendor.length > 0`, true)
	})
}
