package htmlelements_test

import (
	"testing"

	"github.com/onsi/gomega"

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

func TestIDLInterface(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	actual := IdlInterface{
		Name:       "HTMLAnchorElement",
		Attributes: []IdlInterfaceAttribute{NewStringAttribute("target")},
	}
	Expect(actual).To(HaveRendered(lines(
		`type HTMLAnchorElement interface {`,
		`	Target() string`,
		`	SetTarget(string)`,
		`}`,
	)))
}

func TestIDLInterfaceInheritance(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	actual := IdlInterface{
		Name:       "HTMLAnchorElement",
		Inherits:   "HTMLElement",
		Attributes: []IdlInterfaceAttribute{NewStringAttribute("target")},
	}
	Expect(actual).To(HaveRendered(lines(
		`type HTMLAnchorElement interface {`,
		`	HTMLElement`,
		`	Target() string`,
		`	SetTarget(string)`,
		`}`,
	)))
}
