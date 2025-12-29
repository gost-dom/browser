package events_test

import (
	"bytes"
	"testing"

	"github.com/gost-dom/code-gen/events"
	g "github.com/gost-dom/generators"

	// . "github.com/gost-dom/generators/testing/matchers"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
)

func render(t testing.TB, gen g.Generator) string {
	t.Helper()
	var b bytes.Buffer
	err := gen.Generate().Render(&b)
	if err != nil {
		t.Fatalf("Error rendering: %v", err)
	}
	return b.String()
}

func TestGenerateEventInit(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	gen, err := events.GenerateEventInit("KeyboardEventInit", "uievents")
	require.NoError(t, err)
	rendered := render(t, gen)

	g.Expect(rendered).To(gomega.ContainSubstring(
		"type KeyboardEventInit struct {\n\tEventModifierInit",
	))
	g.Expect(rendered).To(gomega.MatchRegexp(
		"\n\tKey *string\n",
	))
}
