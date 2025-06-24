package gomegamatchers

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

// Wrapper of [gomega.Expect] that accepts a [testing.TB] argument.
//
// Gomega's assertion requires first creating a gomega value initialized with a
// testing.T value.
//
// While gomega accepts multiple values, the usefulness is diminished in this
// variation. Implicit error checks doesn't work as well as the base Expect
// function; requiring the error the exist as an explicit value anyway
func Expect(t testing.TB, val any) types.Assertion {
	return gomega.NewWithT(t).Expect(val)
}
