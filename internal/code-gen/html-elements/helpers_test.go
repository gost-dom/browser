package htmlelements_test

import (
	"bytes"
	"errors"
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

func getFileGenerator(packageName, targetFile string) (generators.Generator, error) {
	packageSpecs, _ := htmlelements.GetPackageGeneratorSpecs(packageName)
	for outputFile, spec := range packageSpecs {
		if outputFile == targetFile {
			return htmlelements.CreateGenerator(spec)
		}
	}
	return nil, errors.New("Unknown package")
}

func getIdlInterfaceGenerator(
	packageName string,
	interfaceName string,
) (generators.Generator, error) {
	packageSpecs, _ := htmlelements.GetPackageGeneratorSpecs(packageName)
	for _, v := range packageSpecs {
		if v.InterfaceName == interfaceName {
			g, err := htmlelements.CreateGenerator(v)
			return g, err
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
