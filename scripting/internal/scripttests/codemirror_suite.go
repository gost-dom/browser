package scripttests

import (
	"log/slog"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/input/controller"
	app "github.com/gost-dom/browser/internal/test/integration/test-app"
	"github.com/gost-dom/browser/internal/testing/browsertest"

	"github.com/stretchr/testify/assert"
)

func RunCodeMirrorTests(t *testing.T, e html.ScriptEngine) {
	t.Run("Echoes input", func(t *testing.T) {
		t.Parallel()
		server := app.CreateServer()
		b := browsertest.InitBrowser(t, server, e, WithMinLogLevel(slog.LevelWarn))
		win := b.OpenWindow("/codemirror/index.html")
		doc := win.HTMLDocument()
		editorArea := doc.QuerySelectorHTML(".CodeMirror textarea")
		editorArea.Focus()
		keyboard := controller.KeyboardController{Window: win}
		keyboard.SendText("Hello, world!")
		echo := doc.GetHTMLElementById("echo")
		assert.Equal(t, "Hello, world!", echo.TextContent())
	})
}
