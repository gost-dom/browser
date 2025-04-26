package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/onsi/gomega"
)

type LocationSuite struct {
	ScriptHostSuite
}

func NewLocationSuite(h html.ScriptHost) *LocationSuite {
	return &LocationSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *LocationSuite) TestGlobalScope() {
	s.Expect(s.Eval("typeof Location")).To(Equal("function"))
	s.Expect(s.Eval("Object.getPrototypeOf(location) === Location"))
}

func (s *LocationSuite) TestHrefEqualsDocumentLocation() {
	window := html.NewWindow(
		html.WindowOptions{
			BaseLocation: "http://example.com/foo",
			ScriptHost:   s.scriptHost,
		})
	s.Expect(window.Eval("location.href")).To(Equal("http://example.com/foo"))
}
