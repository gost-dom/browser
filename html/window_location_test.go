package html_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/gost-dom/browser/testing/testservers"
)

type WindowLocationTestSuite struct {
	gosttest.GomegaSuite
	window html.Window
}

func (s *WindowLocationTestSuite) SetupTest() {
	server := testservers.NewAnchorTagNavigationServer()
	s.window = NewWindowFromHandler(server)

}

func TestWindowLocation(t *testing.T) {
	suite.Run(t, new(WindowLocationTestSuite))
}

func (s *WindowLocationTestSuite) TestEmptyWindow() {
	s.Expect(s.window.Location().Href()).To(Equal("about:blank"))
}

func (s *WindowLocationTestSuite) TestPathnam() {
	s.Expect(s.window.Navigate("/index")).To(Succeed())
	s.Expect(s.window.Location().Pathname()).To(Equal("/index"))
}

func (s *WindowLocationTestSuite) TestNavigateToAboutBlank() {
	s.Expect(s.window.Navigate("about:blank")).To(Succeed())
	s.Expect(s.window.Document()).To(HaveH1("Gost-DOM"))
}

func (s *WindowLocationTestSuite) TestNavigateClearsEventHandlers() {
	count := 0
	s.Expect(s.window.Navigate("about:blank")).To(Succeed())
	s.window.AddEventListener(
		"gost-event",
		event.NewEventHandlerFunc(func(e *event.Event) error {
			count++
			return nil
		}))

	s.Expect(s.window.Navigate("/index")).To(Succeed())
	s.window.DispatchEvent(event.New("gost-event", nil))
	s.Expect(count).To(Equal(0))
}

func (s *WindowLocationTestSuite) GetLink(text string) html.HTMLElement {
	s.T().Helper()
	s.Expect(s.window.Navigate("/index")).To(Succeed())
	nodes, err := s.window.Document().QuerySelectorAll("a")
	s.Expect(err).ToNot(HaveOccurred())
	for _, n := range nodes.All() {
		if n.TextContent() == text {
			return n.(html.HTMLElement)
		}
	}
	s.T().Fatalf("Link not found with text: %s", text)
	return nil
}

func (s *WindowLocationTestSuite) TestClickAbsoluteURL() {
	link := s.GetLink("Products from absolute url")
	link.Click()
	s.Expect(s.window.Location().Pathname()).To(Equal("/products"))

}

func (s *WindowLocationTestSuite) TestClickRelativeURL() {
	link := s.GetLink("Products from relative url")
	link.Click()
	s.Expect(s.window.Location().Pathname()).To(Equal("/products"))

}
