package scripttests

import (
	"log/slog"
	"net/http"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/gosthttp"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
)

type ScriptHostSuite struct {
	gosttest.GomegaSuite
	engine html.ScriptEngine
	host   html.ScriptHost
	Window htmltest.WindowHelper
}

func (s *ScriptHostSuite) MustLoadHTML(html string) {
	s.T().Helper()
	s.Assert().NoError(s.Window.LoadHTML(html))
}

func (s *ScriptHostSuite) Host() html.ScriptHost {
	if s.host == nil {
		s.host = s.engine.NewHost(html.ScriptEngineOptions{})
	}
	return s.host
}
func (s *ScriptHostSuite) SetupTest() {
	s.Window = htmltest.NewWindowHelper(s.T(), html.NewWindow(html.WindowOptions{
		Logger:     gosttest.NewTestLogger(s.T()),
		ScriptHost: s.Host(),
	}))
}

// NewWindowLocation replaces the window. The new window has the specified
// location, but doesn't actually load the content.
func (s *ScriptHostSuite) NewWindowLocation(location string) {
	s.Window = htmltest.NewWindowHelper(s.T(), html.NewWindow(html.WindowOptions{
		Logger:       gosttest.NewTestLogger(s.T(), gosttest.MinLogLevel(slog.LevelDebug)),
		ScriptHost:   s.Host(),
		BaseLocation: location,
	}))
}

func (s *ScriptHostSuite) OpenWindow(location string, h http.Handler) html.Window {
	s.Window = htmltest.NewWindowHelper(s.T(), html.NewWindow(html.WindowOptions{
		BaseLocation: location,
		HttpClient:   gosthttp.NewHttpClientFromHandler(h),
		Logger:       gosttest.NewTestLogger(s.T()),
		ScriptHost:   s.Host(),
	}))
	return s.Window
}

func (s *ScriptHostSuite) TeardownTest() {
	s.Window.Close()
}

func (s *ScriptHostSuite) TearDownSuite() {
	if s.host != nil {
		s.host.Close()
	}
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
