package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/scripting/v8host"
)

func TestESMSupport(t *testing.T) {
	scripttests.RunModuleSuite(t, v8host.DefaultEngine())
}
