package html_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
)

func TestHTMLElement(t *testing.T) {
	t.Parallel()
	t.Run("Dataset", testHTMLElementDataset)
}

func testHTMLElementDataset(t *testing.T) {
	t.Parallel()

	doc := htmltest.ParseHTMLDocumentHelper(t,
		`<body>
			<div id="target" 
				data-foo="foo" 
				data-foo-bar="foo bar"
			></div>
		</body>`,
	)
	target := doc.GetHTMLElementById("target")

	v, ok := target.Dataset().Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "foo", v)

	v, ok = target.Dataset().Get("fooBar")
	assert.True(t, ok, "data-foo-bar found")
	assert.Equal(t, "foo bar", v, "data-foo-bar value")
}
