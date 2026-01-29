package domsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
)

func RunDomSuite(t *testing.T, e html.ScriptEngine) {
	wrapSuite := func(test func(*testing.T, html.ScriptEngine)) func(*testing.T) {
		return func(t *testing.T) {
			t.Parallel()
			test(t, e)
		}
	}

	t.Run("Document", wrapSuite(testDocument))
	t.Run("EventTarget", wrapSuite(testEventTarget))
	t.Run("Text", wrapSuite(testText))
}
