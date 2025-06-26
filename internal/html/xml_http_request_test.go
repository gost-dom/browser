package html_test

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"testing"

	"github.com/gost-dom/browser/dom/event"
	. "github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/clock"
	"github.com/gost-dom/browser/internal/gosthttp"
	. "github.com/gost-dom/browser/internal/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/onsi/gomega/types"
)

type stubBrowsingContext struct {
	client http.Client
	url    string
}

func (c stubBrowsingContext) HTTPClient() http.Client  { return c.client }
func (c stubBrowsingContext) LocationHREF() string     { return c.url }
func (c stubBrowsingContext) Logger() *slog.Logger     { return nil }
func (c stubBrowsingContext) Context() context.Context { return nil }

type XMLHTTPRequestTestSuite struct {
	gosttest.GomegaSuite
	handler        http.Handler
	actualHeader   http.Header
	actualMethod   string
	actualReqBody  []byte
	reqErr         error
	responseHeader http.Header
	xhr            XmlHttpRequest
	timer          *clock.Clock
}

func TestXMLHTTPRequest(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(XMLHTTPRequestTestSuite))
}

func (s *XMLHTTPRequestTestSuite) SetupTest() {
	// Create a basic server for testing
	s.timer = clock.New()
	s.handler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		s.actualHeader = req.Header
		s.actualMethod = req.Method
		if req.Body != nil {
			s.actualReqBody, s.reqErr = io.ReadAll(req.Body)
		}
		for k, vs := range s.responseHeader {
			for _, v := range vs {
				w.Header().Add(k, v)
			}
		}
		if s.responseHeader == nil || s.responseHeader.Get("Content-Type") == "" {
			w.Header().Add("Content-Type", "text/plain")
		}
		w.Write([]byte("Hello, World!"))
	})
	s.xhr = NewXmlHttpRequest(
		stubBrowsingContext{client: gosthttp.NewHttpClientFromHandler(s.handler)},
		s.timer,
	)

	s.T().Cleanup(func() {
		// Allow GC after test run
		s.handler = nil
		s.actualHeader = nil
	})

}

func (s *XMLHTTPRequestTestSuite) TestSynchronousRequest() {
	s.xhr.Open("GET", "/dummy", RequestOptionAsync(false))
	s.Expect(s.xhr.Status()).To(Equal(0))
	s.Expect(s.xhr.Send(nil)).To(Succeed())
	// Verify request
	s.Expect(s.actualMethod).To(Equal("GET"))
	// Verify response
	s.Expect(s.xhr.Status()).To(Equal(200))
	// This is the only place we test StatusText; it's dumb wrapper and may
	// be removed.
	s.Expect(s.xhr.StatusText()).To(Equal("OK"))
	s.Expect(s.xhr.ResponseText()).To(Equal("Hello, World!"))
}

func (s *XMLHTTPRequestTestSuite) TestAsynchronousRequest() {
	var (
		loadStarted bool
		loadEnded   bool
		loaded      bool
	)
	s.xhr.Open("GET", "/dummy")
	s.xhr.AddEventListener(
		XHREventLoadstart,
		event.NewEventHandlerFunc(func(e *event.Event) error {
			loadStarted = true
			return nil
		}),
	)
	s.xhr.AddEventListener(
		XHREventLoadend,
		event.NewEventHandlerFunc(func(e *event.Event) error {
			loadEnded = true
			return nil
		}),
	)
	s.xhr.AddEventListener(
		XHREventLoad,
		event.NewEventHandlerFunc(func(e *event.Event) error {
			loaded = true
			return nil
		}),
	)
	s.Expect(s.xhr.Send(nil)).To(Succeed())
	s.Expect(loadStarted).To(BeTrue(), "loadstart emitted")
	s.Expect(s.xhr.Status()).To(Equal(0), "Response should not have been received yet")
	s.Expect(loadEnded).To(BeFalse(), "loadend emitted")
	s.Expect(loaded).To(BeFalse(), "load emitted")

	s.Expect(s.timer.RunAll()).To(Succeed())

	s.Expect(s.xhr.Status()).To(Equal(200))
	s.Expect(s.xhr.ResponseText()).To(Equal("Hello, World!"))
}

func (s *XMLHTTPRequestTestSuite) TestFormdataEncoding() {
	// This test uses blocking requests.
	// This isn't the ususal case, but the test is much easier to write; and
	// code being tested is unrelated to blocking/non-blocking.
	s.xhr.Open("POST", "/dummy", RequestOptionAsync(false))
	formData := NewFormData()
	formData.Append("key1", "Value%42")
	formData.Append("key2", "Value&=42")
	formData.Append("key3", "International? æøå")
	s.xhr.SendBody(formData.GetReader())
	s.Expect(s.reqErr).ToNot(HaveOccurred())
	s.Expect(s.actualMethod).To(Equal("POST"))
	actualReqContentType := s.actualHeader.Get("Content-Type")
	s.Expect(actualReqContentType).To(Equal("application/x-www-form-urlencoded"))
	s.Expect(
		string(s.actualReqBody),
	).To(Equal("key1=Value%2542&key2=Value%26%3D42&key3=International%3F+%C3%A6%C3%B8%C3%A5"))
}

func (s *XMLHTTPRequestTestSuite) TestRequestHeaders() {
	s.xhr.SetRequestHeader("x-test", "42")
	s.xhr.Open("GET", "/dummy", RequestOptionAsync(false))
	s.Expect(s.xhr.Send(nil)).To(Succeed())
	s.Expect(s.actualHeader.Get("x-test")).To(Equal("42"))
}

func (s *XMLHTTPRequestTestSuite) TestResponseHeaders() {
	s.responseHeader = make(http.Header)
	s.responseHeader.Add("X-Test-1", "value1")
	s.responseHeader.Add("X-Test-2", "value2")
	s.responseHeader.Add("Content-Type", "text/plain")
	s.xhr.Open("GET", "/dummy", RequestOptionAsync(false))
	s.xhr.Send(nil)
	s.Expect(
		s.xhr.GetAllResponseHeaders(),
	).To(HaveLines("x-test-1: value1", "x-test-2: value2", "content-type: text/plain"))

	_, found := s.xhr.GetResponseHeader("missing")
	s.Expect(found).To(BeFalse(), "Value of non-existing response header")
}

func (s *XMLHTTPRequestTestSuite) TestSameResponseHeaderAddedTwice() {
	s.responseHeader = make(http.Header)
	s.responseHeader.Add("X-Test-1", "value1")
	s.responseHeader.Add("X-Test-2", "value2")
	s.responseHeader.Add("Content-Type", "text/plain")
	s.responseHeader.Add("x-test-1", "value3")

	s.xhr.Open("GET", "/dummy", RequestOptionAsync(false))
	s.xhr.Send(nil)

	s.Expect(
		s.xhr.GetAllResponseHeaders(),
	).To(HaveLines("x-test-1: value1", "x-test-1: value3", "x-test-2: value2", "content-type: text/plain"))

	header1, found1 := s.xhr.GetResponseHeader("x-test-1")
	header2, found2 := s.xhr.GetResponseHeader("x-test-2")

	s.Expect(header2).To(Equal("value2"), "Value of header specified once")
	s.Expect(header1).To(Equal("value1, value3"), "Value of header specified twice")
	assert.True(s.T(), found1)
	assert.True(s.T(), found2)
}

func (s *XMLHTTPRequestTestSuite) TestCookieVisibility() {
	s.responseHeader = make(http.Header)
	s.responseHeader.Add("X-Test-1", "value1")
	s.responseHeader.Add("X-Test-2", "value2")
	s.responseHeader.Add("Content-Type", "text/plain")

	s.xhr.Open("GET", "/dummy", RequestOptionAsync(false))
	s.xhr.Send(nil)

	s.responseHeader = make(http.Header)
	s.responseHeader.Add("set-cookie", "foobar-should-not-be-visible")
	s.Expect(
		s.xhr.GetAllResponseHeaders(),
	).To(HaveLines("x-test-1: value1", "x-test-2: value2", "content-type: text/plain"))
}

func TestXMLHTTPRequestRedirect(t *testing.T) {
	m := http.NewServeMux()
	m.Handle("GET /redirect", http.RedirectHandler("/redirect-temp", 301))
	m.Handle("GET /redirect-temp", http.RedirectHandler("/redirected-url", 301))
	m.HandleFunc("GET /redirected-url", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Handled"))
	})
	xhr := NewXmlHttpRequest(
		stubBrowsingContext{client: gosthttp.NewHttpClientFromHandler(m)},
		clock.New(),
	)
	xhr.Open("GET", "https://example.com/redirect", RequestOptionAsync(false))
	xhr.Send(nil)

	assert.Equal(t, "https://example.com/redirected-url", xhr.ResponseURL())
}

func HaveLines(expected ...string) types.GomegaMatcher {
	return WithTransform(func(s string) []string {
		lines := strings.Split(s, "\r\n")
		return slices.DeleteFunc(lines, func(line string) bool { return line == "" })
	}, ConsistOf(expected))
}
