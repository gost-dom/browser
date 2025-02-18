package htmlelements_test

import (
	"bytes"
	"testing"

	htmlelements "github.com/gost-dom/code-gen/html-elements"
	"github.com/gost-dom/generators"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

type expect func(actual interface{}, extra ...interface{}) types.Assertion

func newGomega(t *testing.T) expect { return gomega.NewWithT(t).Expect }

func generate(spec string, name string) (generators.Generator, error) {
	x := htmlelements.PacageConfigs[spec]
	r := x[name]
	g, err := htmlelements.CreateGenerator(r)
	return g.GenerateInterface(), err
}

func TestStringifierShouldNotMakeInterfaceFail(t *testing.T) {
	// The stringifier is an unnamed operation and shouldn't be included in the
	// go interface; but will generate an error if not handled in code
	expect := newGomega(t)
	g, err := generate("url", "url_search_params")
	expect(err).NotTo(gomega.HaveOccurred())
	var b bytes.Buffer
	expect(g.Generate().Render(&b)).To(gomega.Succeed())
}
