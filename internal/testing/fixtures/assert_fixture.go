package fixtures

import (
	"github.com/gost-dom/fixture"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/stretchr/testify/assert"
)

type AssertFixture struct {
	fixture.Fixture
	assert *assert.Assertions
	gomega gomega.Gomega
}

func (f *AssertFixture) Assert() *assert.Assertions {
	if f.assert == nil {
		f.assert = assert.New(f.TB)
	}
	return f.assert
}

func (f *AssertFixture) Expect(actual interface{}, extra ...interface{}) types.Assertion {
	if f.gomega == nil {
		f.gomega = gomega.NewWithT(f.TB)
	}
	return f.gomega.Expect(actual, extra...)
}
