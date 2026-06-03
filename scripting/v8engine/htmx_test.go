package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/gost-dom/browser/scripting/v8engine"
)

func TestIntegration_V8(t *testing.T) {
	scripttests.RunHtmxTests(t, v8engine.DefaultEngine())
	scripttests.RunCodeMirrorTests(t, v8engine.DefaultEngine())
}
