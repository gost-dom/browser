package html_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestHTMLElement(t *testing.T) {
	t.Parallel()
	t.Run("Dataset", testHTMLElementDataset)
	t.Run("Tabindex", testHTMLElementTabindex)
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

	target.Dataset().Set("fooBar", "New value")
	v, ok = target.Dataset().Get("fooBar")
	assert.True(t, ok, "data-foo-bar found")
	assert.Equal(t, "New value", v, "data-foo-bar value")

	target.Dataset().Delete("foo-bar")
	_, ok = target.Dataset().Get("fooBar")
	assert.False(t, ok, "data-foo-bar found after delete")
}

func testHTMLElementTabindex(t *testing.T) {
	t.Parallel()
	doc := htmltest.ParseHTMLDocumentHelper(t,
		`<body>
			<div id="target-1"><div>
			<div id="target-2" tabindex="1"><div>
			<div id="target-3" tabindex="foo"><div>
		</body>`,
	)

	assert.Equal(t, -1, doc.GetHTMLElementById("target-1").Tabindex(), "No tabindex")
	assert.Equal(t, -1, doc.GetHTMLElementById("target-3").Tabindex(), "No tabindex='foo'")
	assert.Equal(t, 1, doc.GetHTMLElementById("target-2").Tabindex(), "No tabindex='1'")

	// Set an invalid index sets the content attribute to zero
	div := doc.CreateHTMLElement("div")
	div.SetTabindex(1)
	g := gomega.NewWithT(t)
	g.Expect(div).To(HaveAttribute("tabindex", "1"), "tabindex content attribute")
}
