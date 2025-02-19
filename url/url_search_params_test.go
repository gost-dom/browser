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
	assert.Equal(t, "bar", u.Get("foo"))

	// Same value can be added twice
	u.Append("foo", "baz")
	assert.Equal(t, "foo=bar&foo=baz", u.String())
	assert.Equal(t, []string{"bar", "baz"}, u.GetAll("foo"))

	u = &URLSearchParams{}
	u.Append("foo", "Bar value")
	assert.Equal(t, "foo=Bar+value", u.String(), "Query encoding")
}

func TestParseUrlSearchParams(t *testing.T) {
	// Test percent decoding
	var u, err = ParseURLSearchParams("foo=bar%20value")
	assert.NoError(t, err)
	assert.Equal(t, "bar value", u.Get("foo"))

	// Test + decoding
	u, err = ParseURLSearchParams("foo=bar+value")
	assert.NoError(t, err)
	assert.Equal(t, "bar value", u.Get("foo"))

	// Test leading ? is ignores
	u, err = ParseURLSearchParams("?foo=bar%20value")
	assert.NoError(t, err)
	assert.Equal(t, "bar value", u.Get("foo"))
}
