package domsuite

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func testText(t *testing.T, e html.ScriptEngine) {
	w := browsertest.InitWindow(t, e)
	res := w.MustEval("new Text('Text content')")
	_, ok := res.(dom.Text)
	if !assert.True(t, ok) {
		t.Logf("Not text: %T", res)
	}
}
