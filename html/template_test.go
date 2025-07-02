package html_test

import (
	"testing"

	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/testing/gomega-matchers"
	"github.com/gost-dom/browser/internal/testing/htmltest"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

func TestHTMLTemplateElement(t *testing.T) {
	doc := htmltest.ParseHTMLDocument(t,
		`<body><template id="t"><div id="d"></div></template></body>`)

	template := doc.QuerySelectorHTML("template")
	Expect(t,
		template.ChildNodes().Length()).To(
		Equal(0),
		"template does not have it's HTML children as DOM children")

	Expect(t, doc.Body()).To(
		HaveOuterHTML(`<body><template id="t"><div id="d"></div></template></body>`),
		"Template renders its contents in body's outer HTML",
	)
	Expect(t, template).To(
		HaveOuterHTML(`<template id="t"><div id="d"></div></template>`),
		"Template renders its contents",
	)

	Expect(t, template.InnerHTML()).To(Equal(`<div id="d"></div>`))

	Expect(t, doc.QuerySelectorHTML("template")).To(Equal(template), "template can be queried")

	Expect(t, template.QuerySelectorHTMLOpt("div")).To(BeNil(), "searching <template> children")

	Expect(t, doc.QuerySelectorHTMLOpt("div")).To(
		BeNil(),
		"searching for <template> children from document",
	)
}

func TestDocumentFragmentSetInnerHTML(t *testing.T) {
	doc := htmltest.ParseHTMLDocument(t,
		`<body><template id="t"><div id="d"></div></template></body>`,
	)
	template := htmltest.UnwrapHTMLElement[html.HTMLTemplateElement](
		doc.QuerySelectorHTML("template"),
	)
	template.SetInnerHTML("<p>P1</p><p>P2</p>")
	Expect(t, template.InnerHTML()).To(Equal(
		`<p>P1</p><p>P2</p>`,
	))
}
