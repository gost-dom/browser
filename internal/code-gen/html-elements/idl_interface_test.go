package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/code-gen/idltransform"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/gost-dom/webref/idl"
)

func NewStringAttribute(name string) IdlInterfaceAttribute {
	return IdlInterfaceAttribute{
		Name: name,
		Type: idltransform.NewIdlType(idl.Type{Name: "DOMString"}),
	}
}

var _ = Describe("IdlInterface", func() {
	It("Should generate an interface", func() {
		actual := IdlInterface{
			Name:       "HTMLAnchorElement",
			Attributes: []IdlInterfaceAttribute{NewStringAttribute("target")},
		}
		Expect(actual).To(HaveRendered(
			`type HTMLAnchorElement interface {
	Target() string
	SetTarget(string)
}`))
	})

	It("Should add inherited type", func() {
		actual := IdlInterface{
			Name:       "HTMLAnchorElement",
			Inherits:   "HTMLElement",
			Attributes: []IdlInterfaceAttribute{NewStringAttribute("target")},
		}
		Expect(actual).To(HaveRendered(
			`type HTMLAnchorElement interface {
	HTMLElement
	Target() string
	SetTarget(string)
}`))
	})

	It("Should not sanitize Type", func() {
		actual := IdlInterface{
			Name:       "HTMLAnchorElement",
			Attributes: []IdlInterfaceAttribute{NewStringAttribute("type")},
		}
		Expect(actual).To(HaveRendered(
			`type HTMLAnchorElement interface {
	Type() string
	SetType(string)
}`))

	})
})
