package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DOM Parser", func() {
	It("Should return a Document", func() {
		ctx := NewTestContext()
		Expect(ctx.Eval(`
			const parser = new DOMParser()
			const doc = parser.parseFromString("<div id='target'></div>", "text/html")
		`)).Error().ToNot(HaveOccurred())
		Expect(
			ctx.Eval("Object.getPrototypeOf(doc) === HTMLDocument.prototype"),
		).To(BeTrue(), "result is a Document")
		Expect(
			ctx.Eval("doc === window.document"),
		).To(BeFalse(), "Window.document isn't replaced")
		Expect(
			ctx.Eval("doc.getElementById('target') instanceof HTMLDivElement"),
		).To(BeTrue(), "Element is a div")
	})

	It("Should return an HTMLDocument", func() {
		ctx := NewTestContext()
		Expect(ctx.Eval(`
			const parser = new DOMParser()
			const doc = parser.parseFromString("<div id='target'></div>", "text/html")
		`)).Error().ToNot(HaveOccurred())
		Skip("HTMLDocument not properly implemented")
		Expect(
			ctx.Eval("Object.getPrototypeOf(doc) === HTMLDocument.prototype"),
		).To(BeTrue(), "result is an HTMLDocument")
	})
})
