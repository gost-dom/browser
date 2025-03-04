package html_test

import (
	. "github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/testing/gomega-matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTMLDocument", func() {
	Describe("Empty document created by `NewDocument`", func() {
		It("Should have an HTML document element", func() {
			doc := NewHTMLDocument(nil)
			Expect(doc.DocumentElement()).To(HaveTag("HTML"))
		})

		It("Should have an empty HEAD", func() {
			doc := NewHTMLDocument(nil)
			Expect(doc.Head()).To(HaveTag("HEAD"))
		})

		It("Should have a BODY", func() {
			doc := NewHTMLDocument(nil)
			Expect(doc.Body()).To(HaveTag("BODY"))
		})
	})

	Describe("CreateElement", func() {
		It("Should create the right subclass for an 'a' tag", func() {
			doc := NewHTMLDocument(nil)
			e := doc.CreateElement("a")
			_, ok := e.(HTMLAnchorElement)
			Expect(ok).To(BeTrue(), "Element was an HTMLAnchorElement")
		})

		It("Should create the right subclass for an 'A' tag", func() {
			doc := NewHTMLDocument(nil)
			e := doc.CreateElement("A")
			_, ok := e.(HTMLAnchorElement)
			Expect(ok).To(BeTrue(), "Element was an HTMLAnchorElement")
		})
	})
})
