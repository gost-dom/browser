package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom"
	"github.com/stretchr/testify/require"
)

func TestDocumentType_AllowedPosition(t *testing.T) {
	doc := dom.NewDocument(nil)
	require.NoError(t, doc.Append(doc.CreateDocumentType("HTML")))
	require.ErrorIs(t, doc.Append(doc.CreateDocumentType("HTML")), dom.ErrDom)
}
