package gojahost_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/gojahost"
)

type scriptHostFactory struct{}

func (f scriptHostFactory) New() html.ScriptHost { return gojahost.New() }

func TestGojaHost(t *testing.T) {
	t.Parallel()
	scripttests.RunSuites(t, scriptHostFactory{})
}
