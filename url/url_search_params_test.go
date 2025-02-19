package url_test

import (
	"testing"

	i "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	. "github.com/gost-dom/browser/url"
	"github.com/stretchr/testify/assert"
)

func TestUrlSearchParamsConformsToInterface(t *testing.T) {
	// Declare a var of the interface type to compile time check that the type
	// conforms to the IDL spec as good as possible.
	var u i.URLSearchParams = &URLSearchParams{}
	assert.Equal(t, "", u.String())
}

func TestUrlSearchParamsAddValue(t *testing.T) {
	u := &URLSearchParams{}
	u.Append("foo", "bar")
	assert.Equal(t, "foo=bar", u.String())
	v, found := u.Get("foo")
	assert.True(t, found)
	assert.Equal(t, "bar", v)

	// Same value can be added twice
	u.Append("foo", "baz")
	assert.Equal(t, "foo=bar&foo=baz", u.String())
	assert.Equal(t, []string{"bar", "baz"}, u.GetAll("foo"))

	u = &URLSearchParams{}
	u.Append("foo", "Bar value")
	assert.Equal(t, "foo=Bar+value", u.String(), "Query encoding")
}

func TestGet(t *testing.T) {
	u := &URLSearchParams{}
	u.Append("foo", "bar")

	foundValue, ok := u.Get("foo")
	assert.True(t, ok)
	assert.Equal(t, "bar", foundValue)

	missingValue, ok := u.Get("baz")
	assert.False(t, ok)
	assert.Equal(t, "", missingValue)
}

func TestParseUrlSearchParams(t *testing.T) {
	// Test percent decoding
	var u, err = ParseURLSearchParams("foo=bar%20value")
	assert.NoError(t, err)
	v, found := u.Get("foo")
	assert.True(t, found)
	assert.Equal(t, "bar value", v)

	// Test + decoding
	u, err = ParseURLSearchParams("foo=bar+value")
	assert.NoError(t, err)
	v, found = u.Get("foo")
	assert.True(t, found)
	assert.Equal(t, "bar value", v)

	// Test leading ? is ignores
	u, err = ParseURLSearchParams("?foo=bar%20value")
	assert.NoError(t, err)
	v, found = u.Get("foo")
	assert.True(t, found)
	assert.Equal(t, "bar value", v)
}

func TestURLSearchParamsHas(t *testing.T) {
	var u URLSearchParams
	u.Append("foo", "bar")
	assert.True(t, u.Has("foo"))
	assert.False(t, u.Has("bar"))

	assert.True(t, u.HasValue("foo", "bar"))
	assert.False(t, u.HasValue("foo", "baz"))
}
