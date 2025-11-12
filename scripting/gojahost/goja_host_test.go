package gojahost_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/gojahost"
)

type scriptHostFactory struct{}

// TODO: Fix this
func (f scriptHostFactory) NewHost(
	html.ScriptEngineOptions,
) html.ScriptHost {
	return gojahost.New()
}

func TestGojaHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, scriptHostFactory{})
}
