package v8engine

import (
	"testing"

	"github.com/gost-dom/browser/scripting/internal"
	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/gost-dom/browser/scripting/internal/testing/jsassert"
)

var assertEngine *v8ScriptEngine

// Runs all the shared script tests using the V8 script engine
func TestV8ScriptHost(t *testing.T) {
	scripttests.RunSuites(t, assertEngine)
}

func init() {
	configurer := internal.CreateWindowsConfigurer[jsTypeParam]()
	configurer.AddConfigurerFunc(jsassert.Configure)
	assertEngine = newEngine(configurer)
}
