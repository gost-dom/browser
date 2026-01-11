package scripttests

import (
	"net/http"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

type ScriptHostSuite struct {
	gosttest.GomegaSuite
	engine html.ScriptEngine
	Window htmltest.WindowHelper
}

func (s *ScriptHostSuite) MustLoadHTML(html string) {
	s.T().Helper()
	s.Assert().NoError(s.Window.LoadHTML(html))
}

func (s *ScriptHostSuite) SetupTest() {
	s.Window = browsertest.InitWindow(s.T(), s.engine)
}

// NewWindowLocation replaces the window. The new window has the specified
// location, but doesn't actually load the content.
func (s *ScriptHostSuite) NewWindowLocation(location string) {
	b := browsertest.InitBrowser(s.T(), nil, s.engine)
	s.Window = b.OpenWindow(location)
}

func (s *ScriptHostSuite) OpenWindow(location string, h http.Handler) html.Window {
	b := browsertest.InitBrowser(s.T(), h, s.engine)
	s.Window = b.OpenWindow(location)
	return s.Window
}

func (s *ScriptHostSuite) TeardownTest() {
	s.Window.Close()
}

// Runs a script and returns the evaluated value as a native Go value.
//
// Panics (or generates an error?) if no suitable conversion could be found (i.e.
// gost doesn't implement this yet).
//
// Returns an error if script code throws.
//
// If the return value is not used, call run; to avoid panic/error
func (s *ScriptHostSuite) Eval(script string) (any, error) {
	return s.Window.Eval(script)
}

func (s *ScriptHostSuite) MustEval(script string) any {
	s.T().Helper()
	res, err := s.Eval(script)
	s.Assert().NoError(err)
	return res
}

// Runs a script, and discards the returned value.
//
// Returns an error if the script code throws. Named RunScript to not shadow
// [Suite.Run].
func (s *ScriptHostSuite) RunScript(script string) error {
	return s.Window.Run(script)
}

// MustRunScript runs a script and marks the test as an error if an error occurs.
func (s *ScriptHostSuite) MustRunScript(script string) {
	s.T().Helper()
	s.Assert().NoError(s.RunScript(script))
}

func NewScriptHostSuite(e html.ScriptEngine) *ScriptHostSuite {
	return &ScriptHostSuite{engine: e}
}
