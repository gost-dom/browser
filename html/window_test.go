package html_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
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

type WindowNavigationTestSuite struct {
	gosttest.GomegaSuite
	win htmltest.WindowHelper
}

func (s *WindowNavigationTestSuite) SetupTest() {
	m := http.NewServeMux()
	m.HandleFunc("GET /page-1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<body>
			<a id="link" href="/old-page-2">Link</a>
			<form method="post" action="/old-page-2" id="form">
				<input name="data" value="value" type="text" />
				<button type="submit" id="submit">Submit</button>
			</form>
			</body>`))
	})
	m.HandleFunc("/old-page-2", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-page-2", http.StatusSeeOther)
	})
	m.HandleFunc("/new-page-2", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
		}
		respBody := fmt.Sprintf(
			`<body><h1>Page 2</h1>
				<div id="method">%s</div>
				<div id="request-body">%s</div>
			</body>`, r.Method, string(body),
		)
		w.Write([]byte(respBody))
	})
	m.HandleFunc("/infinite-redirects", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/infinite-redirects", 301)
	})

	b := browsertest.InitBrowser(s.T(), m, nil)
	s.win = b.OpenWindow("https://example.com/page-1")
}

func TestWindowNavigation(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WindowNavigationTestSuite))
}

func (s *WindowNavigationTestSuite) TestRedirectGetRequests() {
	s.Assert().Equal("/page-1", s.win.Location().Pathname())
	s.win.HTMLDocument().GetHTMLElementById("link").Click()
	s.Assert().Equal("https://example.com/new-page-2", s.win.Location().Href())
}

func (s *WindowNavigationTestSuite) TestRedirectNavigate() {
	s.win.Navigate("https://example.com/old-page-2")
	s.Assert().Equal("/new-page-2", s.win.Location().Pathname())
}

func (s *WindowNavigationTestSuite) TestInfinteRedirects() {
	err := s.win.Navigate("/infinite-redirects")
	s.Assert().ErrorIs(err, html.ErrTooManyRedirects, "Error is too many redirects")
}

func (s *WindowNavigationTestSuite) TestNavigateHref() {
	s.win.Navigate("https://example.com/page-1")
	s.win.Navigate("/new-page-2")
	s.Assert().Equal(
		"https://example.com/new-page-2", s.win.LocationHREF(),
		"Navigate resolves the HREF for a local URI",
	)
}
