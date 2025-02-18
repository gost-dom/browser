package htmlelements_test

import (
	"bytes"
	"testing"

	"github.com/onsi/gomega"
)

func TestStringifierShouldNotMakeInterfaceFail(t *testing.T) {
	// The stringifier is an unnamed operation and shouldn't be included in the
	// go interface; but will generate an error if not handled in code
	expect := newGomega(t)
	g, err := generate("url", "url_search_params")
	expect(err).NotTo(gomega.HaveOccurred())
	var b bytes.Buffer
	expect(g.Generate().Render(&b)).To(gomega.Succeed())
}

func TestGenerateGetterAndSetterOnURL(t *testing.T) {
	Expect := newGomega(t)
	Expect(
		generate("url", "url"),
	).To(HaveRendered(gomega.ContainSubstring(
		`ToJSON() (string, error)`)))
}
