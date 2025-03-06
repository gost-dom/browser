package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/stretchr/testify/assert"
)

func TestElementClickEventPropagation(t *testing.T) {
	var (
		targetCalled bool
		parentCalled bool
	)
	doc := ParseHtmlString(`<body><div id="parent"><div id="target"></div></div></body>`)
	target := doc.GetElementById("target")
	target.
		AddEventListener("click", NewTestHandler(func(e *event.Event) {
			targetCalled = true
			e.StopPropagation()
		}))
	doc.GetElementById("parent").AddEventListener("click", NewTestHandler(func(e *event.Event) {
		parentCalled = true
	}))
	result := target.Click()
	assert.True(t, targetCalled, "Target handler called")
	assert.False(t, parentCalled, "Parent handler called")
	assert.True(t, result, "Click return value")
}

func TestElementClickEventPreventDefault(t *testing.T) {
	var (
		targetCalled bool
		parentCalled bool
	)
	doc := ParseHtmlString(`<body><div id="parent"><div id="target"></div></div></body>`)
	target := doc.GetElementById("target")
	target.
		AddEventListener("click", NewTestHandler(func(e *event.Event) {
			targetCalled = true
			e.PreventDefault()
		}))
	doc.GetElementById("parent").AddEventListener("click", NewTestHandler(func(e *event.Event) {
		parentCalled = true
	}))
	result := target.Click()
	assert.True(t, targetCalled, "Target handler called")
	assert.True(t, parentCalled, "Parent handler called")
	assert.False(t, result, "Click return value")
}

func TestElementClickEventBubbles(t *testing.T) {
	var (
		targetCalled bool
		parentCalled bool
	)
	doc := ParseHtmlString(`<body><div id="parent"><div id="target"></div></div></body>`)
	target := doc.GetElementById("target")
	target.AddEventListener("click", NewTestHandler(func(e *event.Event) { targetCalled = true }))
	doc.GetElementById("parent").
		AddEventListener("click", NewTestHandler(func(e *event.Event) { parentCalled = true }))
	result := target.Click()
	assert.True(t, targetCalled, "Target handler called")
	assert.True(t, parentCalled, "Parent handler called")
	assert.True(t, result, "Click return value")
}
