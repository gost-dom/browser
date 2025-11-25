package v8engine_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/v8engine"
)

func TestDatastar(t *testing.T) {
	scripttests.RunDataStarTests(t, v8engine.DefaultEngine())
}
