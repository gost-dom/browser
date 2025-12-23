package htmlelements_test

import (
	"testing"

	"github.com/gost-dom/code-gen/codegentest"
	htmlelements "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/code-gen/packagenames"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"
)

type DomSuite struct {
	suite.Suite
	gomega.Gomega
}

func (s *DomSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

func (s *DomSuite) TestGenerateHTMLCollection() {
	s.T().Skip(
		"HTMLCollection is not generated currently, as it doesn't generate an All() function",
	)
	generator, err := getIdlInterfaceGenerator("dom", "HTMLCollection")
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(generator).To(HaveRendered(ContainSubstring(`Length() int`)))
	s.Expect(generator).To(HaveRendered(ContainSubstring(`Item(int) Element`)))
}

func (s *DomSuite) TestGenerateParentNode() {
	generator, err := htmlelements.GenerateInterface("dom", "dom", "ParentNode")
	s.Expect(err).ToNot(HaveOccurred())
	output := codegentest.RenderInPackage(s.T(), packagenames.Dom, generator)
	s.Expect(output).To(ContainSubstring("Append(...Node) error\n"))
}

func TestGeneratedDomTypes(t *testing.T) {
	suite.Run(t, new(DomSuite))
}
