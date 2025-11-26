package sobekengine_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/sobekengine"
)

func TestSobekHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, sobekengine.DefaultEngine())
}

func TestESMSupport(t *testing.T) {
	scripttests.RunModuleSuite(t, sobekengine.DefaultEngine())
}

func TestScriptFileSupport(t *testing.T) {
	scripttests.RunDownloadScriptSuite(t, sobekengine.DefaultEngine())
}

func TestHTMX(t *testing.T) {
	scripttests.RunHtmxTests(t, sobekengine.DefaultEngine())
}

func TestDatastar(t *testing.T) {
	scripttests.RunDataStarTests(t, sobekengine.DefaultEngine())
}
