package sobekhost_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/sobekhost"
)

func TestSobekHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, sobekhost.GojaEngine{})
}

func TestESMSupport(t *testing.T) {
	scripttests.RunModuleSuite(t, sobekhost.GojaEngine{})
}
