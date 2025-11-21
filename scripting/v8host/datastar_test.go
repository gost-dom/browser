package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/v8host"
)

func TestDatastar(t *testing.T) {
	scripttests.RunDataStarTests(t, v8host.DefaultEngine())
}
