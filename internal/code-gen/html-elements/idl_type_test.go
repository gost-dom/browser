package htmlelements_test

import (
	"testing"

	. "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/webref/idl"
	"github.com/onsi/gomega"
)

var DOMStringType = idl.Type{Name: "DOMString"}

func TestIdlTypeGeneratesCorrespondingGoType(t *testing.T) {
	idlType := IdlType(DOMStringType)
	g := gomega.NewWithT(t)
	g.Expect(idlType).To(HaveRendered("string"))
}
