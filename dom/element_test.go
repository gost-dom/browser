package dom_test

import (
	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/testing/gomega-matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Element", func() {
	var doc Document

	BeforeEach(func() {
		doc = CreateHTMLDocument()
		DeferCleanup(func() { doc = nil })
	})

	Describe("Get/Set attribute", func() {
		It("Should add a new attribute when not existing", func() {
			doc := CreateHTMLDocument()
			elm := doc.CreateElement("div")
			Expect(elm.Attributes().Length()).To(Equal(0))
			elm.SetAttribute("id", "1")
			Expect(elm.Attributes().Length()).To(Equal(1))
		})

		It("Should add overwrite an existing attribute", func() {
			elm := doc.CreateElement("div")
			elm.SetAttribute("id", "1")
			elm.SetAttribute("id", "2")
			Expect(elm).To(HaveAttribute("id", "2"))
			Expect(elm.Attributes().Length()).To(Equal(1))
		})

		It("Should return nil when the attribute does't exist", func() {
			elm := doc.CreateElement("div")
			_, ok := elm.GetAttribute("non-existing")
			Expect(ok).To(BeFalse())
		})
	})

	Describe("Get/SetAttributeNode", func() {
		It("getAttributeNode should return nil on missing name", func() {
			elm := doc.CreateElement("div")
			attr := elm.GetAttributeNode("class")
			Expect(attr).To(BeNil())
		})

		It("getAttributeNode should return a mutable node on a valid name", func() {
			elm := doc.CreateElement("div")
			elm.SetAttribute("class", "foo")
			attr := elm.GetAttributeNode("class")
			Expect(attr).ToNot(BeNil())
			Expect(attr.Value()).To(Equal("foo"), "Attribute value before mutation")
			Expect(attr.Parent()).To(Equal(elm), "Parent on attribute node")
			attr.SetValue("bar")
			actual := elm.GetAttributeNode("class")
			Expect(actual.Value()).To(Equal("bar"), "Attribute value after mutation")
		})

		Describe("SetAttribute", func() {
			It("Should add an attribute if not existing", func() {
				elm := doc.CreateElement("div")
				attr := doc.CreateAttribute("class")
				attr.SetValue("foo")
				result, err := elm.SetAttributeNode(attr)
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(BeNil())
				Expect(elm.Attributes().Length()).To(Equal(1))
				actual, _ := elm.GetAttribute("class")
				Expect(actual).To(Equal("foo"))
			})

			It("Should replace an attribute if it already exists", func() {
				elm := doc.CreateElement("div")
				elm.SetAttribute("class", "bar")
				attr := doc.CreateAttribute("class")
				attr.SetValue("foo")
				result, err := elm.SetAttributeNode(attr)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.Name()).To(Equal("class"))
				Expect(result.Value()).To(Equal("bar"))
				Expect(elm.Attributes().Length()).To(Equal(1))
				actual, _ := elm.GetAttribute("class")
				Expect(actual).To(Equal("foo"))
			})

			It("Should return a DOMError if the attribute belongs to another element", func() {
				elm := doc.CreateElement("div")
				elm2 := doc.CreateElement("div")
				elm2.SetAttribute("class", "bar")
				attributeFromAnotherElement := elm2.GetAttributeNode("class")
				Expect(
					elm.SetAttributeNode(attributeFromAnotherElement),
				).Error().To(BeADOMError())
				Expect(elm.Attributes().Length()).To(Equal(0), "Target elm received attribute?")
				Expect(
					elm2.Attributes().Length(),
				).To(Equal(1), "Target attribute still on source elm")
			})
		})

		Describe("RemoveAttribute", func() {
			It("Should remove the attribute if it is on the element", func() {
				elm := doc.CreateElement("div")
				elm.SetAttribute("class", "bar")
				nodeToRemove := elm.GetAttributeNode("class")
				removedNode, err := elm.RemoveAttributeNode(nodeToRemove)
				Expect(err).ToNot(HaveOccurred())
				Expect(elm.Attributes().Length()).To(Equal(0))
				Expect(removedNode).To(Equal(nodeToRemove))
				Expect(removedNode.Parent()).To(BeNil(), "Attribute parent")
			})

			It("Should return a DOMError when attribute doesn't exist on the node", func() {
				elm := doc.CreateElement("div")
				elm2 := doc.CreateElement("div")
				elm2.SetAttribute("class", "bar")
				attributeFromAnotherElement := elm2.GetAttributeNode("class")
				Expect(
					elm.RemoveAttributeNode(attributeFromAnotherElement),
				).Error().To(BeADOMError())
			})
		})
	})

	Describe("Matches", func() {
		It("Should return true for a simple string matching the root element", func() {
			d := doc.CreateElement("div")
			p := doc.CreateElement("p")
			d.Append(p)
			Expect(d.Matches("div")).To(BeTrue())
		})

		It("Should return false for a simple string matching a child element", func() {
			d := doc.CreateElement("div")
			p := doc.CreateElement("p")
			d.Append(p)
			Expect(d.Matches("p")).To(BeFalse())
		})

		It("Should return true for an existing attribute", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("known-attribute", "")
			Expect(d.Matches("[known-attribute]")).To(BeTrue())
		})

		It("Should return true if one attribute match", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("known-attribute", "")
			Expect(d.Matches("[unknown-attribute], [known-attribute]")).To(BeTrue())
		})

		It("Should return false for a non-existing attribute", func() {
			d := doc.CreateElement("div")
			Expect(d.Matches("[unknown-attribute]")).To(BeFalse())
		})

		It("Should return true if tagname + attribute key=value has right value", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("a", "right")
			Expect(d.Matches(`div[a="right"]`)).To(BeTrue())
		})

		It("Should return true if tagname + attribute key=value has wrong value", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("a", "right")
			Expect(d.Matches(`div[a="wrong"]`)).To(BeFalse())
		})
	})

	It("Should support Get/SetTextContent", func() {
		d := doc.CreateElement("div")
		d.AppendChild(doc.CreateElement("p"))
		d.SetTextContent("Replace the p")
		Expect(d).To(HaveTextContent("Replace the p"))
	})

	Describe("InsertAdjacentHTML", func() {
		It("Should insert correctly 'beforeBegin'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el := doc.GetElementById("2")
			Expect(el.InsertAdjacentHTML(
				"beforebegin",
				"<div>1st new child</div><div>2nd new child</div>",
			)).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div>1st new child</div><div>2nd new child</div><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'afterBegin'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			Expect(
				el.InsertAdjacentHTML(
					"afterbegin",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2"><div>1st new child</div><div>2nd new child</div>El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'beforeEnd'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			Expect(
				el.InsertAdjacentHTML(
					"beforeend",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  <div>1st new child</div><div>2nd new child</div></div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'afterend'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			Expect(
				el.InsertAdjacentHTML(
					"afterend",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div><div>1st new child</div><div>2nd new child</div>
  <div id="3">El 1</div>
</body>`))
		})
	})

	Describe("HTML Rendering", func() {
		It("Should support OuterHTML", func() {
			doc := ParseHtmlString(`<body><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div></body>`)
			Expect(doc.Body().OuterHTML()).To(Equal(`<body><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div></body>`))
			Expect(doc.Body().InnerHTML()).To(Equal(`<div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>`))
		})
	})

	Describe("Click", func() {
		It("Is cancelable and bubbles", func() {
			var event Event
			doc := ParseHtmlString(`<body><div id="target"></div></body>`)
			element := doc.GetElementById("target")
			element.AddEventListener("click", NewEventHandlerFuncWithoutError(func(e Event) {
				event = e
			}))
			element.Click()
			Expect(event.Cancelable()).To(BeTrue())
			Expect(event.Bubbles()).To(BeTrue())
		})
	})

	Describe("Append/Prepend/ReplaceChildren/Children", func() {
		// These should ideaaly be tested on both Element, Document, and
		// DocumentFragment. The functions are defined in the ParentNode
		// mixin interface that all 3 share.

		It("Should add elements in 'Append'", func() {
			doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
			b := doc.Body()
			divE := doc.CreateElement("div")
			divE.SetTextContent("e")
			b.Append(
				doc.CreateText("d"),
				divE,
				doc.CreateText("f"),
			)
			Expect(b).To(HaveOuterHTML(`<body>a<div>b</div>cd<div>e</div>f</body>`))
		})

		It("Should add elements first in 'Prepend'", func() {
			doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
			b := doc.Body()
			divE := doc.CreateElement("div")
			divE.SetTextContent("e")
			b.Prepend(
				doc.CreateText("d"),
				divE,
				doc.CreateText("f"),
			)
			Expect(b).To(HaveOuterHTML(`<body>d<div>e</div>fa<div>b</div>c</body>`))
		})

		It("Should replace elements first in 'ReplaceChildren'", func() {
			doc := ParseHtmlString(`<body>a<div>b</div>c</body>`)
			b := doc.Body()
			divE := doc.CreateElement("div")
			divE.SetTextContent("e")
			b.ReplaceChildren(
				doc.CreateText("d"),
				divE,
				doc.CreateText("f"),
			)
			Expect(b).To(HaveOuterHTML(`<body>d<div>e</div>f</body>`))
		})

		It("Should iterate elements in 'Children'", func() {
			doc := ParseHtmlString(
				`<body>a<div id="el-1">b</div>c<div id="el-2">d</div>e<div name="el-3">f</div>g</body>`,
			)
			b := doc.Body()
			c := b.Children()

			Expect(c.Length()).To(Equal(3))
			Expect(c.Item(0)).To(HaveOuterHTML(`<div id="el-1">b</div>`))
			Expect(c.Item(1)).To(HaveAttribute("id", "el-2"))
			Expect(c.Item(2)).To(HaveAttribute("name", "el-3"))

			Expect(c.Item(-1)).To(BeNil())
			Expect(c.Item(3)).To(BeNil())

			Expect(c.NamedItem("el-1")).To(HaveAttribute("id", "el-1"))
			Expect(c.NamedItem("el-2")).To(HaveAttribute("id", "el-2"))
			Expect(c.NamedItem("el-3")).To(HaveAttribute("name", "el-3"))
		})

		Describe("First/Last element", func() {
			It("Should return elements when they exist", func() {
				doc := ParseHtmlString(
					`<body>a<div id="el-1">b</div>c<div id="el-2">d</div>e<div name="el-3">f</div>g</body>`,
				)
				b := doc.Body()
				Expect(b.FirstElementChild()).To(HaveAttribute("id", "el-1"))
				Expect(b.LastElementChild()).To(HaveOuterHTML(`<div name="el-3">f</div>`))
			})

			It("Should nil when there are only non-element children", func() {
				doc := ParseHtmlString(
					`<body>body text</body>`,
				)
				b := doc.Body()
				Expect(b.FirstElementChild()).To(BeNil())
				Expect(b.LastElementChild()).To(BeNil())
			})
		})

		It("Should handle empty string correctly in 'NamedItem'", func() {
			Skip("Need to research")
		})
	})
})
