package browser_test

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"testing"
	"testing/synctest"

	"github.com/gost-dom/browser"
	"github.com/gost-dom/browser/dom/event"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/gost-dom/browser/v8browser"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BrowserTestSuite struct {
	suite.Suite
}

func (s *BrowserTestSuite) TestReadFromHTTPHandler() {
	browser := v8browser.New(browser.WithHandler(gosttest.StaticHTML("<html></html>")))
	defer browser.Close()

	result, err := browser.Open("/")
	s.Assert().NoError(err)

	element := result.Document().DocumentElement()
	s.Assert().Equal("HTML", element.NodeName())
	s.Assert().Equal("HTML", element.TagName())
}

func TestBrowserClosedOnCancel(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(t.Context())
		browser := v8browser.New(browser.WithContext(ctx))
		cancel()
		synctest.Wait()
		assert.True(t, browser.Closed())
	})
}

func (s *BrowserTestSuite) TestExecuteScript() {
	Expect := gomega.NewWithT(s.T()).Expect
	server := gosttest.HttpHandlerMap{
		"/index.html": gosttest.StaticHTML(
			`<body>
				<div id='target'></div>
				<script>
					const target = document.getElementById('target');
					target.textContent = "42"
				</script>
			</body>`),
	}
	browser := v8browser.New(browser.WithHandler(server))
	s.T().Cleanup(browser.Close)

	win, err := browser.Open("/index.html")
	s.Assert().NoError(err)
	Expect(err).ToNot(HaveOccurred())
	target := win.Document().GetElementById("target")
	Expect(target).To(HaveOuterHTML(`<div id="target">42</div>`))
}

func (s *BrowserTestSuite) TestCancellation() {
	synctest.Test(s.T(), func(t *testing.T) {
		handler := gosttest.NewPipeHandler(t)
		defer handler.Close()
		h := gosttest.HttpHandlerMap{
			"/index.html": gosttest.StaticHTML("body>Dummy</body>"),
			"/data":       handler,
		}

		ctx, cancel := context.WithCancel(t.Context())
		b := v8browser.New(
			browser.WithHandler(h),
			browser.WithContext(ctx),
		)
		w, err := b.Open("http://example.com/index.html")
		if assert.NoError(t, err) {
			assert.NoError(t, w.Run("fetch('/data')"))
			synctest.Wait()
			assert.NotNil(t, handler.Req, "/data requested")
			cancel()
			synctest.Wait()
			assert.True(t, handler.ClientDisconnected, "http client disconnected")
		} else {
			cancel()
		}
	})
}

func TestBrowserSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(BrowserTestSuite))
}

type BrowserNavigationTestSuite struct {
	suite.Suite
	gomega.Gomega
}

func (s *BrowserNavigationTestSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

func (s *BrowserNavigationTestSuite) loadPageA() htmltest.WindowHelper {
	server := newBrowserNavigateTestServer()
	browser := htmltest.NewBrowserHelper(s.T(), v8browser.New(browser.WithHandler(server)))
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

	s.Expect(window.ScriptContext().Eval("typeof loadedA")).
		To(Equal("undefined"), "Global state cleared after clicking a link to navigate")
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
	browser := v8browser.New(browser.WithHandler(http.HandlerFunc(cookieHandler)))
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
	b := v8browser.New(browser.WithHandler(http.HandlerFunc(cookieHandler)))
	win, err := b.Open("http://localhost/")
	Expect(err).ToNot(HaveOccurred())
	el := win.Document().GetElementById("gost")
	Expect(el).To(HaveTextContent(""))

	b = v8browser.New(browser.WithHandler(http.HandlerFunc(cookieHandler)))
	win, err = b.Open("http://localhost/")
	Expect(err).ToNot(HaveOccurred())
	el = win.Document().GetElementById("gost")
	Expect(el).To(HaveTextContent(""))
}

func TestCookies(t *testing.T) {
	suite.Run(t, new(CookiesTestSuite))
}

func TestLogOutput(t *testing.T) {
	var buf bytes.Buffer
	Expect := gomega.NewWithT(t).Expect
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	b := v8browser.New(
		browser.WithHandler(http.HandlerFunc(cookieHandler)),
		browser.WithLogger(logger),
	)
	defer b.Close()
	win, err := b.Open("http://localhost/")
	Expect(err).ToNot(HaveOccurred())
	win.Run("console.log('foo bar')")
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

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	var gost string
	if c, _ := r.Cookie("gost"); c != nil {
		gost = c.Value
	}
	w.Header().Add("Set-Cookie", "gost=Hello, World!")
	fmt.Fprintf(w, `<body><div id="gost">%s</div></body>`, gost)
}

func newBrowserNavigateTestServer() http.Handler {
	return gosttest.HttpHandlerMap{
		"/a.html": gosttest.StaticHTML(
			`<body>
				<h1>Page A</h1>
				<a href="b.html">Load B</a>
				<script>loadedA = "PAGE A"</script>
			</body>`),
		"/b.html": gosttest.StaticHTML(
			`<body>
				<h1>Page B</h1>
				<script>loadedB = "PAGE B"</script>
			</body>`),
	}
}

func TestBrowserOpenRedirect(t *testing.T) {
	b := v8browser.New(browser.WithHandler(
		gosttest.HttpHandlerMap{
			"/old-location": http.RedirectHandler("/new-location", 301),
			"/new-location": gosttest.StaticHTML("<body><h1>Hello</h1></body>"),
		},
	))
	win, err := b.Open("/old-location")
	assert.NoError(t, err)
	heading, _ := win.Document().QuerySelector("h1")
	assert.Equal(t, "Hello", heading.TextContent())
	assert.Equal(t, "/new-location", win.Location().Pathname())
}
