package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDocumentType_AllowedPosition(t *testing.T) {
	doc := dom.NewDocument(nil)
	require.NoError(t, doc.Append(doc.CreateDocumentType("HTML")))
	require.ErrorIs(t, doc.Append(doc.CreateDocumentType("HTML")), dom.ErrDom)
}

func TestDocumentType_IsEqualNode(t *testing.T) {
	doc := dom.NewDocument(nil)
	t1 := doc.CreateDocumentType("HTML")
	t2 := doc.CreateDocumentType("HTML")
	t3 := doc.CreateDocumentType("NOT-HTML")
	el := doc.CreateElement("div")

	require.True(t, t1.IsEqualNode(t2))
	require.False(t, t2.IsEqualNode(t3))
	require.False(t, t2.IsEqualNode(el))
}

func TestDocumentTypeNodeType(t *testing.T) {
	doc := ParseHtmlString(`<!DOCTYPE html>
<html><head></head><body>
</body></html>`)
	node := doc.ChildNodes().Item(0)
	docType, ok := node.(dom.DocumentType)
	require.True(t, ok, "Node is a DocumentType")
	assert.Equal(t, dom.NodeTypeDocumentType, docType.NodeType())
	assert.Equal(t, "html", docType.NodeName(), "DocumentType.NodeName()")
	assert.Equal(t, "html", docType.Name(), "DocumentType.Name()")
}
