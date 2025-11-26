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

func RunSuites(t *testing.T, e html.ScriptEngine) {
	t.Run("SharowRoot", runSuite(NewShadowRootSuite(e)))
	t.Run("DocumentFragment", runSuite(NewDocumentFragmentSuite(e)))
	t.Run("XMLHttpRequest", runSuite(NewXMLHttpRequestSuite(e)))
	t.Run("Location", runSuite(NewLocationSuite(e)))
	t.Run("EventLoop", runSuite(NewEventLoopTestSuite(e)))
	t.Run("Element", runSuite(NewElementSuite(e)))
	t.Run("Window", runSuite(NewWindowTestSuite(e)))
	t.Run("UIEvents", runSuite(NewUIEventTestSuite(e)))
	t.Run("Document", runSuite(NewDocumentSuite(e)))
	t.Run("EventTarget", runSuite(NewEventTargetTestSuite(e)))
	t.Run("FormData", runSuite(NewFormDataSuite(e)))
	t.Run("ClassList", runSuite(NewClassListTestSuite(e)))
	t.Run("Node", runSuite(NewNodeTestSuite(e)))
	t.Run("NamedNodeMap", runSuite(NewNamedNodeMapSuite(e)))
	t.Run("ElementDataset", runSuite(NewDatasetSuite(e)))
	t.Run("NodeList", runSuite(NewNodeListSuite(e)))
	t.Run("AbortController", runSuite(NewAbortControllerSuite(e)))
	t.Run("FetchSuite", runSuite(NewFetchSuite(e)))
	t.Run("Console", func(t *testing.T) { testConsole(t, e) })
	t.Run("Fetch", func(t *testing.T) { testFetch(t, e) })
	t.Run("Streams", func(t *testing.T) { testStreams(t, e) })
	t.Run("CharacterData", func(t *testing.T) { testCharacterData(t, e) })
	t.Run("ParentNode", func(t *testing.T) { testParentNode(t, e) })
	t.Run("CustomEvent", func(t *testing.T) { testCustomEvent(t, e) })
	t.Run("URL", func(t *testing.T) { testURLSuite(t, e) })
}
