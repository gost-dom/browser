package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/scripting/v8host"
	"github.com/onsi/gomega"
)

func TestHtmlAnchorElement(t *testing.T) {
	host := v8host.New()
	t.Cleanup(host.Close)
	win := html.NewWindow(html.WindowOptions{ScriptHost: host})
	defer win.Close()

	g := gomega.NewWithT(t)
	g.Expect(win.Eval(`
		const a = document.createElement("a");
		a.href = "http://example.com/foo";
		a.toString()
	`)).To(gomega.Equal("http://example.com/foo"))
}
