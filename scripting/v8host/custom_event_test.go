package v8host_test

import (
	"testing"

	"github.com/gost-dom/browser/internal/testing/browsertest"
	"github.com/stretchr/testify/assert"
)

func TestCustomEvent(t *testing.T) {
	win := browsertest.InitWindow(t)
	assert.Equal(t, "foo", win.MustEval(`
		const e = new CustomEvent("foo", { detail: { f: "foo", b: "bar" }})
		const d = e.detail
		d.f
	`))
}
