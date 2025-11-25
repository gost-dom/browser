package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/v8engine"
)

// Runs all the shared script tests using the V8 script engine
func TestV8ScriptHost(t *testing.T) {
	scripttests.RunSuites(t, v8engine.DefaultEngine())
}
