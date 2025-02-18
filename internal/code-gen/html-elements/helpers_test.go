package htmlelements_test

import (
	"bytes"
	"fmt"
	"testing"

	htmlelements "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/generators"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

type generator = generators.Generator

type expect func(actual interface{}, extra ...interface{}) types.Assertion

func newGomega(t *testing.T) expect { return gomega.NewWithT(t).Expect }

type BaseGenerator interface{ GenerateInterface() generator }

func getIdlInterfaceGenerator(
	packageName string,
	interfaceName string,
) (generators.Generator, error) {
	packageSpecs, _ := htmlelements.GetPackageGeneratorSpecs(packageName)
	for _, v := range packageSpecs {
		if v.InterfaceName == interfaceName {
			g, err := htmlelements.CreateGenerator(v)
			return g.GenerateInterface(), err
		}
	}
	return nil, fmt.Errorf(
		"getIdlInterfaceGenerator: IDL Interface %s not configured for package %s",
		interfaceName,
		packageName,
	)
}

func render(g generator) (string, error) {
	var b bytes.Buffer
	err := g.Generate().Render(&b)
	return b.String(), err
}
