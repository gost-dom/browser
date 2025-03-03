package dom_test

import (
	"log/slog"
	"testing"

	"github.com/gost-dom/browser/internal/test"
	"github.com/gost-dom/browser/logger"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func init() {
	logger.SetDefault(test.CreateTestLogger(slog.LevelInfo))
}

func TestDomTypes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Browser Suite")
}
