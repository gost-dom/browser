package html_test

import (
	"net/http"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/suite"
)

type HTMLFormElementWithCheckboxTestSuite struct {
	suite.Suite
}

func TestHTMLFormElementWithCheckbox(t *testing.T) {
	suite.Run(t, new(HTMLFormElementWithCheckboxTestSuite))
}

// httpRequestRecorder is a very simple http.Handler that just records the
// incoming requests, and returns a 200 status
type httpRequestRecorder struct {
	t        testing.TB
	requests []*http.Request
}

func (rec *httpRequestRecorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rec.requests = append(rec.requests, r)
}

// Single asserts that a single request was made
func (r httpRequestRecorder) Single() *http.Request {
	r.t.Helper()
	if len(r.requests) != 1 {
		r.t.Errorf("Expected single recorded request. Got: %d", len(r.requests))
	}
	return r.requests[0]
}

func (s *HTMLFormElementWithCheckboxTestSuite) TestSubmitWithCheckboxes() {
	rec := &httpRequestRecorder{t: s.T()}
	win := htmltest.NewWindowHelper(s.T(), NewWindowFromHandler(rec))
	win.LoadHTML(`<body>
		<form method="post">
			<input id="check-1" name="check-me-1" type="checkbox" />
			<label id="lbl-1" for="check">Check me 1</label>
			<input id="check-2" name="check-me-2" type="checkbox" />
			<label id="lbl-2" for="check">Check me 2</label>
		</form>
	</body>`)
	form := win.HTMLDocument().QuerySelectorHTML("form").(html.HTMLFormElement)
	check := win.HTMLDocument().GetHTMLElementById("check-1").(html.HTMLInputElement)
	check.SetChecked(true)
	form.Submit()
	req := rec.Single()
	s.Assert().Equal([]string{"on"}, req.PostForm["check-me-1"])
	s.Assert().Empty(req.PostForm["check-me-2"])
}
