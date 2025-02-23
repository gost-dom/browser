package scripttests

import (
	"github.com/gost-dom/browser/html"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type GomegaSuite struct {
	suite.Suite
	gomega.Gomega
}

func (s *GomegaSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

type ScriptHostSuite struct {
	GomegaSuite
	scriptHost html.ScriptHost
	window     html.Window
}

func (s *ScriptHostSuite) SetupTest() {
	s.GomegaSuite.SetupTest()
	s.window = html.NewWindow(html.WindowOptions{
		ScriptHost: s.scriptHost,
	})
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

type LocationSuite struct {
	ScriptHostSuite
}

func NewLocationSuite(h html.ScriptHost) *LocationSuite {
	return &LocationSuite{ScriptHostSuite: *NewScriptHostSuite(h)}
}

func (s *LocationSuite) TestGlobalScope() {
	s.Expect(s.eval("typeof Location")).To(Equal("function"))
	s.Expect(s.eval("Object.getPrototypeOf(location) === Location"))
}

func (s *LocationSuite) TestHrefEqualsDocumentLocation() {
	window := html.NewWindow(
		html.WindowOptions{
			BaseLocation: "http://example.com/foo",
			ScriptHost:   s.scriptHost,
		})
	s.Expect(window.Eval("location.href")).To(Equal("http://example.com/foo"))
}
