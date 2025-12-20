package htmlsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
)

func testHtmlDocument(t *testing.T, e html.ScriptEngine) {
	win := browsertest.InitWindow(t, e)
	t.Run("Location is null on new document", func(t *testing.T) {
		win.MustRun(`
			gost.assertNotNull(document.location)
			gost.assertNull(new Document().location)	
		`)
	})
}
