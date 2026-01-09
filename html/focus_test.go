package html_test

import (
	"fmt"
	"testing"

	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom/event"
	"github.com/gost-dom/browser/internal/testing/eventtest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
)

func TestFocusEvents(t *testing.T) {
	assert := assert.New(t)
	doc := htmltest.ParseHTMLDocument(t, activeDocumentTestHtml)
	var logs []string
	var logHandler = func(name string) EventHandler {
		return eventtest.NewTestHandler(func(e *Event) {
			logs = append(logs, fmt.Sprintf("%s - %s", e.Type, name))
		})
	}
	nodes, err := doc.QuerySelectorAll("div, input, button")
	assert.NoError(err)
	for e := range nodes.All() {
		id, _ := (e.(Element).GetAttribute("id"))
		e.AddEventListener("focus", logHandler(id))
		e.AddEventListener("blur", logHandler(id))
		e.AddEventListener("focusin", logHandler(id))
		e.AddEventListener("focusout", logHandler(id))
	}
	target := doc.GetHTMLElementById("input-1-1-1")
	target.Focus()
	assert.Equal([]string{
		"focus - input-1-1-1",
		"focusin - input-1-1-1",
		"focusin - child-1-1",
		"focusin - root-1",
	}, logs)
	assert.Equal(target, doc.ActiveElement())

	logs = nil
	newTarget := doc.GetHTMLElementById("button-1-2-1")
	newTarget.Focus()

	assert.Equal([]string{
		"blur - input-1-1-1",
		"focusout - input-1-1-1",
		"focusout - child-1-1",
		"focusout - root-1",

		"focus - button-1-2-1",
		"focusin - button-1-2-1",
		"focusin - child-1-2",
		"focusin - root-1",
	}, logs)

	logs = nil
	newTarget.Blur()
	assert.Equal([]string{
		"blur - button-1-2-1",
		"focusout - button-1-2-1",
		"focusout - child-1-2",
		"focusout - root-1",
	}, logs)
	assert.Equal(doc.Body(), doc.ActiveElement())
}

func TestAutofocus(t *testing.T) {
	doc := htmltest.ParseHTMLDocument(t,
		`<body><form><input id="i" name="i" type="text" autofocus/></form></body>`,
	)
	input := doc.GetElementById("i")
	assert.Equal(t, input, doc.ActiveElement())
}

const activeDocumentTestHtml = `<body>
	<div id="root-1">
		<div id="child-1-1">
			<input id="input-1-1-1" type="text" />
			<input id="input-1-1-2" type="text" />
		</div>
		<div id="child-1-2">
			<button id="button-1-2-1">Button 1</button>
		</div>
	</div>
<body>`
