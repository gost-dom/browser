package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/scripting/internal/dom/domsuite"
	"github.com/gost-dom/browser/scripting/internal/html/htmlsuite"
	"github.com/gost-dom/browser/scripting/internal/uievents/uieventssuite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func runSuite(s suite.TestingSuite) func(t *testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		suite.Run(t, s)
	}
}

func RunBasicSuite(t *testing.T, e html.ScriptEngine) {
	var w = browsertest.InitWindow(t, e)
	assert.NotNil(t, w.MustEval("globalThis"), "globalThis is nil")
	assert.NotNil(t, w.MustEval("window"), "window is nil")
	assert.Equal(t, "Window", w.MustEval("Object.getPrototypeOf(globalThis).constructor.name"))
	assert.Equal(t, "Window", w.MustEval("Object.getPrototypeOf(window).constructor.name"))
	assert.True(t, w.MustEval("window === globalThis").(bool))
}

func RunSuites(t *testing.T, e html.ScriptEngine) {
	t.Run("ScriptEngineBehaviour", func(t *testing.T) { testScriptEngineBehaviour(t, e) })
	t.Run("ShadowRoot", runSuite(NewShadowRootSuite(e)))
	t.Run("DocumentFragment", runSuite(NewDocumentFragmentSuite(e)))
	t.Run("XMLHttpRequest", runSuite(NewXMLHttpRequestSuite(e)))
	t.Run("Location", runSuite(NewLocationSuite(e)))
	t.Run("EventLoop", runSuite(NewEventLoopTestSuite(e)))
	t.Run("Element", runSuite(NewElementSuite(e)))
	t.Run("Window", runSuite(NewWindowTestSuite(e)))
	t.Run("UIEvents", func(t *testing.T) { uieventssuite.RunUieventsSuite(t, e) })
	t.Run("FormData", runSuite(NewFormDataSuite(e)))
	t.Run("ClassList", runSuite(NewClassListTestSuite(e)))
	t.Run("Node", runSuite(NewNodeTestSuite(e)))
	t.Run("NamedNodeMap", runSuite(NewNamedNodeMapSuite(e)))
	t.Run("ElementDataset", runSuite(NewDatasetSuite(e)))
	t.Run("NodeList", runSuite(NewNodeListSuite(e)))
	t.Run("AbortController", runSuite(NewAbortControllerSuite(e)))
	t.Run("Console", func(t *testing.T) { testConsole(t, e) })
	t.Run("Fetch", func(t *testing.T) { testFetch(t, e) })
	t.Run("Streams", func(t *testing.T) { testStreams(t, e) })
	t.Run("CharacterData", func(t *testing.T) { testCharacterData(t, e) })
	t.Run("ParentNode", func(t *testing.T) { testParentNode(t, e) })
	t.Run("CustomEvent", func(t *testing.T) { testCustomEvent(t, e) })
	t.Run("URL", func(t *testing.T) { testURLSuite(t, e) })
	t.Run("DOMParser", func(t *testing.T) { testDomParser(t, e) })
	t.Run("MutationObserver", func(t *testing.T) { testMutationObserver(t, e) })
	t.Run("Error handling", func(t *testing.T) { testErrorHandling(t, e) })
	t.Run("html", func(t *testing.T) { htmlsuite.RunHtmlSuite(t, e) })
	t.Run("dom", func(t *testing.T) { domsuite.RunDomSuite(t, e) })
}
