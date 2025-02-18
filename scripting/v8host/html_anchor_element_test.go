package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/onsi/gomega"
)

func TestHtmlAnchorElement(t *testing.T) {
	win := html.NewWindow(html.WindowOptions{ScriptHost: host})
	defer win.Close()

	g := gomega.NewWithT(t)
	g.Expect(win.Eval(`
		const a = document.createElement("a");
		a.href = "http://example.com/foo";
		a.toString()
	`)).To(gomega.Equal("http://example.com/foo"))
}
