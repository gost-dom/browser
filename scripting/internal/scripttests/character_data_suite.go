package scripttests

import (
	"testing"

	"github.com/gost-dom/browser/html"
	"github.com/stretchr/testify/assert"
)

func testCharacterData(t *testing.T, e html.ScriptEngine) {
	win := initWindow(t, e, nil)
	win.LoadHTML(`<body><div /></body>`)
	bodyHTML := win.MustEval(`
		const t = document.createTextNode("foo")
		const b = document.body
		b.insertBefore(t, b.firstChild)
		b.outerHTML
	`)
	assert.Equal(t, `<body>foo<div></div></body>`, bodyHTML)
	assert.Equal(t, "Text", win.MustEval(`Object.getPrototypeOf(t).constructor.name`))

	bodyHTML = win.MustEval(`
		const t2 = t.cloneNode()
		b.insertBefore(t2)
		b.outerHTML
	`)
	assert.Equal(t, `<body>foo<div></div>foo</body>`, bodyHTML)
}
