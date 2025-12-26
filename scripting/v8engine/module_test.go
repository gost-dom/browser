package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/gost-dom/browser/scripting/v8engine"
)

func TestESMSupport_V8(t *testing.T) {
	scripttests.RunModuleSuite(t, v8engine.DefaultEngine())
}
