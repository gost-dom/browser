package sobekhost_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/sobekhost"
)

func TestSobekHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, sobekhost.DefaultEngine())
}

func TestESMSupport(t *testing.T) {
	scripttests.RunModuleSuite(t, sobekhost.DefaultEngine())
}

func TestScriptFileSupport(t *testing.T) {
	scripttests.RunDownloadScriptSuite(t, sobekhost.DefaultEngine())
}

func TestHTMX(t *testing.T) {
	scripttests.RunHtmxTests(t, sobekhost.DefaultEngine())
}

func TestDatastar(t *testing.T) {
	scripttests.RunDataStarTests(t, sobekhost.DefaultEngine())
}
