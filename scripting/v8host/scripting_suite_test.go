package v8host_test

import (
	"log/slog"
	"testing"

	"github.com/gost-dom/browser/internal/test"
	"github.com/gost-dom/browser/logger"
	. "github.com/gost-dom/browser/scripting/v8host"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScripting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Scripting Suite")
}

var host *V8ScriptHost

func init() {
	logger.SetDefault(test.CreateTestLogger(slog.LevelWarn))

	host = New()

	AfterSuite(func() {
		host.Close()
	})
}
