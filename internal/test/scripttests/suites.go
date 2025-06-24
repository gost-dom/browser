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
	t.Run("SharowRoot", runSuite(NewShadowRootSuite(h.New())))
	t.Run("DocumentFragment", runSuite(NewDocumentFragmentSuite(h.New())))
	t.Run("XMLHttpRequest", runSuite(NewXMLHttpRequestSuite(h.New())))
	t.Run("Location", runSuite(NewLocationSuite(h.New())))
	t.Run("EventLoop", runSuite(NewEventLoopTestSuite(h.New())))
	t.Run("Window", runSuite(NewWindowTestSuite(h.New())))
	t.Run("UIEvents", runSuite(NewUIEventTestSuite(h.New())))
	t.Run("Document", runSuite(NewDocumentSuite(h.New())))
	t.Run("EventTarget", runSuite(NewEventTargetTestSuite(h.New())))
	t.Run("FormData", runSuite(NewFormDataSuite(h.New())))
	t.Run("ClassList", runSuite(NewClassListTestSuite(h.New())))
	t.Run("Node", runSuite(NewNodeTestSuite(h.New())))
	t.Run("NamedNodeMap", runSuite(NewNamedNodeMapSuite(h.New())))
	t.Run("ElementDataset", runSuite(NewDatasetSuite(h.New())))
	t.Run("NodeList", runSuite(NewNodeListSuite(h.New())))
	t.Run("AbortController", runSuite(NewAbortControllerSuite(h.New())))
	t.Run("FetchSuite", runSuite(NewFetchSuite(h.New())))
	t.Run("Fetch", func(t *testing.T) { testFetch(t, h.New()) })
}
