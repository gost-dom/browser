package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
)

func TestSelectElementValue(t *testing.T) {
	win := htmltest.NewWindowHTML(t, `
		<body>
			<form method="post">
				<select id="field-id" name="field">
					<option id="opt1" value="">Initial value</option>
					<option id="opt2" value="value 1" selected>Value 1</option>
					<option id="opt3" value="value 2">Value 2</option>
				</select>
			</form>
		</body>`)
	sel := htmltest.WinQuerySelector[html.HTMLSelectElement](win, "select")
	op2 := win.Document().GetElementById("opt2").(html.HTMLOptionElement)
	op3 := win.Document().GetElementById("opt3").(html.HTMLOptionElement)

	assert.True(t, op2.Selected())
	assert.False(t, op3.Selected())
	sel.SetValue("value 2")
	assert.False(t, op2.Selected())
	assert.True(t, op3.Selected())
}

func TestHTMLSelectElement_FomSubmission(t *testing.T) {
	rec := gosttest.NewHTTPRequestRecorder(t, ParseFormHandler)
	b := browsertest.InitBrowser(t, rec, nil)
	win := b.NewWindow()

	t.Run("No field is initially selected", func(t *testing.T) {
		win.LoadHTML(`
			<body>
				<form method="post">
					<select id="field-id" name="field">
						<option value="">Initial value</option>
						<option value="value 1">Value 1</option>
						<option value="value 2">Value 2</option>
					</select>
				</form>
			</body>`)
		form := htmltest.WinQuerySelector[html.HTMLFormElement](win, "form")
		form.Submit()

		req := rec.Single()
		req.ParseForm()
		assert.Equal(t, []string{""}, req.PostForm["field"])
	})

	t.Run("A field is initially selected", func(t *testing.T) {
		rec.Clear()
		win.LoadHTML(`
			<body>
				<form method="post">
					<select id="field-id" name="field">
						<option value="">Initial value</option>
						<option value="value 1" selected>Value 1</option>
						<option value="value 2">Value 2</option>
					</select>
				</form>
			</body>`)
		sel := htmltest.WinQuerySelector[html.HTMLSelectElement](win, "select")
		assert.Equal(t, "value 1", sel.Value())
		form := htmltest.WinQuerySelector[html.HTMLFormElement](win, "form")
		form.Submit()

		req := rec.Single()
		req.ParseForm()
		assert.Equal(t, []string{"value 1"}, req.PostForm["field"])
	})
}
