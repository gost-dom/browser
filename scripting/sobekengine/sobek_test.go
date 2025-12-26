package sobekengine

import (
	"testing"

	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/gost-dom/browser/scripting/internal/testing/jsassert"
)

var assertEngine *scriptEngine

func TestSobekHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, assertEngine)
}

func TestESMSupport_Sobek(t *testing.T) {
	scripttests.RunModuleSuite(t, DefaultEngine())
}

func TestScriptFileSupport(t *testing.T) {
	scripttests.RunDownloadScriptSuite(t, DefaultEngine())
}

func TestHTMX_Sobek(t *testing.T) {
	scripttests.RunHtmxTests(t, DefaultEngine())
}

func TestDatastar(t *testing.T) {
	scripttests.RunDataStarTests(t, DefaultEngine())
}

func TestBasics(t *testing.T) {
	scripttests.RunBasicSuite(t, assertEngine)
}

func init() {
	configurer := internal.CreateWindowsConfigurer[jsTypeParam]()
	configurer.AddConfigurerFunc(jsassert.Configure)
	assertEngine = newEngine(configurer)
}
