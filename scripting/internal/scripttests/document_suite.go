package scripttests

import (
	"github.com/gost-dom/browser/html"
)

type DocumentTestSuite struct {
	ScriptHostSuite
}

func NewDocumentSuite(h html.ScriptEngine) *DocumentTestSuite {
	return &DocumentTestSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *DocumentTestSuite) TestLocationAttribute() {
	s.Assert().
		Equal(true,
			s.MustEval("document.location instanceof Location"),
			"document.location instanceof Location")
	s.Assert().
		Equal(true,
			s.MustEval("document.location === location"),
			"document.location is same as global location object")
}
