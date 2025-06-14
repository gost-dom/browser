package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/suite"
)

func runSuite(s suite.TestingSuite) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		suite.Run(t, s)
	}
}

type ScriptHostFactory interface{ New() html.ScriptHost }

func RunSuites(t *testing.T, h ScriptHostFactory) {
	t.Run("Location", runSuite(NewLocationSuite(h.New())))
	t.Run("EventLoop", runSuite(NewEventLoopTestSuite(h.New())))
	t.Run("Window", runSuite(NewWindowTestSuite(h.New())))
	t.Run("UIEvents", runSuite(NewUIEventTestSuite(h.New())))
	t.Run("Document", runSuite(NewDocumentSuite(h.New())))
	t.Run("EventTarget", runSuite(NewEventTargetTestSuite(h.New())))
	t.Run("ClassList", runSuite(NewClassListTestSuite(h.New())))
	t.Run("Node", runSuite(NewNodeTestSuite(h.New())))
	t.Run("ElementDataset", runSuite(NewDatasetSuite(h.New())))
}
