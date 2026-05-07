package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
