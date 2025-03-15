package gojahost_test

import (
	"log/slog"
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/test"
	. "github.com/gost-dom/browser/internal/test/scripttests"
	"github.com/gost-dom/browser/logger"
	. "github.com/gost-dom/browser/scripting/gojahost"

	"github.com/dop251/goja"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var testSuite = NewScriptTestSuite(New(), "goja", SkipDOM)

func newCtx() html.ScriptContext { return testSuite.NewContext() }

func FormatException(value any) (result string, ok bool) {
	var exception *goja.Exception
	if exception, ok = value.(*goja.Exception); ok {
		result = exception.String()
	}
	return
}

func init() {
	logger.SetDefault(test.CreateTestLogger(slog.LevelInfo))

	format.RegisterCustomFormatter(FormatException)
}

func TestGojaDriver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GojaDriver Suite")
}
