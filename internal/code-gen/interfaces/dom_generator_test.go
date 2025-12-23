package interfaces_test

import (
	"testing"

	"github.com/gost-dom/code-gen/codegentest"
	"github.com/gost-dom/code-gen/interfaces"
	"github.com/gost-dom/code-gen/packagenames"

	. "github.com/onsi/gomega"
)

func TestDomParentNode(t *testing.T) {
	s := NewGomegaWithT(t)
	generator, err := interfaces.GenerateInterface("dom", "dom", "ParentNode")
	s.Expect(err).ToNot(HaveOccurred())
	output := codegentest.RenderInPackage(t, packagenames.Dom, generator)
	s.Expect(output).To(ContainSubstring("Append(...Node) error\n"))
}
