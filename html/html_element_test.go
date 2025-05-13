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
	t.Run("Tabindex", testHTMLElementTabIndex)
	t.Run("Autofocus", testHTMLElementAutofocus)
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

	assert.Equal(t, []string{"foo", "fooBar"}, target.Dataset().Keys())

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

func testHTMLElementTabIndex(t *testing.T) {
	t.Parallel()
	doc := htmltest.ParseHTMLDocumentHelper(t,
		`<body>
			<div id="target-1"><div>
			<div id="target-2" tabindex="1"><div>
			<div id="target-3" tabindex="foo"><div>
		</body>`,
	)

	assert.Equal(t, -1, doc.GetHTMLElementById("target-1").TabIndex(), "No tabindex")
	assert.Equal(t, -1, doc.GetHTMLElementById("target-3").TabIndex(), "No tabindex='foo'")
	assert.Equal(t, 1, doc.GetHTMLElementById("target-2").TabIndex(), "No tabindex='1'")

	// Set an invalid index sets the content attribute to zero
	div := doc.CreateHTMLElement("div")
	div.SetTabIndex(1)
	g := gomega.NewWithT(t)
	g.Expect(div).To(HaveAttribute("tabindex", "1"), "tabindex content attribute")
}

func testHTMLElementAutofocus(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	testHelper := htmltest.HTMLHelper{T: t}
	{
		//
		div := testHelper.CreateHTMLElement("div")

		assert.False(div.Autofocus(), "Autofocus on empty property")

		// Content attribute values, not "false" does not become "true"
		for _, val := range []string{"", "foo", "false"} {
			div.SetAttribute("autofocus", val)
			assert.Truef(div.Autofocus(), "Autofocus content attribute: %s", val)
		}
	}

	{
		// Set value
		div := testHelper.CreateHTMLElement("div")

		div.SetAutofocus(true)
		assert.True(div.Autofocus(), "Autofocus IDL attribute")
		assert.Equal("", div.AttributeValue("autofocus"), "Autofocus IDL attribute")

		div.SetAutofocus(false)
		assert.False(div.Autofocus(), "Autofocus IDL attribute")
		assert.False(div.HasAttribute("autofocus"), "Autofocus content attribute exists")
	}
}
