package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/v8host"
)

type scriptHostFactory struct{}

func (f scriptHostFactory) New() html.ScriptHost { return v8host.New() }

// Runs all the shared script tests using the V8 script engine
func TestV8ScriptHost(t *testing.T) {
	scripttests.RunSuites(t, &v8host.DefaultEngine)
}
