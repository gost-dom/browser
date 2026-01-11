package scripttests

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	. "github.com/onsi/gomega"
)

type LocationSuite struct {
	ScriptHostSuite
}

func NewLocationSuite(h html.ScriptEngine) *LocationSuite {
	return &LocationSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *LocationSuite) TestGlobalScope() {
	s.Expect(s.Eval("typeof Location")).To(Equal("function"))
	s.Expect(s.Eval("Object.getPrototypeOf(location) === Location"))
}

func (s *LocationSuite) TestHrefEqualsDocumentLocation() {
	b := browsertest.InitBrowser(s.T(), nil, s.engine)
	window := b.OpenWindow("http://example.com/foo")
	s.Expect(window.Eval("location.href")).To(Equal("http://example.com/foo"))
}
