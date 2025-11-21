package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/v8host"
)

func TestHTMX(t *testing.T) {
	scripttests.RunHtmxTests(t, v8host.DefaultEngine())
}
