package sobekhost_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/sobekhost"
)

type scriptHostFactory struct{}

func (f scriptHostFactory) NewHost(
	html.ScriptEngineOptions,
) html.ScriptHost {
	return sobekhost.New()
}

func TestSobekHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, scriptHostFactory{})
}
