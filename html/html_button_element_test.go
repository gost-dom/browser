package html_test

import (
	"testing"

	. "github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
)

func TestHTMLButtonElementClick(t *testing.T) {
	doc := ParseHTMLDocumentHelper(t,
		"<body><button id='btn' type='submit'>Click me!</button></body>",
	)
	button := doc.GetHTMLElementById("btn")
	assert.True(t, button.Click())
}
