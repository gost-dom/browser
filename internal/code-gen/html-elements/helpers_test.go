package htmlelements_test

import (
	"bytes"
	"testing"

	htmlelements "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/generators"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

type generator = generators.Generator

type expect func(actual interface{}, extra ...interface{}) types.Assertion

func newGomega(t *testing.T) expect { return gomega.NewWithT(t).Expect }

func generateType(spec string, name string) (generators.Generator, error) {
	x := htmlelements.PacageConfigs[spec]
	r := x[name]
	g, err := htmlelements.CreateGenerator(r)
	return g.GenerateInterface(), err
}

func render(g generator) (string, error) {
	var b bytes.Buffer
	err := g.Generate().Render(&b)
	return b.String(), err
}
