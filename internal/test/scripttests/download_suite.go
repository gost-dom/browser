package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func RunDownloadScriptSuite(t *testing.T, e html.ScriptEngine) {
	t.Run("Load script from source", func(t *testing.T) {
		const indexHTML = `
		<!DOCTYPE html>
		<html>
			<head><script src="module.js"></script></head>
			<body>
				<h1>Module test</h1>
				<div id="tgt"></div>
			</body>
		</html>`
		const moduleJS = `
			document.addEventListener("DOMContentLoaded", () => {
				document.getElementById("tgt").textContent="CONTENT";
			})
		`
		server := gosttest.HttpHandlerMap{
			"/index.html": gosttest.StaticHTML(indexHTML),
			"/module.js":  gosttest.StaticJS(moduleJS),
		}
		win, err := initBrowser(t, e, server).Open("https://example.com/index.html")
		assert.NoError(t, err)
		g := gomega.NewWithT(t)
		g.Expect(win.Document().GetElementById("tgt")).To(HaveTextContent("CONTENT"))

	})
}
