package url_test

import (
	"testing"

	i "github.com/gost-dom/browser/internal/interfaces/url-interfaces"
	. "github.com/gost-dom/browser/url"
	"github.com/stretchr/testify/assert"
)

func TestUrlSearchParams(t *testing.T) {
	// Declare a var of the interface type to compile time check that we adhere
	// to the interface
	var u i.URLSearchParams = &URLSearchParams{}
	u.Append("foo", "bar")
	assert.Equal(t, "foo=bar", u.String())
}
