package htmlsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
)

func testMessageChannel(t *testing.T, e html.ScriptEngine) {
	w := browsertest.InitWindow(t, e)
	w.MustRun(`
		const ch = new MessageChannel()
	`)
}
