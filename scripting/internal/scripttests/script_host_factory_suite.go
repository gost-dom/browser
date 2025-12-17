package scripttests

import (
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

type ScriptHostFactorySuite struct {
	gosttest.GomegaSuite
	factory html.ScriptEngine
	host    html.ScriptHost
	Window  htmltest.WindowHelper
}

func (s *ScriptHostFactorySuite) SetupTest() {
	logger := gosttest.NewTestLogger(s.T())
	s.host = s.factory.NewHost(html.ScriptEngineOptions{
		Logger: logger,
	})
	s.Window = htmltest.NewWindowHelper(s.T(), html.NewWindow(html.WindowOptions{
		Logger:     gosttest.NewTestLogger(s.T()),
		ScriptHost: s.host,
	}))
}

func NewScriptHostFactorySuite(f html.ScriptEngine) *ScriptHostFactorySuite {
	return &ScriptHostFactorySuite{factory: f}
}

// Runs a script, and discards the returned value.
//
// Returns an error if the script code throws. Named RunScript to not shadow
// [Suite.Run].
func (s *ScriptHostFactorySuite) RunScript(script string) error {
	return s.Window.Run(script)
}

// MustRunScript runs a script and marks the test as an error if an error
// occurs.
func (s *ScriptHostFactorySuite) MustRunScript(script string) {
	s.T().Helper()
	s.Assert().NoError(s.RunScript(script))
}

// Eval runs a script and returns the evaluated value as a native Go value.
//
// Returns an error if no suitable conversion could be found or if the
// conversion is not implemented.
//
// Returns an error if script code throws an exception.
//
// If the return value is not needed, you can use RunScript instead to avoid
// dealing with errors if return value conversion is not possible.
func (s *ScriptHostFactorySuite) Eval(script string) (any, error) {
	return s.Window.Eval(script)
}

func (s *ScriptHostFactorySuite) MustEval(script string) any {
	s.T().Helper()
	res, err := s.Eval(script)
	s.Assert().NoError(err)
	return res
}
