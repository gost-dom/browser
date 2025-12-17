package htmlsuite

import (
	"testing"

	"github.com/gost-dom/browser/html"
)

func RunHtmlSuite(t *testing.T, e html.ScriptEngine) {
	wrapSuite := func(test func(*testing.T, html.ScriptEngine)) func(*testing.T) {
		return func(t *testing.T) { test(t, e) }
	}

	t.Run("HTMLAnchorElement", wrapSuite(testHtmlAnchorElement))
}
