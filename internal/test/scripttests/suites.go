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

type scriptEngineHost struct{ engine html.ScriptEngine }

// Deprecated: Get rid of this
func (h scriptEngineHost) New() html.ScriptHost {
	return h.engine.NewHost(html.ScriptEngineOptions{})
}

func RunSuites(t *testing.T, e html.ScriptEngine) {
	h := scriptEngineHost{e}
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
	t.Run("Fetch", func(t *testing.T) { testFetch(t, h.New()) })
	t.Run("Streams", func(t *testing.T) { testStreams(t, h) })
	t.Run("CharacterData", func(t *testing.T) { testCharacterData(t, h) })
}
