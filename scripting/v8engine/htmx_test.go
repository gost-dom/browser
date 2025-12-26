package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/gost-dom/browser/scripting/v8engine"
)

func TestHTMX_V8(t *testing.T) {
	scripttests.RunHtmxTests(t, v8engine.DefaultEngine())
}
