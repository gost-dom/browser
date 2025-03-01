package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/suite"
)

func runSuite(s suite.TestingSuite) func(t *testing.T) {
	return func(t *testing.T) { suite.Run(t, s) }
}

func RunSuites(t *testing.T, h html.ScriptHost) {
	t.Run("Location", runSuite(NewLocationSuite(h)))
	t.Run("EventLoop", runSuite(NewEventLoopTestSuite(h)))
	t.Run("Window", runSuite(NewWindowTestSuite(h)))
	t.Run("UIEvents", runSuite(NewUIEventTestSuite(h)))
}
