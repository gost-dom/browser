package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/stretchr/testify/assert"
)

func TestAttribute_Namespace(t *testing.T) {
	doc := CreateHTMLDocument()

	t.Run("Attribute with no namespace", func(t *testing.T) {
		attr, err := doc.CreateAttribute("attr")
		assert.NoError(t, err)
		assert.Equal(t, "", attr.NamespaceURI())
		assert.Equal(t, "attr", attr.LocalName())
		assert.Equal(t, "attr", attr.Name())
		assert.Equal(t, "attr", attr.NodeName())
		assert.Equal(t, "", attr.Prefix())
	})

	t.Run("Attr has no prefix", func(t *testing.T) {
		attr, err := doc.CreateAttributeNS("http://example.com/gost-dom", "attr")
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com/gost-dom", attr.NamespaceURI())
		assert.Equal(t, "attr", attr.LocalName())
		assert.Equal(t, "attr", attr.Name())
		assert.Equal(t, "attr", attr.NodeName())
		assert.Equal(t, "", attr.Prefix())
	})

	t.Run("Attr has prefix", func(t *testing.T) {
		attr, err := doc.CreateAttributeNS("http://example.com/gost-dom", "gost:attr")
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com/gost-dom", attr.NamespaceURI())
		assert.Equal(t, "attr", attr.LocalName())
		assert.Equal(t, "gost:attr", attr.Name())
		assert.Equal(t, "gost:attr", attr.NodeName())
		assert.Equal(t, "gost", attr.Prefix())
	})

	t.Run("Attr has ':' in local name", func(t *testing.T) {
		attr, err := doc.CreateAttributeNS("http://example.com/gost-dom", "gost:attr:x")
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com/gost-dom", attr.NamespaceURI())
		assert.Equal(t, "attr:x", attr.LocalName())
		assert.Equal(t, "gost:attr:x", attr.Name())
		assert.Equal(t, "gost:attr:x", attr.NodeName())
		assert.Equal(t, "gost", attr.Prefix())
	})

	t.Run("Attr has 'xml' prefix", func(t *testing.T) {
		_, err1 := doc.CreateAttributeNS(
			"http://www.w3.org/XML/1998/namespace",
			"xml:something",
		)
		assert.NoError(t, err1)
		_, err2 := doc.CreateAttributeNS("http://example.com/wrong-namespace", "xml:something")
		assert.ErrorIs(t, err2, dom.ErrDom)
		assert.ErrorIs(t, err2, dom.ErrInvalidCharacter)
	})

	t.Run("Attr has 'xmlns' prefix", func(t *testing.T) {
		_, err1 := doc.CreateAttributeNS(
			"http://www.w3.org/2000/xmlns/",
			"xmlns:something",
		)
		assert.NoError(t, err1)
		_, err2 := doc.CreateAttributeNS("http://example.com/wrong-namespace", "xmlns:something")
		assert.ErrorIs(t, err2, dom.ErrDom)
		assert.ErrorIs(t, err2, dom.ErrInvalidCharacter)
	})
}
