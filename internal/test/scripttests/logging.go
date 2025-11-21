package scripttests

import (
	"bytes"
	"fmt"
	"log/slog"
	"testing"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func testLogging(t *testing.T, e html.ScriptEngine) {
	var buf bytes.Buffer
	Expect := gomega.NewWithT(t).Expect
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	logger.Info(fmt.Sprintf("Logger: %v", logger))
	b := browser.New(
		browser.WithScriptEngine(e),
		browser.WithLogger(logger),
	)
	defer b.Close()
	win := b.NewWindow()
	Expect(win.Run("console.log('foo bar')")).To(Succeed())
	Expect(buf.String()).To(ContainSubstring("foo bar"))

	buf.Reset()
	win.DispatchEvent(event.NewCustomEvent("dummy", event.CustomEventInit{}))
	Expect(buf.String()).To(ContainSubstring(`msg="Dispatch event"`))

	buf.Reset()
	win.Document().Body().AppendChild(win.Document().CreateElement("div"))
	win.Document().Body().DispatchEvent(event.NewCustomEvent("dummy", event.CustomEventInit{}))
	Expect(buf.String()).To(ContainSubstring(`msg=Node.AppendChild`))
	Expect(buf.String()).To(ContainSubstring(`msg="Dispatch event"`))
	buf.Reset()

}
