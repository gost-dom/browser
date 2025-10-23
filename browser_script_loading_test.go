package browser_test

import (
	"strings"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/stretchr/testify/assert"
)

func TestBrowserScriptLoading(t *testing.T) {
	server := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(`
			<!doctype html>
			<html>
				<head>
					<script src="script1.js"></script>
					<script src="script2.js" defer></script>
					<script src="script3.js" type="module"></script>
					<script src="script4.js"></script>
				<head>
				<body>
					<h1>Script Test Page</h1>
					<div id="target"></div>
				</body>
			</html>`),
		"/script1.js": gosttest.StaticJS(`
				window.events = window.events || []
				window.events.push("script 1 loaded")
				document.addEventListener("DOMContentLoaded", () => {
					window.events.push("script 1 DOMContentLoaded")
				})
				document.addEventListener("load", () => {
					document.getElementById("target").textContent = events.join(",")
				})
			`),
		"/script2.js": gosttest.StaticJS(`
				window.events = window.events || []
				window.events.push("script 2 loaded")
				document.addEventListener("DOMContentLoaded", () => {
					window.events.push("script 2 DOMContentLoaded")
				})
			`),
		"/script3.js": gosttest.StaticJS(`
				window.events = window.events || []
				window.events.push("script 3 loaded")
				document.addEventListener("DOMContentLoaded", () => {
					window.events.push("script 3 DOMContentLoaded")
				})
			`),
		"/script4.js": gosttest.StaticJS(`
				window.events = window.events || []
				window.events.push("script 4 loaded")
				document.addEventListener("DOMContentLoaded", () => {
					window.events.push("script 4 DOMContentLoaded")
				})
			`),
	}
	b := browser.New(
		browser.WithHandler(server),
		browser.WithLogger(gosttest.NewTestLogger(t)),
	)
	w, err := b.Open("/index.html")
	assert.NoError(t, err)

	content := w.Document().GetElementById("target").TextContent()
	parts := strings.Split(content, ",")
	assert.Equal(t, []string{
		"script 1 loaded",
		"script 4 loaded",
		"script 2 loaded",
		"script 3 loaded",
		"script 1 DOMContentLoaded",
		"script 4 DOMContentLoaded",
		"script 2 DOMContentLoaded",
		"script 3 DOMContentLoaded",
	}, parts)

}
