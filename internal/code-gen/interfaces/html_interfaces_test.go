package interfaces_test

import (
	"testing"

	"github.com/gost-dom/code-gen/interfaces"
	. "github.com/gost-dom/generators/testing/matchers"
	"github.com/onsi/gomega"
)

func TestHTMLHistoryInterface(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	gen, err := interfaces.GenerateInterface("html", "htmlinterfaces", "History")
	Expect(err).ToNot(gomega.HaveOccurred())

	Expect(gen).To(HaveRenderedSubstring("\tLength() int\n"))
	Expect(gen).ToNot(HaveRenderedSubstring("ScrollRestoration()"))

	Expect(
		gen,
	).To(HaveRenderedSubstring("\tState() HistoryState\n"), "History state 'overrides' default type")
}
