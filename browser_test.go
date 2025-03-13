package browser_test

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom/event"
	. "github.com/gost-dom/browser/internal/testing"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	. "github.com/gost-dom/browser/testing/gomega-matchers"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type BrowserTestSuite struct {
	suite.Suite
}

func (s *BrowserTestSuite) TestReadFromHTTPHandler() {
	Expect := gomega.NewWithT(s.T()).Expect
	handler := (http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "text/html") // For good measure, not used yet"
		w.Write([]byte("<html></html>"))
	})
	browser := NewBrowserFromHandler(handler)
	result, err := browser.Open("/")
	Expect(err).ToNot(HaveOccurred())
	element := result.Document().DocumentElement()

	Expect(element.NodeName()).To(Equal("HTML"))
	Expect(element.TagName()).To(Equal("HTML"))
}

func (s *BrowserTestSuite) TestExecuteScript() {
	Expect := gomega.NewWithT(s.T()).Expect
	// This is not necessarily desired behaviour right now.
	server := http.NewServeMux()
	server.Handle(
		"GET /index.html",
		http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(`<body>
					<div id='target'></div>
					<script>
						const target = document.getElementById('target');
						target.textContent = "42"
					</script>
				</body>`))
		}),
	)
	browser := NewBrowserFromHandler(server)
	s.T().Cleanup(browser.Close)

	win, err := browser.Open("/index.html")
	Expect(err).ToNot(HaveOccurred())
	target := win.Document().GetElementById("target")
	Expect(target).To(HaveOuterHTML(Equal(`<div id="target">42</div>`)))
}

func TestBrowserSuite(t *testing.T) {
	suite.Run(t, new(BrowserTestSuite))
}

type BrowserNavigationTestSuite struct {
	suite.Suite
	gomega.Gomega
}

func (s *BrowserNavigationTestSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

func (s *BrowserNavigationTestSuite) loadPageA() WindowHelper {
	server := newBrowserNavigateTestServer()
	browser := NewBrowserHelper(s.T(), NewBrowserFromHandler(server))
	window := browser.OpenWindow("/a.html")
	return window
}

func (s *BrowserNavigationTestSuite) TestPageAHasLoaded() {
	Expect := gomega.NewWithT(s.T()).Expect
	window := s.loadPageA()
	heading, _ := window.Document().QuerySelector("h1")
	Expect(heading).To(HaveTextContent(Equal("Page A")))
	Expect(window.ScriptContext().Eval("loadedA")).To(Equal("PAGE A"))
}

func (s *BrowserNavigationTestSuite) TestClickLink() {
	window := s.loadPageA()
	window.HTMLDocument().QuerySelectorHTML("a").Click()

	heading, _ := window.Document().QuerySelector("h1")
	s.Expect(heading).To(HaveTextContent(Equal("Page B")))
	s.Expect(window.ScriptContext().Eval("loadedB")).To(Equal("PAGE B"))

	// The global state should have been cleared
	s.Expect(window.ScriptContext().Eval("typeof loadedA")).To(Equal("undefined"))
}

func (s *BrowserNavigationTestSuite) TestNavigationAbortedByEventHandler() {
	window := s.loadPageA()
	anchor := window.HTMLDocument().QuerySelectorHTML("a")
	anchor.AddEventListener(
		"click",
		event.NewEventHandlerFunc(event.NoError((*event.Event).PreventDefault)),
	)
	anchor.Click()
	heading, _ := window.Document().QuerySelector("h1")
	s.Expect(heading).To(HaveTextContent(Equal("Page A")))
}

func TestBrowserNavigation(t *testing.T) {
	suite.Run(t, new(BrowserNavigationTestSuite))
}

type CookiesTestSuite struct {
	suite.Suite
}

func (s *CookiesTestSuite) TestCookiesArePersistedInSameBrowser() {
	Expect := gomega.NewWithT(s.T()).Expect
	browser := NewBrowserFromHandler(http.HandlerFunc(cookieHandler))
	win, err := browser.Open("http://localhost/")
	Expect(err).ToNot(HaveOccurred())
	el := win.Document().GetElementById("gost")
	Expect(el).To(HaveTextContent(""))

	Expect(win.Navigate("http://localhost/")).To(Succeed())
	el = win.Document().GetElementById("gost")
	Expect(el).To(HaveTextContent("Hello, World!"))
}

func (s *CookiesTestSuite) TestCookiesAreNotReusedInNewBrowser() {
	Expect := gomega.NewWithT(s.T()).Expect
	browser := NewBrowserFromHandler(http.HandlerFunc(cookieHandler))
	win, err := browser.Open("http://localhost/")
	Expect(err).ToNot(HaveOccurred())
	el := win.Document().GetElementById("gost")
	Expect(el).To(HaveTextContent(""))

	browser = NewBrowserFromHandler(http.HandlerFunc(cookieHandler))
	win, err = browser.Open("http://localhost/")
	Expect(err).ToNot(HaveOccurred())
	el = win.Document().GetElementById("gost")
	Expect(el).To(HaveTextContent(""))

}

func TestCookies(t *testing.T) {
	suite.Run(t, new(CookiesTestSuite))
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	var gost string
	if c, _ := r.Cookie("gost"); c != nil {
		gost = c.Value
	}
	w.Header().Add("Set-Cookie", "gost=Hello, World!")
	w.Write([]byte(fmt.Sprintf(`<body><div id="gost">%s</div></body>`, gost)))
}

func newBrowserNavigateTestServer() http.Handler {
	server := http.NewServeMux()
	server.HandleFunc("GET /a.html",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(
				`<body>
					<h1>Page A</h1>
					<a href="b.html">Load B</a>
					<script>loadedA = "PAGE A"</script>
				</body>`))
		})

	server.HandleFunc("GET /b.html",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(`
				<body>
					<h1>Page B</h1>
					<script>loadedB = "PAGE B"</script>
				</body>`))
		})

	return server
}
