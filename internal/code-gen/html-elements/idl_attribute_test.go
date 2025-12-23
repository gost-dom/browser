package htmlelements_test

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/gost-dom/code-gen/codegentest"
	. "github.com/gost-dom/code-gen/html-elements"
	g "github.com/gost-dom/generators"
	. "github.com/gost-dom/generators/testing/matchers"
)

func TestIDLAttributeGetterAndSetter(t *testing.T) {
	t.Run("Temporary disabled solution", func(t *testing.T) {
		Expect := gomega.NewWithT(t).Expect
		t.Skip("Temporary different solution")
		Expect(IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		}).To(HaveRendered(codegentest.Lines(
			"func (e *htmlAnchorElement) Target() string {",
			"\treturn e.target",
			"}",
			"",
			"func (e *htmlAnchorElement) SetTarget(val string) {",
			"\te.target = val",
			"https://docs.couchdb.org/en/stable/api/index.html}")))
	})

	t.Run("Generates getter and setter", func(t *testing.T) {
		Expect := gomega.NewWithT(t).Expect
		Expect(IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		}).To(HaveRendered(codegentest.Lines(
			`func (e *htmlAnchorElement) Target() string {`,
			`	result, _ := e.GetAttribute("target")`,
			`	return result`,
			`}`,
			``,
			`func (e *htmlAnchorElement) SetTarget(val string) {`,
			`	e.SetAttribute("target", val)`,
			`}`,
		)))
	})

	t.Run("Don't sanitize exported names", func(t *testing.T) {
		// Some unexported names, like `type` conflict with go reserved words.
		// This test is about NOT sanitizing exported names
		Expect := gomega.NewWithT(t).Expect
		actual := IDLAttribute{
			AttributeName: "type",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
			ReadOnly: true,
		}
		Expect(actual).To(HaveRenderedSubstring(`Type() string`))
		Expect(actual).ToNot(HaveRenderedSubstring(`SetType()`))
	})
}

func TestIDLAttributeReadonlyGetter(t *testing.T) {
	Expect := gomega.NewWithT(t).Expect
	actual := IDLAttribute{
		AttributeName: "target",
		Receiver: Receiver{
			Name: g.Id("e"),
			Type: g.NewType("htmlAnchorElement").Pointer(),
		},
		ReadOnly: true,
	}
	Expect(actual).To(HaveRenderedSubstring(`Target() string`))
	Expect(actual).ToNot(HaveRenderedSubstring(`SetTarget()`))
}
