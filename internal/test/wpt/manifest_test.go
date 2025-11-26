package main_test

import (
	"strings"
	"testing"

	. "github.com/gost-dom/browser/internal/test/wpt"

	"github.com/stretchr/testify/assert"
)

func TestManifestFile(t *testing.T) {
	// This isn't a _test_ - it was merely a tool to get feedback while
	// implementing JSON deserialization of the manifest
	data, err := LoadManifest()
	if !assert.NoError(t, err) {
		return
	}

	i := 0
	for testCase := range data.All() {
		if !strings.HasPrefix(testCase.Path, "dom") {
			continue
		}
		i++
		if i > 100 {
			break
		}
		t.Log(testCase.Path)
	}
}
