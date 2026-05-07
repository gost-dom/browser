package cssomviewsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func RunDomSuite(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e, browsertest.WithHtml(
		`<body><div id="target"></div></body>`,
	))
	win.MustRun(`
		const target = document.getElementById("target")
		const rect = target.getBoundingClientRect()
	`)
	assert.Equal(t, "number", win.MustEval(`typeof rect.x`))
	assert.Equal(t, "number", win.MustEval(`typeof rect.y`))
}
