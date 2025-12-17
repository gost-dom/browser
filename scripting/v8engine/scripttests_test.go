package v8engine

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

// Runs all the shared script tests using the V8 script engine
func TestV8ScriptHost(t *testing.T) {
	scripttests.RunSuites(t, assertEngine)
}
