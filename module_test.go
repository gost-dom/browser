package browser_test

import (
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	const indexHTML = `
		<!DOCTYPE html>
		<html>
			<head><script src="module.js" type="module"></script></head>
			<body>
				<h1>Module test</h1>
				<div id="tgt"></div>
			</body>
		</html>`
	const moduleJS = `
		// Verify that compilation doesn't fail on export
		export const foo = 43;
		document.addEventListener("DOMContentLoaded", () => {
			document.getElementById("tgt").textContent="CONTENT";
		})
	`
	server := gosttest.StaticFileServer{
		"https://example.com/index.html": [2]string{"text/html", indexHTML},
		"https://example.com/module.js":  [2]string{"text/javascript", moduleJS},
	}
	browser := browser.New(
		browser.WithHandler(server),
		browser.WithLogger(gosttest.NewTestLogger(t)),
	)
	win, err := browser.Open("https://example.com/index.html")
	assert.NoError(t, err)
	g := gomega.NewWithT(t)
	g.Expect(
		win.Document().GetElementById("tgt")).To(HaveTextContent("CONTENT"))
}
