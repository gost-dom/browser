package scripttests

import (
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
)

type AbortControllerSuite struct {
	ScriptHostSuite
}

func NewAbortControllerSuite(h html.ScriptHost) *AbortControllerSuite {
	return &AbortControllerSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *AbortControllerSuite) TestGlobals() {
	s.Expect(s.Eval("typeof AbortController")).To(Equal("function"))
}
