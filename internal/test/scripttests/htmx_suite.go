package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	app "github.com/gost-dom/browser/internal/test/integration/test-app"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
)

func RunHtmxTests(t *testing.T, e html.ScriptEngine) {
	t.Run("Click to increment counter", func(t *testing.T) {
		t.Parallel()

		expect := gomega.NewWithT(t).Expect
		server := app.CreateServer()
		b := htmltest.NewBrowserHelper(t, browsertest.InitBrowser(t, server, e))
		win, err := b.Open("/counter/index.html")
		expect(err).ToNot(HaveOccurred())
		counter := win.Document().GetElementById("counter").(html.HTMLElement)
		expect(counter).To(HaveInnerHTML(Equal("Count: 1")))
		counter.Click()
		counter = win.Document().GetElementById("counter").(html.HTMLElement)
		expect(counter).To(HaveInnerHTML(Equal("Count: 2")))
	})

	t.Run("Click hx-get link", func(t *testing.T) {
		t.Parallel()

		expect := gomega.NewWithT(t).Expect
		server := app.CreateServer()
		b := htmltest.NewBrowserHelper(t, browsertest.InitBrowser(t, server, e))
		win := b.OpenWindow("/navigation/page-a.html")
		expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue())
		expect(win.ScriptContext().Eval("window.pageB")).To(BeNil())

		// Click an hx-get link
		win.HTMLDocument().GetHTMLElementById("link-to-b").Click()

		expect(
			win.ScriptContext().Eval("window.pageA"),
		).To(BeTrue(), "Script context cleared from first page")
		expect(win.ScriptContext().Eval("window.pageB")).To(
			BeTrue(), "Scripts executed on second page")
		expect(win.Document()).To(
			HaveH1("Page B"), "Page heading", "Heading updated, i.e. htmx swapped")
		expect(win.Location().Pathname()).To(Equal("/navigation/page-a.html"), "Location updated")
	})

	t.Run("Clock boosted link", func(t *testing.T) {
		t.Parallel()

		expect := gomega.NewWithT(t).Expect
		server := app.CreateServer()
		b := htmltest.NewBrowserHelper(t, browsertest.InitBrowser(t, server, e))
		win, err := b.Open("/navigation/page-a.html")
		expect(err).ToNot(HaveOccurred())

		// Click an hx-boost link
		expect(win.ScriptContext().Eval("window.pageA")).ToNot(BeNil())
		expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue())
		expect(win.ScriptContext().Eval("window.pageB")).To(BeNil())
		win.Document().GetElementById("link-to-b-boosted").(html.HTMLElement).Click()

		expect(win.ScriptContext().Eval("window.pageA")).ToNot(BeNil(), "A")
		expect(win.ScriptContext().Eval("window.pageB")).ToNot(BeNil(), "B")
		expect(win.ScriptContext().Eval("window.pageA")).To(
			BeTrue(), "Script context cleared from first page")
		expect(win.ScriptContext().Eval("window.pageB")).To(
			BeTrue(), "Scripts executed on second page")
		expect(win.Document()).To(
			HaveH1("Page B"), "Page heading", "Heading updated, i.e. htmx swapped")
		expect(win.Location().Pathname()).To(Equal("/navigation/page-b.html"), "Location updated")
	})

	t.Run("Form submit", func(t *testing.T) {
		t.Parallel()

		expect := gomega.NewWithT(t).Expect
		server := app.CreateServer()
		b := htmltest.NewBrowserHelper(t, browsertest.InitBrowser(t, server, e))
		win, err := b.Open("/forms/form-1.html")
		expect(err).ToNot(HaveOccurred())
		i1 := win.Document().GetElementById("field-1")
		i1.SetAttribute("value", "Foo")

		btn := win.Document().GetElementById("submit-btn").(html.HTMLElement)
		expect(len(server.Requests)).To(Equal(2), "No of requests _before_ click")
		btn.Click()
		expect(len(server.Requests)).To(Equal(3), "No of requests _after_ click")
		elm := win.Document().GetElementById("field-value-1")
		expect(elm).To(HaveTextContent("Foo"))
	})
}
