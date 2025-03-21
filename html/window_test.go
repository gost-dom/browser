package html_test

import (
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/domslices"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/gost-dom/browser/testing/testservers"
)

type WindowTestSuite struct {
	gosttest.GomegaSuite
}

func (s *WindowTestSuite) TestDocumentIsAnHTMLDocument() {
	win, err := NewWindowReader(strings.NewReader("<html><body></body></html>"))
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(win.Document().DocumentElement()).To(BeHTMLElement())
}

func (s *WindowTestSuite) TestDocumentWithDOCTYPE() {
	win, err := NewWindowReader(strings.NewReader("<!DOCTYPE HTML><html><body></body></html>"))
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(win.Document().FirstChild().NodeType()).To(Equal(dom.NodeTypeDocumentType))
}

func TestWindow(t *testing.T) {
	suite.Run(t, new(WindowTestSuite))
}

var _ = Describe("Window", func() {
	Describe("Location()", func() {
		var window Window

		BeforeEach(func() {
			server := testservers.NewAnchorTagNavigationServer()
			DeferCleanup(func() { server = nil })
			window = NewWindowFromHandler(server)
		})

		It("Should be about:blank", func() {
			Expect(window.Location().Href()).To(Equal("about:blank"))
		})

		It("Should return the path loaded from", func() {
			Expect(window.Navigate("/index")).To(Succeed())
			Expect(window.Location().Pathname()).To(Equal("/index"))
		})

		Describe("Navigate", func() {
			It("Should load a blank page when loading about:blank", func() {
				Expect(window.Navigate("about:blank")).To(Succeed())
				Expect(window.Document()).To(HaveH1("Gost-DOM"))
			})

			It("Should clear event handlers", func() {
				count := 0
				Expect(window.Navigate("about:blank")).To(Succeed())
				window.AddEventListener(
					"gost-event",
					event.NewEventHandlerFunc(func(e *event.Event) error {
						count++
						return nil
					}))

				Expect(window.Navigate("/index")).To(Succeed())
				window.DispatchEvent(event.New("gost-event", nil))
				Expect(count).To(Equal(0))
			})
		})

		Describe("User navigation (clicking links)", func() {
			var links []dom.Node

			BeforeEach(func() {
				Expect(window.Navigate("/index")).To(Succeed())
				nodes, err := window.Document().QuerySelectorAll("a")
				Expect(err).ToNot(HaveOccurred())
				links = nodes.All()
			})

			It("Should update when using a link with absolute url", func() {
				link, ok := domslices.SliceFindFunc(links, func(n dom.Node) bool {
					return n.TextContent() == "Products from absolute url"
				})
				Expect(ok).To(BeTrue())
				link.(html.HTMLElement).Click()
				Expect(window.Location().Pathname()).To(Equal("/products"))
			})

			It("Should update when using a link with relative url", func() {
				link, ok := domslices.SliceFindFunc(links, func(n dom.Node) bool {
					return n.TextContent() == "Products from relative url"
				})
				Expect(ok).To(BeTrue())
				link.(html.HTMLElement).Click()
				Expect(window.Location().Pathname()).To(Equal("/products"))
			})
		})
	})
})
