package htmlelements_test

import (
	"testing"

	. "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/webref/idl"
	"github.com/onsi/gomega"
)

var DOMStringType = idl.Type{Name: "DOMString"}
var DOMBooleanType = idl.Type{Name: "boolean"}

func TestIdlTypeGeneratesCorrespondingGoType(t *testing.T) {
	g := gomega.NewWithT(t)
	g.Expect(IdlType(DOMStringType)).To(HaveRendered("string"))
	g.Expect(IdlType(DOMBooleanType)).To(HaveRendered("bool"))
}

func TestIdlTypeGeneratesSequenceTypes(t *testing.T) {
	g := gomega.NewWithT(t)
	stringSequence := idl.Type{
		Kind:      idl.KindSequence,
		TypeParam: &DOMStringType,
	}
	g.Expect(IdlType(stringSequence)).To(HaveRendered("[]string"))
}
