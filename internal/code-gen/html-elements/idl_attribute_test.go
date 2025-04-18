package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/gost-dom/code-gen/html-elements"
	g "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
)

var _ = Describe("IDLAttribute", func() {
	It("Should generate a getter and setter", func() {
		Skip("Temporary different solution")
		Expect(IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		}).To(HaveRendered(
			`func (e *htmlAnchorElement) Target() string {
	return e.target
}

func (e *htmlAnchorElement) SetTarget(val string) {
	e.target = val
}`))
	})

	It("Should generate a getter and setter", func() {
		Expect(IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		}).To(HaveRendered(
			`func (e *htmlAnchorElement) Target() string {
	result, _ := e.GetAttribute("target")
	return result
}

func (e *htmlAnchorElement) SetTarget(val string) {
	e.SetAttribute("target", val)
}`))
	})

	It("Should generate a getter when readOnly", func() {
		actual := IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
			ReadOnly: true,
		}
		Expect(actual).To(HaveRendered(ContainSubstring(`Target() string`)))
		Expect(actual).ToNot(HaveRendered(ContainSubstring(`SetTarget()`)))
	})

	It("Should NOT sanitize Type", func() {
		actual := IDLAttribute{
			AttributeName: "type",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
			ReadOnly: true,
		}
		Expect(actual).To(HaveRendered(ContainSubstring(`Type() string`)))
		Expect(actual).ToNot(HaveRendered(ContainSubstring(`SetType()`)))
	})
})
