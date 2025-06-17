package idltransform_test

import (
	"testing"

	"github.com/gost-dom/code-gen/idltransform"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/gost-dom/webref/idl"
	"github.com/onsi/gomega"
)

var DOMStringType = idl.Type{Name: "DOMString"}
var DOMBooleanType = idl.Type{Name: "boolean"}

func TestIdlTypeGeneratesCorrespondingGoType(t *testing.T) {
	g := gomega.NewWithT(t)
	g.Expect(idltransform.NewIdlType(DOMStringType)).To(HaveRendered("string"))
	g.Expect(idltransform.NewIdlType(DOMBooleanType)).To(HaveRendered("bool"))
}

func TestIdlTypeGeneratesSequenceTypes(t *testing.T) {
	g := gomega.NewWithT(t)
	stringSequence := idl.Type{
		Kind:      idl.KindSequence,
		TypeParam: &DOMStringType,
	}
	g.Expect(idltransform.NewIdlType(stringSequence)).To(HaveRendered("[]string"))
}
