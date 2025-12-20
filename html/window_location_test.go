package html_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

type WindowLocationTestSuite struct {
	gosttest.GomegaSuite
	window html.Window
}

func (s *WindowLocationTestSuite) SetupTest() {
	server := newAnchorTagNavigationServer()
	b := browsertest.InitBrowser(s.T(), server, nil)
	s.window = b.NewWindow()
}

func TestWindowLocation(t *testing.T) {
	suite.Run(t, new(WindowLocationTestSuite))
}

func (s *WindowLocationTestSuite) TestEmptyWindow() {
	s.Expect(s.window.Location().Href()).To(Equal("about:blank"))
}

func (s *WindowLocationTestSuite) TestPathname() {
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
	s.window.DispatchEvent(&event.Event{Type: "gost-event"})
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

func (s *WindowLocationTestSuite) TestParseLocationWithoutQuery() {
	win := s.openWindow("http://localhost:9999/foo/bar")
	location := win.Location()
	s.Expect(location.Host()).To(Equal("localhost:9999"), "host")
	s.Expect(location.Hash()).To(Equal(""), "hash")
	s.Expect(location.Hostname()).To(Equal("localhost"), "hostname")
	s.Expect(location.Href()).To(Equal("http://localhost:9999/foo/bar"), "href")
	s.Expect(location.Origin()).To(Equal("http://localhost:9999"), "origin")
	s.Expect(location.Pathname()).To(Equal("/foo/bar"), "Pathname")
	s.Expect(location.Port()).To(Equal("9999"), "port")
	s.Expect(location.Protocol()).To(Equal("http:"), "protocol")
	s.Expect(location.Search()).To(Equal(""), "query")
}

func (s *WindowLocationTestSuite) TestParseLocationWithQuery() {
	win := s.openWindow("http://localhost:9999/foo/bar?q=baz")
	location := win.Location()
	s.Expect(location.Host()).To(Equal("localhost:9999"), "host")
	s.Expect(location.Hostname()).To(Equal("localhost"), "hostname")
	s.Expect(location.Href()).To(Equal("http://localhost:9999/foo/bar?q=baz"), "href")
	s.Expect(location.Origin()).To(Equal("http://localhost:9999"), "origin")
	s.Expect(location.Pathname()).To(Equal("/foo/bar"), "Pathname")
	s.Expect(location.Port()).To(Equal("9999"), "port")
	s.Expect(location.Protocol()).To(Equal("http:"), "protocol")
	s.Expect(location.Search()).To(Equal("?q=baz"), "query")

}

func (s *WindowLocationTestSuite) TestParseLocationWithoutFragment() {
	win := s.openWindow("http://localhost:9999/foo#heading-1")
	location := win.Location()
	s.Expect(location.Host()).To(Equal("localhost:9999"), "host")
	s.Expect(location.Hash()).To(Equal("#heading-1"), "hash")
	s.Expect(location.Hostname()).To(Equal("localhost"), "hostname")
	s.Expect(location.Href()).To(Equal("http://localhost:9999/foo#heading-1"), "href")
}

func (s *WindowLocationTestSuite) openWindow(location string) html.Window {
	handler := http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) { res.Write([]byte("<html></html>")) },
	)
	windowOptions := html.WindowOptions{
		HttpClient: gosthttp.NewHttpClientFromHandler(handler),
	}
	win, err := html.OpenWindowFromLocation(location, windowOptions)
	assert.NoError(s.T(), err)
	s.T().Cleanup(func() {
		if win != nil {
			win.Close()
		}
	})
	return win
}

func newAnchorTagNavigationServer() http.Handler {
	server := http.NewServeMux()
	server.HandleFunc("GET /index",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(
				`<body>
					<h1>Index</h1>
					<a href="products">Products from relative url</a>
					<a href="/products">Products from absolute url</a>
				</body>`))
		})

	server.HandleFunc("GET /products",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(
				`<body>
					<h1>Products</h1>
				</body>`))
		})

	return server
}
