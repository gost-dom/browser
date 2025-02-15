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

func TestGeneratedDomTypes(t *testing.T) {
	suite.Run(t, new(DomSuite))
}
