package dom_test

import (
	"testing"

	"github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
)

func TestElementClickEventPropagation(t *testing.T) {
	t.Parallel()
	var (
		targetCalled bool
		parentCalled bool
	)
	doc := htmltest.ParseHTMLDocumentHelper(t,
		`<body><div id="parent"><div id="target"></div></div></body>`,
	)
	target := doc.GetHTMLElementById("target")
	target.
		AddEventListener("click", eventtest.NewTestHandler(func(e *event.Event) {
			targetCalled = true
			e.StopPropagation()
		}))
	doc.GetElementById("parent").
		AddEventListener("click", eventtest.NewTestHandler(func(e *event.Event) {
			parentCalled = true
		}))
	target.Click()
	assert.True(t, targetCalled, "Target handler called")
	assert.False(t, parentCalled, "Parent handler called")
}

func TestElementClickEventPreventDefault(t *testing.T) {
	t.Parallel()
	var (
		targetCalled bool
		parentCalled bool
	)
	doc := htmltest.ParseHTMLDocumentHelper(
		t,
		`<body><div id="parent"><div id="target"></div></div></body>`,
	)
	target := doc.GetHTMLElementById("target")
	target.
		AddEventListener("click", eventtest.NewTestHandler(func(e *event.Event) {
			targetCalled = true
			e.PreventDefault()
		}))
	doc.GetElementById("parent").
		AddEventListener("click", eventtest.NewTestHandler(func(e *event.Event) {
			parentCalled = true
		}))
	target.Click()
	assert.True(t, targetCalled, "Target handler called")
	assert.True(t, parentCalled, "Parent handler called")
}

func TestElementClickEventBubbles(t *testing.T) {
	t.Parallel()
	var (
		targetCalled bool
		parentCalled bool
	)
	doc := htmltest.ParseHTMLDocumentHelper(
		t,
		`<body><div id="parent"><div id="target"></div></div></body>`,
	)
	target := doc.GetHTMLElementById("target")
	target.AddEventListener(
		"click",
		eventtest.NewTestHandler(func(e *event.Event) { targetCalled = true }),
	)
	doc.GetElementById("parent").
		AddEventListener("click", eventtest.NewTestHandler(func(e *event.Event) { parentCalled = true }))
	target.Click()
	assert.True(t, targetCalled, "Target handler called")
	assert.True(t, parentCalled, "Parent handler called")
}
