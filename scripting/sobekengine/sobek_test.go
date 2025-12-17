package sobekengine

import (
	"testing"

	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/js"
	"github.com/gost-dom/browser/scripting/internal/testing/jsassert"
)

var assertEngine = newEngine(
	internal.DefaultInitializer[jsTypeParam](),
	js.ConfigurerFunc[jsTypeParam](jsassert.Configure[jsTypeParam]),
)

func TestSobekHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, assertEngine)
}

func TestESMSupport(t *testing.T) {
	scripttests.RunModuleSuite(t, DefaultEngine())
}

func TestScriptFileSupport(t *testing.T) {
	scripttests.RunDownloadScriptSuite(t, DefaultEngine())
}

func TestHTMX(t *testing.T) {
	scripttests.RunHtmxTests(t, DefaultEngine())
}

func TestDatastar(t *testing.T) {
	scripttests.RunDataStarTests(t, DefaultEngine())
}
