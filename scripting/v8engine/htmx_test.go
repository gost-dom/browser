package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/scripting/internal/scripttests"
	"github.com/gost-dom/browser/scripting/v8engine"
)

func TestHTMX(t *testing.T) {
	scripttests.RunHtmxTests(t, v8engine.DefaultEngine())
}
