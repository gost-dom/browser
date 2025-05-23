package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestMutationObserver(t *testing.T) {
	win := initBrowser(t, nil).NewWindow()
	assert.NoError(t, win.LoadHTML(`<body>
		<div id="parent"><div id="child"></div><div>
	</body>`))

	assert.NoError(t, win.Run(`
		const events = []
		MutationRecord.prototype.toJSON = function() { return {type: this.type} }
		const parent = document.getElementById("parent")
		const obs = new MutationObserver(recs => {
			for (const rec of recs) {
				events.push(rec.type)
			}
		})
		obs.observe(parent, {
			childList: true, subtree: true
		})
		parent.appendChild(document.createElement("div"))
	`))
	g := gomega.NewWithT(t)
	g.Expect(win.Eval("events.join(',')")).To(Equal("childList"))

	parent := win.Document().GetElementById("parent").(html.HTMLElement)
	parent.SetInnerHTML(`<div id="bar"></div>`)
	win.Clock().RunAll()

	g.Expect(win.Eval("events.join(',')")).To(Equal("childList,childList"))
}
