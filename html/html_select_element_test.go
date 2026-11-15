package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/gost-dom/browser/internal/testing/gosttest"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	"github.com/stretchr/testify/assert"
)

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
