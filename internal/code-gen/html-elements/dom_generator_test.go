package htmlelements_test

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/suite"

	. "github.com/gost-dom/code-gen/html-elements"
	g "github.com/gost-dom/generators"
)

func GenerateURL() (g.Generator, error) {
	g, err := CreateGenerator(FileGenerationConfig["url"])
	return g.GenerateInterface(), err
}

func generateDomType(r HTMLGeneratorReq) (g.Generator, error) {
	g, err := CreateGenerator(r)
	return g.GenerateInterface(), err
}

func GenerateParentNode() (g.Generator, error) {
	g, err := CreateGenerator(
		HTMLGeneratorReq{
			InterfaceName:     "ParentNode",
			SpecName:          "dom",
			GenerateInterface: true,
		})
	return g.GenerateInterface(), err
}

type DomSuite struct {
	suite.Suite
	gomega.Gomega
}

func (s *DomSuite) SetupTest() {
	s.Gomega = gomega.NewWithT(s.T())
}

func (s *DomSuite) TestGenerateGetterAndSetterOnURL() {
	s.Expect(
		generateDomType(FileGenerationConfig["url"]),
	).To(HaveRendered(ContainSubstring(
		`ToJSON() (string, error)`)))
}

func (s *DomSuite) TestGenerateHTMLCollection() {
	generator, err := generateDomType(HTMLGeneratorReq{
		InterfaceName:     "HTMLCollection",
		SpecName:          "dom",
		GenerateInterface: true,
	})
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(generator).To(HaveRendered(ContainSubstring(`Length() int`)))
	s.Expect(generator).To(HaveRendered(ContainSubstring(`Item(int) Element`)))
}

func (s *DomSuite) TestGenerateParentNode() {
	generator, err := generateDomType(HTMLGeneratorReq{
		InterfaceName:     "ParentNode",
		SpecName:          "dom",
		GenerateInterface: true,
	})
	s.Expect(err).ToNot(HaveOccurred())
	s.Expect(generator).To(HaveRendered(ContainSubstring("Append(nodes ...Node)\n")))
}

/*
ParentNode generates this:
	Children() string
	FirstElementChild() string
	LastElementChild() string
	ChildElementCount() string
	Prepend() (string, error)
	Append() (string, error)
	ReplaceChildren() (string, error)
	QuerySelector() (string, error)
	QuerySelectorAll() (string, error)
*/

func TestGeneratedDomTypes(t *testing.T) {
	suite.Run(t, new(DomSuite))
}
