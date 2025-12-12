package entity_test

import (
	"testing"

	. "github.com/gost-dom/browser/internal/entity"
	"github.com/stretchr/testify/assert"
)

var _ Components = &Entity{}

type Key struct{}
type NonKey struct{}

type StringKeyType string
type NonStringKeyType string

var KeyVal StringKeyType = "Key"
var NonKeyVal NonStringKeyType = "Key"

func TestEntityComponents(t *testing.T) {
	t.Run("Retrieve same value", func(t *testing.T) {
		var e Entity
		SetComponent(&e, Key{}, "foobar")
		actual, ok := Component(&e, Key{})
		assert.True(t, ok)
		assert.Equal(t, "foobar", actual)
	})

	t.Run("Insert invalid value", func(t *testing.T) {
		var e Entity
		assert.Panics(t, func() {
			SetComponent(&e, []string{}, "Foobar")
		})
		assert.Panics(t, func() {
			Component(&e, []string{})
		})
	})

	t.Run("It treats keys as nominal types", func(t *testing.T) {
		var e Entity
		SetComponent(&e, Key{}, "Foobar")
		val, found := Component(&e, NonKey{})
		assert.Nil(t, val)
		assert.False(t, found)

		val, found = Component(&e, struct{}{})
		assert.Nil(t, val)
		assert.False(t, found)
	})

	t.Run("It treats string keys as nominal types", func(t *testing.T) {
		var e Entity
		SetComponent(&e, KeyVal, "Foobar")
		val, found := Component(&e, NonKeyVal)
		assert.Nil(t, val)
		assert.False(t, found)
	})
}
