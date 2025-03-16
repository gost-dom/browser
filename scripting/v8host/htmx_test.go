package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/html"
	app "github.com/gost-dom/browser/internal/test/integration/htmx-app"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

var _ = Describe("HTMX Tests", Ordered, func() {
	var b htmltest.BrowserHelper
	var server *app.TestServer

	BeforeEach(func() {
		server = app.CreateServer()
		b = htmltest.NewBrowserHelper(GinkgoTB(), browser.NewBrowserFromHandler(server))
		DeferCleanup(func() {
			b.Close()
		})
	})

	It("Should increment the counter example", func() {
		win, err := b.Open("/counter/index.html")
		Expect(err).ToNot(HaveOccurred())
		counter := win.Document().GetElementById("counter").(html.HTMLElement)
		Expect(counter).To(HaveInnerHTML(Equal("Count: 1")))
		counter.Click()
		counter = win.Document().GetElementById("counter").(html.HTMLElement)
		Expect(counter).To(HaveInnerHTML(Equal("Count: 2")))
	})

	It("Should not update the location when a link has hx-get", func() {
		Skip("Need sync")
		win := b.OpenWindow("/navigation/page-a.html")
		Expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue())
		Expect(win.ScriptContext().Eval("window.pageB")).To(BeNil())

		// Click an hx-get link
		win.HTMLDocument().GetHTMLElementById("link-to-b").Click()

		Expect(
			win.ScriptContext().Eval("window.pageA"),
		).To(BeTrue(), "Script context cleared from first page")
		Expect(win.ScriptContext().Eval("window.pageB")).To(
			BeTrue(), "Scripts executed on second page")
		Expect(win.Document()).To(
			HaveH1("Page B"), "Page heading", "Heading updated, i.e. htmx swapped")
		Expect(win.Location().Pathname()).To(Equal("/navigation/page-a.html"), "Location updated")
	})

	It("Should update the location when a link with href is boosted", func() {
		win, err := b.Open("/navigation/page-a.html")
		Expect(err).ToNot(HaveOccurred())

		// Click an hx-boost link
		Expect(win.ScriptContext().Eval("window.pageA")).ToNot(BeNil())
		Expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue())
		Expect(win.ScriptContext().Eval("window.pageB")).To(BeNil())
		win.Document().GetElementById("link-to-b-boosted").(html.HTMLElement).Click()

		Expect(win.ScriptContext().Eval("window.pageA")).ToNot(BeNil(), "A")
		Expect(win.ScriptContext().Eval("window.pageB")).ToNot(BeNil(), "B")
		Expect(win.ScriptContext().Eval("window.pageA")).To(
			BeTrue(), "Script context cleared from first page")
		Expect(win.ScriptContext().Eval("window.pageB")).To(
			BeTrue(), "Scripts executed on second page")
		Expect(win.Document()).To(
			HaveH1("Page B"), "Page heading", "Heading updated, i.e. htmx swapped")
		Expect(win.Location().Pathname()).To(Equal("/navigation/page-b.html"), "Location updated")
	})

	It("Should submit forms", func() {
		Skip("Need to update")
		win, err := b.Open("/forms/form-1.html")
		Expect(err).ToNot(HaveOccurred())
		i1 := win.Document().GetElementById("field-1")
		i1.SetAttribute("value", "Foo")
		b := win.Document().GetElementById("submit-btn").(html.HTMLElement)
		Expect(len(server.Requests)).To(Equal(2), "No of requests _before_ click")
		b.Click()
		Expect(len(server.Requests)).To(Equal(3), "No of requests _after_ click")
		elm := win.Document().GetElementById("field-value-1")
		Expect(elm).To(HaveTextContent("Foo"))
	})
})
