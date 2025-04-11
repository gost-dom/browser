package html_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
)

func TestHTMLLabelElement(t *testing.T) {
	var inputClicked bool
	doc := htmltest.NewHTMLDocumentHelper(t, nil)
	input := doc.CreateElement("input")
	input.SetAttribute("id", "input")
	label := doc.CreateHTMLElement("label").(html.HTMLLabelElement)
	label.SetHTMLFor("input")
	doc.Body().Append(input, label)
	input.AddEventListener("click", eventtest.NewTestHandler(func(*event.Event) {
		inputClicked = true
	}))
	label.Click()

	assert.True(t, inputClicked, "Clicking a label should cause the input to be clicked")

	inputClicked = false
	label.AddEventListener("click", eventtest.PreventDefaultHandler())
	label.Click()
	assert.False(t, inputClicked, "Clicking a label should cause the input to be clicked")
}
