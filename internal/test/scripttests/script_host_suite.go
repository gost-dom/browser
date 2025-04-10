package scripttests

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

type ScriptHostSuite struct {
	gosttest.GomegaSuite
	scriptHost html.ScriptHost
	window     htmltest.WindowHelper
}

func (s *ScriptHostSuite) mustLoadHTML(html string) {
	s.T().Helper()
	s.Assert().NoError(s.window.LoadHTML(html))
}

func (s *ScriptHostSuite) SetupTest() {
	s.window = htmltest.NewWindowHelper(s.T(), html.NewWindow(html.WindowOptions{
		ScriptHost: s.scriptHost,
	}))
}

func (s *ScriptHostSuite) OpenWindow(location string) html.Window {
	err := s.window.Navigate(location)
	s.Assert().NoError(err)
	return s.window
}

func (s *ScriptHostSuite) TeardownTest() {
	s.window.Close()
}

// Runs a script and returns the evaluated value as a native Go value.
//
// Panics (or generates an error?) if no suitable conversion could be found (i.e.
// gost doesn't implement this yet).
//
// Returns an error if script code throws.
//
// If the return value is not used, call run; to avoid panic/error
func (s *ScriptHostSuite) eval(script string) (any, error) {
	return s.window.Eval(script)
}

func (s *ScriptHostSuite) mustRun(script string) {
	s.T().Helper()
	s.Assert().NoError(s.run(script))
}

func (s *ScriptHostSuite) mustEval(script string) any {
	s.T().Helper()
	res, err := s.eval(script)
	s.Assert().NoError(err)
	return res
}

// Runs a script, and discards the returned value.
//
// Returns an error if the script code throws.
func (s *ScriptHostSuite) run(script string) error {
	return s.window.Run(script)
}

func NewScriptHostSuite(h html.ScriptHost) *ScriptHostSuite {
	return &ScriptHostSuite{
		scriptHost: h,
	}
}
