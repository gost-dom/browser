package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/suite"
)

func RunSuites(t *testing.T, h html.ScriptHost) {
	suite.Run(t, NewLocationSuite(h))
}
